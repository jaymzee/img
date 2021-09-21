package term

// #cgo CFLAGS:
// #cgo LDFLAGS:
// #include "fb_query.h"
import "C"
import (
	"fmt"
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
		xres, yres, err := queryfb("/dev/fb0")
		// fallback on error instead of fail
		if err == nil {
			return &Winsize{ws.Row, ws.Col, xres, yres}
		}
	}
	return &Winsize{ws.Row, ws.Col, ws.Xpixel, ws.Ypixel}
}

// copied from fb package because it avoids import cycles
// query the framebuffer for it's screen dimensions and bits per pixel
func queryfb(device string) (uint16, uint16, error) {
	var fbinfo C.struct_fb_var_screeninfo

	if C.query_fb(C.CString(device), &fbinfo) != 0 {
		return 0, 0, fmt.Errorf("%s: permission denied")
	}

	return uint16(fbinfo.xres), uint16(fbinfo.yres), nil
}
