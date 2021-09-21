package term

import (
	"golang.org/x/sys/unix"
	"os"
)

// GetWinsize returns the console window size. If the os doesn't have an
// api to determine this, a default window size of 24 80 0 0 is returned.
func GetWinsize() *Winsize {
	ws, err := unix.IoctlGetWinsize(int(os.Stdout.Fd()), unix.TIOCGWINSZ)
	if err != nil {
		return newWinsize()
	}

	if Isaconsole() {
		si, err := QueryFramebuffer("/dev/fb0")
		// fallback on error instead of fail
		if err == nil {
			return &Winsize{ws.Row, ws.Col, uint16(si.Xres), uint16(si.Yres)}
		}
	}
	return &Winsize{ws.Row, ws.Col, ws.Xpixel, ws.Ypixel}
}
