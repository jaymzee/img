package term

import "regexp"

func Isaconsole() bool {
	pattern := regexp.MustCompile(`/dev/tty\d`)
	return pattern.MatchString(TtyName())
}

type ScreenInfo struct {
	Xres  uint
	Yres  uint
	Xresv uint
	Yresv uint
	Bpp   uint
}
