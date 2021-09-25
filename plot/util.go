package plot

import "math"

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
