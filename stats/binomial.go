package stat

import (
	"code.google.com/p/go-fn/fn"
)

// BinomialPMF is Binomial distribution's Probability Mass Function
func BinomialPMF(ρ float64, n int64) func(i int64) float64 {
	return func(i int64) float64 {
		p := pow(ρ, float64(i)) * pow(1-ρ, float64(n-i))
		p *= fn.Γ(float64(n+1)) / (fn.Γ(float64(i+1)) * fn.Γ(float64(n-i+1)))
		return p
	}
}

// BinomialPMFAt return value of Binomial distribution's Probability Mass Function at k
func BinomialPMFAt(ρ float64, n, k int64) float64 {
	pmf := BinomialPMF(ρ, n)
	return pmf(k)
}

// BinomialLnPMF is  Natural logarithm of Probability Mass Function for the Binomial distribution
func BinomialLnPMF(ρ float64, n int64) func(i int64) float64 {
	return func(i int64) float64 {
		p := log(ρ)*float64(i) + log(1-ρ)*float64(n-i)
		p += fn.LnΓ(float64(n+1)) - fn.LnΓ(float64(i+1)) - fn.LnΓ(float64(n-i+1))
		return p
	}
}

// NextBinomial return value in binomial distribution
func NextBinomial(ρ float64, n int64) (result int64) {
	for i := int64(0); i <= n; i++ {
		result += NextBernoulli(ρ)
	}
	return
}

// Binomial is binomial distribution function
func Binomial(ρ float64, n int64) func() int64 {
	return func() int64 { return NextBinomial(ρ, n) }
}

// BinomialCDFtrivial is Cumulative Distribution Function for the Binomial distribution, trivial implementation
func BinomialCDFtrivial(ρ float64, n int64) func(k int64) float64 {
	return func(k int64) float64 {
		var p float64
		var i int64
		pmf := BinomialPMF(ρ, n)
		for i = 0; i <= k; i++ {
			p += pmf(i)
		}
		return p
	}
}

// BinomialCDF is Cumulative Distribution Function for the Binomial distribution
func BinomialCDF(ρ float64, n int64) func(k int64) float64 {
	return func(k int64) float64 {
		p := BetaCDFAt((float64)(n-k), (float64)(k+1), 1-ρ)
		return p
	}
}

// BinomialCDFAt return value of binomial distribution's Cumulative Distribution Function at k
func BinomialCDFAt(ρ float64, n, k int64) float64 {
	cdf := BinomialCDF(ρ, n)
	return cdf(k)
}
