package cmigemo

import "syscall"

func openLibrary() (uintptr, error) {
	h, err := syscall.LoadLibrary("libmigemo.dll")
	return uintptr(h), err
}
