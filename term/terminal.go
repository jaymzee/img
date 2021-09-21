package term

// #cgo CFLAGS:
// #cgo LDFLAGS:
// #include "cursor.h"
import "C"
import (
	"fmt"
	"os"
	"regexp"
)

// Isatty determines if stdout is a tty (is it a char mode device?)
func Isatty() bool {
	fileInfo, _ := os.Stdout.Stat()
	return fileInfo.Mode()&os.ModeCharDevice != 0
}

// Isaconsole determines if stdout is a console (/dev/tty1 thru /dev/tty6)
func Isaconsole() bool {
	pattern := regexp.MustCompile(`/dev/tty\d`)
	return pattern.MatchString(TtyName())
}

// GetCursorCood queries the terminal for it's current position on screen
func GetCursorCoord() (int, int, error) {
	var x, y C.int

	if C.getCursor(&x, &y) != 0 {
		return 0, 0, fmt.Errorf("failed to read cursor location")
	}
	return int(x), int(y), nil
}
