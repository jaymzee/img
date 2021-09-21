package term

// #cgo CFLAGS:
// #cgo LDFLAGS:
// #include "query_fb.h"
import "C"
import "fmt"

func QueryFramebuffer(device string) (*ScreenInfo, error) {
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
