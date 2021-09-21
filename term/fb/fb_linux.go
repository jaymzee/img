package fb

// #cgo CFLAGS:
// #cgo LDFLAGS:
// #include "fb_query.h"
import "C"
import "fmt"

// Query queries the framebuffer for it's screen dimensions and bits per pixel
func Query(device string) (*ScreenInfo, error) {
	var fbinfo C.struct_fb_var_screeninfo

	if C.query_framebuffer(C.CString(device), &fbinfo) != 0 {
		return nil, fmt.Errorf("%s: permission denied")
	}

	return &ScreenInfo{
		Xres:  uint(fbinfo.xres),
		Yres:  uint(fbinfo.yres),
		Xresv: uint(fbinfo.xres_virtual),
		Yresv: uint(fbinfo.yres_virtual),
		Bpp:   uint(fbinfo.bits_per_pixel),
	}, nil
}
