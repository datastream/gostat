package stat

// BernoulliPMF is bernoulli distribution's Probability mass function
func BernoulliPMF(ρ float64) func(k int64) float64 {
	return func(k int64) float64 {
		if k < 0 || k > 1 {
			panic("k is not 0 or 1")
		}
		if k == 1 {
			return ρ
		}
		return 1 - ρ
	}
}

// BernoulliPMFAt is bernoulli distribution's Probability mass function for k
func BernoulliPMFAt(ρ float64, k int64) float64 {
	pmf := BernoulliPMF(ρ)
	return pmf(k)
}

// BernoulliLnPMF is bernoulli distribution's Probability mass function with log
func BernoulliLnPMF(ρ float64) func(k int64) float64 {
	return func(k int64) float64 {
		if k == 1 {
			return log(ρ)
		}
		return log(1 - ρ)
	}
}

// NextBernoulli return next value of bernoullil distribution
func NextBernoulli(ρ float64) int64 {
	if NextUniform() < ρ {
		return 1
	}
	return 0
}

// Bernoulli distribution function
func Bernoulli(ρ float64) func() int64 { return func() int64 { return NextBernoulli(ρ) } }

// BernoulliCDF is Bernoulli distribution's CDF
func BernoulliCDF(ρ float64) func(k int64) float64 {
	return func(k int64) float64 {
		if k < 0 || k > 1 {
			panic("k is not 0 or 1")
		}
		if k == 1 {
			return 1
		}
		return 1 - ρ
	}
}
