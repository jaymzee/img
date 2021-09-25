package plot

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
)

// RenderImage renders the plot as a paletted image
func (plt *Plot) RenderImage() *image.Paletted {
	offset := 100
	imgwidth := plt.W + offset
	rect := image.Rect(0, 0, imgwidth, plt.H)
	palette := color.Palette([]color.Color{
		color.RGBA{0, 0, 0, 255},       // black
		color.RGBA{255, 0, 0, 255},     // red
		color.RGBA{0, 255, 0, 255},     // green
		color.RGBA{255, 255, 0, 255},   // yellow
		color.RGBA{0, 0, 255, 255},     // blue
		color.RGBA{255, 0, 255, 255},   // magenta
		color.RGBA{0, 255, 255, 255},   // cyan
		color.RGBA{255, 255, 255, 255}, // white
	})
	img := image.NewPaletted(rect, palette)
	for i := 0; i < plt.H; i++ {
		for j := 0; j < min(plt.W, plt.N); j++ {
			y := plt.Data[j]
			if (plt.Dots && i == y) || (!plt.Dots && i >= y) {
				img.Pix[i*imgwidth+offset+j] = plt.LineColor
			}
		}
	}
	AddLabel(img, fmt.Sprintf("%11.3e", plt.Ymax), 0, 0, 7)
	AddLabel(img, fmt.Sprintf("%11.3e", plt.Ymin), 0, plt.H-8, 7)
	return img
}

// AddLabel adds a text label to a paletted image in a simple 8x8 font
func AddLabel(img *image.Paletted, s string, x, y int, color uint8) {
	width := img.Bounds().Dx()
	data := []byte(s)
	for _, char := range data {
		c := char & 0x7f // strip off 8th bit
		for i := 0; i < 8; i++ {
			d := font8x8basic[c][i]
			for j := 0; j < 8; j++ {
				offset := x + j + (y+i)*width
				if d&1 == 1 {
					img.Pix[offset] = color
				}
				d = d >> 1
			}
		}
		x += 8
	}
}

// RenderASCII renders the plot as ascii art
func (plt *Plot) RenderASCII() string {
	buf := new(bytes.Buffer)

	for i := 0; i < plt.H; i++ {
		if i == 0 {
			fmt.Fprintf(buf, "\n%11.3e |", plt.Ymax)
		} else if i == plt.H-1 {
			fmt.Fprintf(buf, "\n%11.3e |", plt.Ymin)
		} else {
			fmt.Fprintf(buf, "\n            |")
		}
		for j := 0; j < min(plt.W, plt.N); j++ {
			if plt.Data[j] == i {
				fmt.Fprint(buf, "*")
			} else {
				fmt.Fprint(buf, " ")
			}
		}
	}
	fmt.Fprintln(buf)

	return buf.String()
}

// RenderPNG renders the plot as a PNG image
func (plt *Plot) RenderPNG() []byte {
	var buf bytes.Buffer
	enc := png.Encoder{CompressionLevel: png.BestSpeed}
	err := enc.Encode(&buf, plt.RenderImage())
	if err != nil {
		panic(err)
	}
	return buf.Bytes()
}
