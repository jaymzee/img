package plot

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"math"
)

type FuncF64 func(float64) float64

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

func Id(x float64) float64 {
	return x
}

func PlotFunc(x []float64, f, g FuncF64, W, H int) *Plot {
	N := len(x)

	// resample to fit screen
	y := make([]float64, W)
	M := (N-1)/W + 1
	i := 0
	j := 0
	t := 0.0
	for _, xn := range x {
		t += g(xn)
		j++
		if j == M {
			y[i] = f(t / float64(M))
			i++
			j = 0
			t = 0.0
		}
	}
	if j > 0 {
		y[i] = f(t / float64(j))
	}
	actualW := i

	// rescale and plot
	ymin, ymax := minmax(y[:])
	data := make([]int, W)
	for n, yn := range y {
		data[n] = H - 1 - int((yn-ymin)/(ymax-ymin)*float64(H-1))
	}
	return &Plot{data, ymin, ymax, actualW, H, N, 0x00ff00ff, true}
}

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

func (plt *Plot) RenderAscii() string {
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
	} else {
		return b
	}
}
