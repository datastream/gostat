package stat

import (
	"code.google.com/p/go-fn/fn"
)

// MultinomialPMF is multinomial distribution's pmf
func MultinomialPMF(θ []float64, n int64) func(x []int64) float64 {
	return func(x []int64) float64 {
		if len(x) != len(θ) {
			return 0
		}
		l := fOne
		totalx := iZero
		for i := 0; i < len(x); i++ {
			l *= pow(θ[i], float64(x[i]))
			l /= fn.Γ(float64(x[i] + 1))
			totalx += x[i]
		}
		if totalx != n {
			return 0
		}
		l *= fn.Γ(float64(totalx + 1))
		return l
	}
}

// MultinomialLnPMF is multinomial distribution's lnpmf
func MultinomialLnPMF(θ []float64, n int64) func(x []int64) float64 {
	return func(x []int64) float64 {
		if len(x) != len(θ) {
			return negInf
		}
		l := fZero
		totalx := iZero
		for i := 0; i < len(x); i++ {
			l += log(θ[i]) * float64(x[i])
			l -= fn.LnΓ(float64(x[i] + 1))
			totalx += x[i]
		}
		if totalx != n {
			return negInf
		}
		l += fn.LnΓ(float64(totalx + 1))
		return l
	}
}

// NextMultinomial return random value in Multinomial distribution
func NextMultinomial(θ []float64, n int64) []int64 {
	x := make([]int64, len(θ))
	chooser := Choice(θ)
	for i := iZero; i < n; i++ {
		x[chooser()]++
	}
	return x
}

// Multinomial is multinomial distribution function
func Multinomial(θ []float64, n int64) func() []int64 {
	return func() []int64 {
		return NextMultinomial(θ, n)
	}
}
