package plot

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
)

// FuncF64 is a function that takes a float64 and returns a float64
type FuncF64 func(float64) float64

// Plot is the intermediate plot format that will be rendered as ascii or a PNG
type Plot struct {
	Data      []int
	Ymin      float64
	Ymax      float64
	W         int
	H         int
	N         int
	LineColor uint32
	Dots      bool
}

// ID is the identity function for float64. It returns what it is given.
// Useful when you need a noop function.
func ID(x float64) float64 {
	return x
}

// Compose plots g(/f(x)) where /f(x) is the mean of f(x) values
// for that pixel of the plot.
func Compose(g, f FuncF64, x []float64, width, height int) *Plot {
	N := len(x)

	// resample to fit screen
	y := make([]float64, width)
	M := (N-1)/width + 1
	i := 0
	j := 0
	t := 0.0
	for _, xn := range x {
		t += f(xn)
		j++
		if j == M {
			y[i] = g(t / float64(M))
			i++
			j = 0
			t = 0.0
		}
	}
	if j > 0 {
		y[i] = g(t / float64(j))
	}
	actualW := i

	// rescale and plot
	ymin, ymax := minmax(y[:])
	data := make([]int, width)
	for n, yn := range y {
		data[n] = height - 1 - int((yn-ymin)/(ymax-ymin)*float64(height-1))
	}
	return &Plot{data, ymin, ymax, actualW, height, N, 0x00ff00ff, true}
}

// RenderImage renders the plot as a paletted image
func (plt *Plot) RenderImage() *image.Paletted {
	offset := 2
	imgwidth := plt.W + offset
	rect := image.Rect(0, 0, imgwidth, plt.H)
	palette := color.Palette([]color.Color{
		color.RGBA{0, 0, 0, 255},
		color.RGBA{
			uint8(plt.LineColor >> 24),
			uint8(plt.LineColor >> 16),
			uint8(plt.LineColor >> 8),
			uint8(plt.LineColor),
		},
	})
	img := image.NewPaletted(rect, palette)
	for i := 0; i < plt.H; i++ {
		for j := 0; j < min(plt.W, plt.N); j++ {
			y := plt.Data[j]
			if (plt.Dots && i == y) || (!plt.Dots && i >= y) {
				img.Pix[i*imgwidth+offset+j] = 1
			}
		}
	}
	return img
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

func minmax(xs []float64) (min float64, max float64) {
	min = math.MaxFloat64
	max = -math.MaxFloat64
	for _, x := range xs {
		if x > max {
			max = x
		}
		if x < min {
			min = x
		}
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
