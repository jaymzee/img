// +build linux darwin freebsd solaris

package term

// #cgo CFLAGS:
// #cgo LDFLAGS:
// #include <unistd.h>
import "C"
import (
	"fmt"
	"os"
	"regexp"
)

func TtyName() string {
	return C.GoString(C.ttyname(C.STDIN_FILENO))
}

// GetCursorCood queries the terminal for it's current position on screen
func GetCursorCoord() (int, int, error) {
	var x, y C.int

	if C.getCursor(&x, &y) != 0 {
		return 0, 0, fmt.Errorf("failed to read cursor location")
	}
	return int(x), int(y), nil
}
