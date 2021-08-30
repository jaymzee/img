// +build linux darwin freebsd

package term

import (
	"golang.org/x/sys/unix"
	"os"
)

func GetWinsize() *Winsize {
	ws, err := unix.IoctlGetWinsize(int(os.Stdout.Fd()), unix.TIOCGWINSZ)
	if err != nil {
		return newWinsize()
	}
	return &Winsize{ws.Row, ws.Col, ws.Xpixel, ws.Ypixel}
}
