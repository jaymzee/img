package fb

// #cgo CFLAGS:
// #cgo LDFLAGS:
// #include "fb.h"
import "C"

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/jaymzee/img/term"
	"image"
	"image/png"
	"unsafe"
)

// WriteImageToFramebuffer takes a png image and Writes the raw RGBA pixel
// data to the device named. Only Paletted RGBA PNG images are supported.
func WriteImage(device string, data []byte) error {
	// get screen dimensions, bit depth, text cell size, framebuffer padding
	// winsize := term.GetWinsize()
	// cellwidth := winsize.Xres / winsize.Cols
	// cellheight := winsize.Yres / winsize.Rows
	scrinfo := term.QueryFramebuffer(device)
	if scrinfo.Bpp != 32 {
		return fmt.Errorf("display must be 32 bits per pixel %v", scrinfo)
	}
	var pad int = int(scrinfo.Xresv) - int(scrinfo.Xres)
	if pad < 0 {
		return fmt.Errorf("virtual screen smaller than actual %v", scrinfo)
	}

	// decode PNG data
	reader := bytes.NewReader(data)
	img, err := png.Decode(reader)
	if err != nil {
		return err
	}

	// write image data to buffer
	var buf bytes.Buffer
	if img, ok := img.(*image.Paletted); ok {
		bounds := img.Bounds()
		offset := 0
		for i := 0; i < bounds.Dy(); i++ {
			for j := 0; j < bounds.Dx(); j++ {
				colorindx := img.Pix[offset]
				rgba := img.Palette[colorindx]
				binary.Write(&buf, binary.LittleEndian, rgba)
				offset++
			}
		}
	} else {
		return fmt.Errorf("image not in expected format")
	}

	// write buffer to framebuffer
	var fbinfo C.struct_fbinfo
	var imginfo C.struct_image
	fbinfo.xres = C.int(scrinfo.Xres)
	fbinfo.yres = C.int(scrinfo.Yres)
	fbinfo.pad = C.int(pad)
	fbinfo.device = C.CString("/dev/fb0")
	imginfo.xres = C.int(img.Bounds().Dx())
	imginfo.yres = C.int(img.Bounds().Dy())
	imgdata := buf.Bytes()
	imginfo.length = C.int(len(imgdata))
	if C.write_image(&imginfo, &fbinfo, (*C.char)(unsafe.Pointer(&imgdata[0]))) == 0 {
		return fmt.Errorf("failed to write data to framebuffer")
	}
	return nil
}
