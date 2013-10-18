// Poisson distribution

package stat

import (
	"code.google.com/p/go-fn/fn"
	"math"
)

// PoissonLnPMF is poisson distribution's lnpmf
/*
func PoissonLnPMF(λ float64) (foo func(i int64) float64) {
	pmf := PoissonPMF(λ)
	return func(i int64) (p float64) {
		return log(pmf(i))
		//p = -λ +log(λ)*float64(i)
		//x := log(fn.Γ(float64(i)+1))
		//_ = x
		//p -= fn.LnΓ(float64(i)+1)
		//return p
	}
}
*/
func PoissonLnPMF(λ float64) func(k int64) float64 {
	return func(k int64) (p float64) {
		i := float64(k)
		a := log(λ) * i
		b := log(fn.Γ(i + 1))
		p = a - b - λ
		return p
	}
}

/*
func PoissonPMF(λ float64) func(k int64) float64 {
	return func(k int64) float64 {
		p := NextExp(-λ) * pow(λ, float64(k)) / Γ(float64(k)+1)
		return p
	}
}

func PoissonPMF(λ float64) func(k int64) float64 {
	return func(k int64) float64 {
		p := math.Exp(-λ) * pow(λ, float64(k)) / Γ(float64(k)+1)
		return p
	}
}
*/

// PoissonPMF is poisson distribution's pmf
func PoissonPMF(λ float64) func(k int64) float64 {
	pmf := PoissonLnPMF(λ)
	return func(k int64) float64 {
		p := math.Exp(pmf(k))
		return p
	}
}

// PoissonPMFAt is poisson distribution's pmf at k
func PoissonPMFAt(λ float64, k int64) float64 {
	pmf := PoissonPMF(λ)
	return pmf(k)
}

// NextPoisson return value in poisson distribution
func NextPoisson(λ float64) int64 {
	// this can be improved upon
	i := iZero
	t := exp(-λ)
	p := fOne
	for ; p > t; p *= NextUniform() {
		i++
	}
	return i
}

// Poisson is poisson distribution function
func Poisson(λ float64) func() int64 {
	return func() int64 {
		return NextPoisson(λ)
	}
}

// PoissonCDF is poisson distribution's cdf
func PoissonCDF(λ float64) func(k int64) float64 {
	return func(k int64) float64 {
		var p float64
		var i int64
		pmf := PoissonPMF(λ)
		for i = 0; i <= k; i++ {
			p += pmf(i)
		}
		return p
	}
}

// PoissonCDFA is poisson distribution's cdf analytic solution
func PoissonCDFA(λ float64) func(k int64) float64 { // analytic solution, less precision
	return func(k int64) float64 {
		p := math.Exp(math.Log(fn.IΓint(k+1, λ)) - (fn.LnFact(float64(k))))
		return p
	}
}

// PoissonCDFAt return poisson distribution's cdf at k
func PoissonCDFAt(λ float64, k int64) float64 {
	cdf := PoissonCDF(λ)
	return cdf(k)
}

// LnPoissonCDFA is ln poisson distribution cdf analytic solution
func LnPoissonCDFA(λ float64) func(k int64) float64 { // analytic solution, less precision
	return func(k int64) float64 {
		k1 := (float64)(k + 1)
		return log(fn.IΓ(k1, λ)) - fn.LnFact(float64(k))
	}
}
