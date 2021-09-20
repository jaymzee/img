package fb

import (
	"bytes"
	"fmt"
	"github.com/jaymzee/img/term"
	"image"
	"image/png"
	"io"
	"os"
)

// WriteImageToFramebuffer takes a png image and Writes the raw RGBA pixel
// data to the device named. Only RGBA png images are supported.
func WriteImage(device string, data []byte) error {
	// get screen dimensions, bit depth, text cell size, framebuffer padding
	winsize := term.GetWinsize()
	// cellwidth := winsize.Xres / winsize.Cols
	// cellheight := winsize.Yres / winsize.Rows
	scrinfo := term.QueryFramebuffer(device)
	if scrinfo.Xres != uint(winsize.Xres) || scrinfo.Yres != uint(winsize.Yres) || scrinfo.Bpp < 1 {
		return fmt.Errorf("bad ScreenInfo %v %v", scrinfo, winsize)
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

	// write image data to framebuffer
	if img, ok := img.(*image.RGBA); ok {
		bounds := img.Bounds()
		width := bounds.Dx()
		height := bounds.Dy()
		pix := img.Pix
		f, err := os.OpenFile(device, os.O_RDWR, 0644)
		if err != nil {
			return err
		}
		defer f.Close()
		offset := 0
		for j := 0; j < width; j++ {
			for i := 0; i < height; i++ {
				f.Write(pix[offset : offset+4])
				offset++
			}
			f.Seek(io.SeekCurrent, pad)
		}
	} else {
		return fmt.Errorf("image not in expected format")
	}

	return nil
}
