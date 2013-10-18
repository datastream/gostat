package stat

// Inverse Gamma distribution (not to be confused with Inverse CDF of Gamma distribution)

import (
	"code.google.com/p/go-fn/fn"
	"math"
)

// InvGammaPDF is Inverse Gamma distribution: probability density function
func InvGammaPDF(a, b float64) func(x float64) float64 {
	return func(x float64) float64 {
		return math.Exp(a*math.Log(b) - fn.LnΓ(a) - (a+1)*math.Log(x) - b*1.0/x)
	}
}

// InvGammaLnPDF is Inverse Gamma distribution's natural logarithm of the probability density function
func InvGammaLnPDF(a, b float64) func(x float64) float64 {
	return func(x float64) float64 {
		return a*math.Log(b) - fn.LnΓ(a) - (a+1)*math.Log(x) - b*1.0/x
	}
}

// InvGammaPDFAt return Inverse Gamma distribution's probability density function at x
func InvGammaPDFAt(a, b float64) func(x float64) float64 {
	return func(x float64) float64 {
		return math.Exp(a*math.Log(b) - fn.LnΓ(a) - (a+1)*math.Log(x) - b*1.0/x)
	}
}

// InvGammaCDF is Inverse Gamma distribution's cumulative distribution function
func InvGammaCDF(a, b float64) func(x float64) float64 {
	return func(x float64) float64 {
		return 1 - fn.IΓ(a, b*1.0/x)
	}
}

// InvGammaCDFAt is Inverse Gamma distribution's value of the cumulative distribution function at x
func InvGammaCDFAt(a, b, x float64) float64 {
	cdf := InvGammaCDF(a, b)
	return cdf(x)
}
