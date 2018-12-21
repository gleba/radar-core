package ux

import "math"

func Round2zero(x float64) float64 {
	return math.Ceil(x*100) / 100
}

func Round4zero(x float64) float64 {
	return math.Ceil(x*10000) / 10000
}
