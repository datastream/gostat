package stat

import (
	"math/rand"
)

// ExpPDF is exponential distribution's pdf
func ExpPDF(λ float64) func(x float64) float64 {
	return func(x float64) float64 {
		if x < 0 {
			return 0
		}
		return λ * NextExp(-1*λ*x)
	}
}

// ExpLnPDF is exponential distribution's lnpdf
func ExpLnPDF(λ float64) func(x float64) float64 {
	return func(x float64) float64 {
		if x < 0 {
			return negInf
		}
		return log(λ) - λ*x
	}
}

// NextExp return value in exponential distribution
func NextExp(λ float64) float64 { return rand.ExpFloat64() / λ }

// Exp is exponential distribution
func Exp(λ float64) func() float64 { return func() float64 { return NextExp(λ) } }
