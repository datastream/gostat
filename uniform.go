// Uniform (Flat) distribution

package stat

import (
	"math/rand"
)

// UniformPDF is uniform distribution's pdf
func UniformPDF() func(x float64) float64 {
	return func(x float64) float64 {
		if 0 <= x && x <= 1 {
			return 1
		}
		return 0
	}
}

// UniformLnPDF is uniform distribution's lnpdf
func UniformLnPDF() func(x float64) float64 {
	return func(x float64) float64 {
		if 0 <= x && x <= 1 {
			return 0
		}
		return negInf
	}
}

// NextUniform is rand float64
var NextUniform = rand.Float64

// Uniform is uniform distribution function
func Uniform() func() float64 { return NextUniform }
