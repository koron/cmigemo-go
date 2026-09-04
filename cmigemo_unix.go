//go:build darwin || freebsd || linux

package cmigemo

import (
	"fmt"
	"runtime"

	"github.com/ebitengine/purego"
)

func getMigemoLibrary() string {
	switch runtime.GOOS {
	case "darwin":
		return "libmigemo.1.dylib"
	case "freebsd":
		return "libmigemo.so.1"
	case "linux":
		return "libmigemo.so.1"
	default:
		panic(fmt.Errorf("GOOS=%s is not supported", runtime.GOOS))
	}
}

func openLibrary() (uintptr, error) {
	name := getMigemoLibrary()
	return purego.Dlopen(name, purego.RTLD_NOW|purego.RTLD_GLOBAL)
}
