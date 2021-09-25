package term

// #cgo CFLAGS:
// #cgo LDFLAGS:
// #include "noop.h"
import "C"
import (
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
