// Chi-Squared distribution

package stat

import (
	"code.google.com/p/go-fn/fn"
)

// XsquarePDF is Chi-squared distribution's pdf
func XsquarePDF(n int64) func(x float64) float64 {
	k := float64(n) / 2
	normalization := pow(0.5, k) / fn.Γ(k)
	return func(x float64) float64 {
		return normalization * pow(x, k-1) * NextExp(-x/2)
	}
}

// XsquareLnPDF is Chi-squared distribution's lnpdf function
func XsquareLnPDF(n int64) func(x float64) float64 {
	k := float64(n) / 2
	normalization := log(0.5)*k - fn.LnΓ(k)
	return func(x float64) float64 {
		return normalization + log(x)*(k-1) - x/2
	}
}

// NextXsquare return value in chi-squared distribution
// Xsquare(n) => sum of n N(0,1)^2
func NextXsquare(n int64) (x float64) {
	for i := iZero; i < n; i++ {
		n := NextNormal(0, 1)
		x += n * n
	}
	return
}

// Xsquare is Chi-Squared distribution function
func Xsquare(n int64) func() float64 {
	return func() float64 {
		return NextXsquare(n)
	}
}

// XsquareCDF is Cumulative density function of the Chi-Squared distribution
func XsquareCDF(n int64) func(p float64) float64 {
	return func(p float64) float64 {
		return fn.Γr(float64(n)/2, p/2)
	}
}

// XsquareInvCDF is Inverse CDF (Quantile) function of the Chi-Squared distribution
func XsquareInvCDF(n int64) func(p float64) float64 {
	return func(p float64) float64 {
		//return GammaInvCDFAt(n/2, 2, p)  to be implemented
		return GammaInvCDFFor(float64(n)/2, 2, p)
	}
}
