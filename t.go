package stat

import (
	"code.google.com/p/go-fn/fn"
)

// StudentsTPDF is students' T distribution's pdf
func StudentsTPDF(ν float64) func(x float64) float64 {
	normalization := fn.Γ((ν+1)/2) / (sqrt(ν*π) * fn.Γ(ν/2))
	return func(x float64) float64 {
		return normalization * pow(1+x*x/ν, -(ν+1)/2)
	}
}

// StudentsTLnPDF is students' T distribution's ln pdf
func StudentsTLnPDF(ν float64) func(x float64) float64 {
	normalization := fn.LnΓ((ν+1)/2) - log(sqrt(ν*π)) - fn.LnΓ(ν/2)
	return func(x float64) float64 {
		return normalization + log(1+x*x/ν)*-(ν+1)/2
	}
}

// NextStudentsT  N(0, 1)*sqrt(ν/NextGamma(ν/2, 2))
func NextStudentsT(ν float64) float64 {
	return NextNormal(0, 1) * sqrt(ν/NextGamma(ν/2, 2))
}

// StudentsT is students' T distribution function
func StudentsT(ν float64) func() float64 {
	return func() float64 {
		return NextStudentsT(ν)
	}
}
