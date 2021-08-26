// +build !windows

package term

import (
	"os"
	"syscall"
	"unsafe"
)

const (
	_TIOCGWINSZ = 0x5413 // On OSX use 1074295912. Thanks zeebo
)

type Winsize struct {
	Rows uint16
	Cols uint16
	Xres uint16
	Yres uint16
}

func GetWinsize() (*Winsize, error) {
	ws := new(Winsize)

	r1, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(_TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)),
	)

	if int(r1) == -1 {
		return nil, os.NewSyscallError("GetWinsize", errno)
	}
	return ws, nil
}
