package fb

// #cgo CFLAGS:
// #cgo LDFLAGS:
// #include "fb.h"
import "C"

import (
	"bytes"
	"fmt"
	"github.com/jaymzee/img/term"
	"image"
	"image/color"
	"image/png"
	"unsafe"
)

type ScreenInfo struct {
	Xres  uint
	Yres  uint
	Xresv uint
	Yresv uint
	Bpp   uint
}

// WriteImage takes a png image and Writes the raw RGBA pixel
// data to the device named. Only Paletted RGBA PNG images are supported.
func WriteImage(device string, data []byte) error {
	// get screen dimensions, bit depth, text cell size, framebuffer padding
	winsize := term.GetWinsize()
	cellwidth := int(winsize.Xres / winsize.Cols)
	cellheight := int(winsize.Yres / winsize.Rows)
	_, cur_y, err := term.GetCursorCoord()
	if err != nil {
		return err
	}
	scrinfo, err := Query(device)
	if err != nil {
		return err
	}
	if scrinfo.Bpp != 32 {
		return fmt.Errorf("%s: display must be 32 bits per pixel, got %v", device, scrinfo)
	}
	var pad int = int(scrinfo.Xresv) - int(scrinfo.Xres)
	if pad < 0 {
		return fmt.Errorf("%s: xres_virtual less than xres, got %v", device, scrinfo)
	}

	// decode PNG data
	reader := bytes.NewReader(data)
	img, err := png.Decode(reader)
	if err != nil {
		return err
	}

	// write image data to the C buffer
	dx := img.Bounds().Dx()
	dy := img.Bounds().Dy()
	cimg := C.new_image(C.int(dx), C.int(dy))
	if cimg == nil {
		return fmt.Errorf("cgo memory allocation failed")
	}
	if img, ok := img.(*image.Paletted); ok {
		n := 0
		pix := (*[1 << 24]C.char)(unsafe.Pointer(&cimg.pix))
		for _, indx := range img.Pix {
			if rgba, ok := img.Palette[indx].(color.RGBA); ok {
				pix[n+0] = (C.char)(rgba.B)
				pix[n+1] = (C.char)(rgba.G)
				pix[n+2] = (C.char)(rgba.R)
				n += 4
			} else {
				return fmt.Errorf("expected palette to be in RGBA format")
			}
		}
	} else {
		return fmt.Errorf("expected a paletted PNG image")
	}

	// write pixels to framebuffer
	fbinfo := C.struct_fbinfo{
		xres:   C.int(scrinfo.Xres),
		yres:   C.int(scrinfo.Yres),
		pad:    C.int(pad),
		device: C.CString(device),
	}
	for i := 0; i < dy / cellheight + 1; i++ {
		fmt.Println();
	}
	if C.write_image(cimg, C.int(6 * cellwidth), C.int(cur_y * cellheight), &fbinfo) != 0 {
		return fmt.Errorf("%s: write to framebuffer failed", device)
	}
	return nil
}
