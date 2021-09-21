package term

// #cgo CFLAGS:
// #cgo LDFLAGS:
// #include "query_fb.h"
import "C"

func QueryFramebuffer(device string) *ScreenInfo {
	var fbinfo C.struct_fb_var_screeninfo
	C.query_framebuffer(C.CString(device), &fbinfo)

	return &ScreenInfo{
		Xres:  uint(fbinfo.xres),
		Yres:  uint(fbinfo.yres),
		Xresv: uint(fbinfo.xres_virtual),
		Yresv: uint(fbinfo.yres_virtual),
		Bpp:   uint(fbinfo.bits_per_pixel),
	}
}
