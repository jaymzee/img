// +build darwin freebsd solaris windows plan9

package term

func QueryFramebuffer() (*ScreenInfo, error) {
	var si ScreenInfo
	return &si, nil
}
