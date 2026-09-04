package cmigemo_test

import (
	"testing"

	"github.com/koron/cmigemo-go"
)

func Test_Version(t *testing.T) {
	ver := cmigemo.Version()
	if ver == "" {
		t.Fatal("empty version")
	}
	t.Logf("migemo version: %s", ver)
}

func Test_Open(t *testing.T) {
	m, err := cmigemo.Open(`/usr/local/share/cmigemo/utf-8/migemo-dict`)
	if err != nil {
		t.Fatalf("can't open: %s", err)
	}
	defer m.Close()

	s := m.Query("aka")
	t.Logf("aka -> %q", s)
}
