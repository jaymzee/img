package term

import "os"

// Isatty determines if stdout is a tty (is it a char mode device?)
func Isatty() bool {
	fileInfo, _ := os.Stdout.Stat()
	return fileInfo.Mode()&os.ModeCharDevice != 0
}
