// +build linux darwin freebsd solaris

package term

// #cgo CFLAGS:
// #cgo LDFLAGS:
// #include <unistd.h>
import "C"

func TtyName() string {
	return C.GoString(C.ttyname(C.STDIN_FILENO))
}
