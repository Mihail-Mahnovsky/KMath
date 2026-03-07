package kmath

import "math"

func Sygmoid(inputV float64) float64 {
	return 1.0 / (1.0 + math.Exp(inputV))
}
