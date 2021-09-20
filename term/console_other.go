// +build darwin freebsd solaris windows plan9

package term

func QueryFramebuffer() *ScreenInfo {
	var si ScreenInfo
	return &si
}
