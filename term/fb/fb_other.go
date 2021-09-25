// +build darwin freebsd solaris windows plan9

package fb

import (
	"fmt"
)

// Query queries the framebuffer for it's screen dimensions
func Query(device string) (*ScreenInfo, error) {
	var si ScreenInfo
	return &si, nil
}

// WriteImage takes a png image and Writes the raw RGBA pixel
// data to the device named. Only Paletted RGBA PNG images are supported.
func WriteImage(device string, data []byte) error {
	return fmt.Errorf("WriteImage: not supported on this OS")
}
