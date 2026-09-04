// Package cmigemo provides a wrapper of koron/cmigemo without CGO.
package cmigemo

import (
	"fmt"
	"os"
	"regexp"
	"sync"
	"unsafe"

	"github.com/ebitengine/purego"
)

var (
	migemoLib            uintptr
	migemoVersion        func() string
	migemoOpen           func(string) uintptr
	migemoClose          func(uintptr)
	migemoQuery          func(uintptr, string) uintptr
	migemoRelease        func(uintptr, uintptr)
	migemoSetOperator    func(uintptr, string) int
	migemoSetEscapeChars func(uintptr, string)
)

var initLib = sync.OnceFunc(func() {
	p, err := openLibrary()
	if err != nil {
		panic(fmt.Sprintf("failed to open cmigemo library: %s", err))
	}
	migemoLib = p
	purego.RegisterLibFunc(&migemoVersion, p, "migemo_version")
	purego.RegisterLibFunc(&migemoOpen, p, "migemo_open")
	purego.RegisterLibFunc(&migemoClose, p, "migemo_close")
	purego.RegisterLibFunc(&migemoQuery, p, "migemo_query")
	purego.RegisterLibFunc(&migemoRelease, p, "migemo_release")
	purego.RegisterLibFunc(&migemoSetOperator, p, "migemo_set_operator")
	purego.RegisterLibFunc(&migemoSetEscapeChars, p, "migemo_set_escape_chars")
})

func ptr2str(p uintptr) string {
	ptr := *(*unsafe.Pointer)(unsafe.Pointer(&p))
	if ptr == nil {
		return ""
	}
	var l int
	// Limit to 1 MiB for safety.
	for ; l < 1024*1024; l++ {
		if *(*byte)(unsafe.Add(ptr, uintptr(l))) == '\x00' {
			break
		}
	}
	return string(unsafe.Slice((*byte)(ptr), l))
}

func isRegularFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.Mode().IsRegular()
}

func Version() string {
	initLib()
	return migemoVersion()
}

type Migemo struct {
	pmo uintptr
}

func Open(migemoDict string) (*Migemo, error) {
	initLib()
	if !isRegularFile(migemoDict) {
		return nil, fmt.Errorf("migemo dictionary not found: %s", migemoDict)
	}
	pmo := migemoOpen(migemoDict)
	if pmo == 0 {
		return nil, fmt.Errorf("migemo_open failed: %s", migemoDict)
	}
	migemoSetEscapeChars(pmo, `\.+*?()|[]{}^$`)
	return &Migemo{pmo: pmo}, nil
}

func (mo *Migemo) Close() {
	if mo.pmo != 0 {
		migemoClose(mo.pmo)
		mo.pmo = 0
	}
}

func (mo *Migemo) Query(query string) string {
	p := migemoQuery(mo.pmo, query)
	if p == 0 {
		return ""
	}
	s := ptr2str(p)
	migemoRelease(mo.pmo, p)
	return s
}

func (mo *Migemo) Regexp(query string) (*regexp.Regexp, error) {
	s := mo.Query(query)
	if s == "" {
		return nil, fmt.Errorf("migemo_query returns an empty string for: %q", query)
	}
	return regexp.Compile(s)
}
