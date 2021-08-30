package term

import (
	"golang.org/x/sys/unix"
	"os"
)

func GetWinsize() (*Winsize, error) {
	ws, err := unix.IoctlGetWinsize(int(os.Stdout.Fd()), unix.TIOCGWINSZ)
	if err != nil {
		return nil, os.NewSyscallError("GetWinsize", err)
	}

	return &Winsize{ws.Row, ws.Col, ws.Xpixel, ws.Ypixel}, nil
}
