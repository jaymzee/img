package plot

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
	LineColor uint8
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
	return &Plot{data, ymin, ymax, actualW, height, N, 2, true}
}
