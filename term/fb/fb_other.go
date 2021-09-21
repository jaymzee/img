// +build darwin freebsd solaris windows plan9

package fb

// Query queries the framebuffer for it's screen dimensions
func Query(device string) (*ScreenInfo, error) {
	var si ScreenInfo
	return &si, nil
}
