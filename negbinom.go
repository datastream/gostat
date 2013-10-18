package stat

// Negative Binomial distribution

import (
	"code.google.com/p/go-fn/fn"
	"math"
)

/*
// Does not pass the test, so commented out
func NegativeBinomialPMF(ρ float64, r int64) func(i int64) float64 {
	return func(k int64) float64 {
		return float64(Choose(k+r-1, r-1)) * pow(ρ, float64(r)) * pow(1-ρ, float64(k))
	}
}
*/

// NegativeBinomialPMF is Negative Binomial distribution's pmf
func NegativeBinomialPMF(ρ float64, r int64) func(k int64) float64 {
	return func(k int64) float64 {
		return fn.BinomCoeff(k+r-1, k) * math.Pow(1-ρ, float64(r)) * math.Pow(ρ, float64(k))
	}
}

// NegativeBinomialPMFAt return Negative Binomial distribution PMF at k
func NegativeBinomialPMFAt(ρ float64, r, k int64) float64 {
	pmf := NegativeBinomialPMF(ρ, r)
	return pmf(k)
}

// NegativeBinomialLnPMF is Negative Binomial distribution's LnPMF
func NegativeBinomialLnPMF(ρ float64, r int64) func(i int64) float64 {
	return func(k int64) float64 {
		return fn.LnChoose(k+r-1, r-1) + log(ρ)*float64(r) + log(1-ρ)*float64(k)
	}
}

// NextNegativeBinomial return value in negative binomial distribution
func NextNegativeBinomial(ρ float64, r int64) int64 {
	k := iZero
	for r >= 0 {
		i := NextBernoulli(ρ)
		r -= i
		k += (1 - i)
	}
	return k
}

// NegativeBinomial  => number of NextBernoulli(ρ) failures before r successes
func NegativeBinomial(ρ float64, r int64) func() int64 {
	return func() int64 {
		return NextNegativeBinomial(ρ, r)
	}
}

// NegativeBinomialCDF is negative binomial distribution's cdf
func NegativeBinomialCDF(ρ float64, r int64) func(k int64) float64 {
	return func(k int64) float64 {
		IP := BetaCDFAt(float64(k+1), float64(r), ρ)
		return 1 - IP
	}
}

// NegativeBinomialCDFAt return negative binomial distribution's cdf at k
func NegativeBinomialCDFAt(ρ float64, r, k int64) float64 {
	cdf := NegativeBinomialCDF(ρ, r)
	return cdf(k)
}
