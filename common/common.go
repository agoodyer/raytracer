package common

import (
	"math"
)

const (
	Infinity = math.MaxFloat64
	Pi       = 3.1415926535897932385
)

// const infinity float64 = math.MaxFloat64

// const pi float64 = 3.1415926535897932385

func degrees_to_radians(degrees float64) float64 {
	return degrees * Pi / 180.0
}