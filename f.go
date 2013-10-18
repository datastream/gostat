package stat

// F-distribution, alias Fisher-Snedecor distribution
import (
	"code.google.com/p/go-fn/fn"
	"fmt"
)

// FPDF is F-distribution's PDF
func FPDF(d1 float64, d2 float64) func(x float64) float64 {
	normalization := 1 / fn.B(d1/2, d2/2)
	return func(x float64) float64 {
		return normalization * sqrt(pow(d1*x, d1)*pow(d2, d2)/pow(d1*x+d2, d1+d2)) / x
	}
}

// FLnPDF is F-distribution's lnPDF
func FLnPDF(d1 float64, d2 float64) func(x float64) float64 {
	normalization := -fn.LnB(d1/2, d2/2)
	return func(x float64) float64 {
		return normalization + log(d1*x)*d1/2 + log(d2)*d2/2 - log(d1*x+d2)*(d1+d2)/2 - log(x)
	}
}

// NextF return value in F-distribution
func NextF(d1 int64, d2 int64) float64 {
	return (NextXsquare(d1) * float64(d2)) / (NextXsquare(d2) * float64(d1))
}

// F is F-distribution function
func F(d1 int64, d2 int64) func() float64 {
	return func() float64 {
		return NextF(d1, d2)
	}
}

// FCDF is F-distribution's CDF
func FCDF(df1, df2 float64) func(x float64) float64 {
	return func(x float64) float64 {
		y := df1 * x / (df1*x + df2)
		return fn.BetaIncReg(df1/2.0, df2/2.0, y)
	}
}

// FCDFAt return  F-distribution's CDF at x
func FCDFAt(df1, df2, x float64) float64 {
	cdf := FCDF(df1, df2)
	return cdf(x)
}

// FInvCDF return Inverse CDF (Quantile) function of F-distribution
func FInvCDF(df1, df2 float64) func(p float64) float64 {
	return func(p float64) float64 {
		if p < 0.0 {
			panic(fmt.Sprintf("p < 0"))
		}
		if p > 1.0 {
			panic(fmt.Sprintf("p > 1.0"))
		}
		if df1 < 1.0 {
			panic(fmt.Sprintf("df1 < 1"))
		}
		if df2 < 1.0 {
			panic(fmt.Sprintf("df2 < 1"))
		}

		return ((1/BetaInvCDFFor(df2/2, df1/2, 1-p) - 1) * df2 / df1)
	}
}

// FInvCDFFor return Value of the inverse CDF of F-distribution for probability p
func FInvCDFFor(df1, df2, p float64) float64 {
	cdf := FInvCDF(df1, df2)
	return cdf(p)
}
