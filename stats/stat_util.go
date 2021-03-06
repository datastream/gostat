package stat

import (
	"math"
)

var fZero float64
var fOne = float64(1.0)
var iZero int64
var iOne = int64(1)

var negInf = math.Inf(-1)

var log = math.Log
var exp = math.Exp
var sqrt = math.Sqrt
var pow = math.Pow

const π = float64(math.Pi)

// RejectionSample rejectionsample
func RejectionSample(targetDensity func(float64) float64, sourceDensity func(float64) float64, source func() float64, K float64) float64 {
	x := source()
	for ; NextUniform() >= targetDensity(x)/(K*sourceDensity(x)); x = source() {

	}
	return x
}

// ShuffleInt64 function
func ShuffleInt64(x []int64) {
	n := int64(len(x))
	for i := iZero; i < n; i++ {
		j := i + NextRange(n-i)
		t := x[i]
		x[i] = x[j]
		x[j] = t
	}
}

// ShuffleFloat64 function
func ShuffleFloat64(x []float64) {
	n := int64(len(x))
	for i := iZero; i < n; i++ {
		j := i + NextRange(n-i)
		t := x[i]
		x[i] = x[j]
		x[j] = t
	}
}

// Shuffle function
func Shuffle(x []interface{}) {
	n := int64(len(x))
	for i := iZero; i < n; i++ {
		j := i + NextRange(n-i)
		t := x[i]
		x[i] = x[j]
		x[j] = t
	}
}

func maxFloat64(x []float64) float64 {
	first := x[0]
	if len(x) > 1 {
		rest := maxFloat64(x[1:len(x)])
		if rest > first {
			first = rest
		}
	}
	return first
}
func maxInt64(x []int64) int64 {
	first := x[0]
	if len(x) > 1 {
		rest := maxInt64(x[1:len(x)])
		if rest > first {
			first = rest
		}
	}
	return first
}

func copyInt64(x []int64, n int64) []int64 {
	newx := make([]int64, n)
	for i := 0; i < len(x) && i < int(n); i++ {
		newx[i] = x[i]
	}
	return newx
}

func copyFloat64(x []float64, n int64) []float64 {
	newx := make([]float64, n)
	for i := 0; i < len(x) && i < int(n); i++ {
		newx[i] = x[i]
	}
	return newx
}
