package stat

import (
	"math/rand"
)

// RangePMF is range distribution's pmf
func RangePMF(n int64) func(i int64) float64 {
	return func(i int64) float64 {
		return fOne / float64(n)
	}
}

// LnRangePMF is log range distribution's pmf
func LnRangePMF(n int64) func(i int64) float64 {
	return func(i int64) float64 {
		return -log(float64(n))
	}
}

// NextRange reutn value in range distribution
func NextRange(n int64) int64 {
	return rand.Int63n(n)
}

// Range is range distribution function
func Range(n int64) func() int64 {
	return func() int64 {
		return NextRange(n)
	}
}
