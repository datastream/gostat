package stat

// Gamma distribution
// k > 0		shape parameter
// θ (Theta) > 0	scale parameter

import (
	"code.google.com/p/go-fn/fn"
	"fmt"
	"math"
)

/* did not pass test, so commented out
// GammaPDF is gamma distribution's Probability density function
func GammaPDF(α float64, λ float64) func(x float64) float64 {
	expPart := ExpPDF(λ)
	return func(x float64) float64 {
		if x < 0 {
			return 0
		}
		return expPart(x) * pow(λ*x, α-1) / Γ(α)
	}
}
*/

// GammaPDF is gamma distribution's Probability density function
func GammaPDF(k float64, θ float64) func(x float64) float64 {
	return func(x float64) float64 {
		if x < 0 {
			return 0
		}
		return pow(x, k-1) * exp(-x/θ) / (fn.Γ(k) * pow(θ, k))
	}
}

// GammaLnPDF is Natural logarithm of the probability density function
func GammaLnPDF(α float64, λ float64) func(x float64) float64 {
	expPart := ExpLnPDF(λ)
	return func(x float64) float64 {
		if x < 0 {
			return negInf
		}
		return expPart(x) + (α-1)*log(λ*x) - fn.LnΓ(α)
	}
}

// NextGamma return Random value drawn from the distribution
func NextGamma(α float64, λ float64) float64 {
	//if α is a small integer, this way is faster on my laptop
	if α == float64(int64(α)) && α <= 15 {
		x := NextExp(λ)
		for i := 1; i < int(α); i++ {
			x += NextExp(λ)
		}
		return x
	}

	if α < 0.75 {
		return RejectionSample(GammaPDF(α, λ), ExpPDF(λ), Exp(λ), 1)
	}

	//Tadikamalla ACM '73
	a := α - 1
	b := 0.5 + 0.5*sqrt(4*α-3)
	c := a * (1 + b) / b
	d := (b - 1) / (a * b)
	s := a / b
	p := 1.0 / (2 - exp(-s))
	var x, y float64
	for i := 1; ; i++ {
		u := NextUniform()
		if u > p {
			var e float64
			for e = -log((1 - u) / (1 - p)); e > s; e = e - a/b {
			}
			x = a - b*e
			y = a - x
		} else {
			x = a - b*log(u/p)
			y = x - a
		}
		u2 := NextUniform()
		if log(u2) <= a*log(d*x)-x+y/b+c {
			break
		}
	}
	return x / λ
}

// Gamma is gamma distribution function
func Gamma(α float64, λ float64) func() float64 {
	return func() float64 { return NextGamma(α, λ) }
}

// GammaCDF is Cumulative distribution function, analytic solution, did not pass some tests!
func GammaCDF(k float64, θ float64) func(x float64) float64 {
	return func(x float64) float64 {
		if k < 0 || θ < 0 {
			panic(fmt.Sprintf("k < 0 || θ < 0"))
		}
		if x < 0 {
			return 0
		}
		return fn.Iγ(k, x/θ) / fn.Γ(k)
	}
}

// GammaCDFint is Cumulative distribution function, for integer k only
func GammaCDFint(k int64, θ float64) func(x float64) float64 {
	return func(x float64) float64 {
		if k < 0 || θ < 0 {
			panic(fmt.Sprintf("k < 0 || θ < 0"))
		}
		if x < 0 {
			return 0
		}
		return fn.Iγint(k, x/θ) / fn.Γ(float64(k))
	}
}

/*
// Cumulative distribution function, using gamma incomplete integral  DOES NOT WORK !!!
func GammaCDF(k float64, θ float64) func(x float64) float64 {
	return func(x float64) float64 {
		if k < 0 || θ < 0 {
			panic(fmt.Sprintf("k < 0 || θ < 0"))
		}
		if x < 0 {
			return 0
		}
		return IGam(θ, k*x)
	}
}
*/

// GammaPDFAt return Value of the probability density function at x
func GammaPDFAt(k, θ, x float64) float64 {
	pdf := GammaPDF(k, θ)
	return pdf(x)
}

// GammaCDFAt return Value of the cumulative distribution function at x
func GammaCDFAt(k, θ, x float64) float64 {
	cdf := GammaCDF(k, θ)
	return cdf(x)
}

// GammaInvCDF is Inverse CDF (Quantile) function
func GammaInvCDF(k float64, θ float64) func(x float64) float64 {
	return func(x float64) float64 {
		var eps, yNew, h float64
		eps = 1e-4
		y := k * θ
		yOld := y
	L:
		for i := 0; i < 100; i++ {
			h = (GammaCDFAt(k, θ, yOld) - x) / GammaPDFAt(k, θ, yOld)
			yNew = yOld - h
			if yNew <= eps {
				yNew = yOld / 10
				h = yOld - yNew
			}
			if math.Abs(h) < eps {
				break L
			}
			yOld = yNew
		}
		return yNew
	}
}

// GammaInvCDFFor return Value of the inverse CDF for probability p
func GammaInvCDFFor(k, θ, p float64) float64 {
	cdf := GammaInvCDF(k, θ)
	return cdf(p)
}
