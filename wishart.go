package stat

import (
	"code.google.com/p/go-fn/fn"
	m "github.com/skelterjohn/go.matrix"
)

// WishartPDF is wishart distribution's pdf
func WishartPDF(n int, V *m.DenseMatrix) func(W *m.DenseMatrix) float64 {
	p := V.Rows()
	Vdet := V.Det()
	Vinv, _ := V.Inverse()
	normalization := pow(2, -0.5*float64(n*p)) *
		pow(Vdet, -0.5*float64(n)) /
		fn.Γ(0.5*float64(n))
	return func(W *m.DenseMatrix) float64 {
		VinvW, _ := Vinv.Times(W)
		return normalization * pow(W.Det(), 0.5*float64(n-p-1)) *
			exp(-0.5*VinvW.Trace())
	}
}

// WishartLnPDF is wishart distribution's lnpdf
func WishartLnPDF(n int, V *m.DenseMatrix) func(W *m.DenseMatrix) float64 {

	p := V.Rows()
	Vdet := V.Det()
	Vinv, _ := V.Inverse()
	normalization := log(2)*(-0.5*float64(n*p)) +
		log(Vdet)*(-0.5*float64(n)) -
		fn.LnΓ(0.5*float64(n))
	return func(W *m.DenseMatrix) float64 {
		VinvW, _ := Vinv.Times(W)
		return normalization +
			log(W.Det())*0.5*float64(n-p-1) -
			0.5*VinvW.Trace()
	}
}

// NextWishart return metric in wishart distribution
func NextWishart(n int, V *m.DenseMatrix) *m.DenseMatrix {
	return Wishart(n, V)()
}

// Wishart is wishart distribution function
func Wishart(n int, V *m.DenseMatrix) func() *m.DenseMatrix {
	p := V.Rows()
	zeros := m.Zeros(p, 1)
	rowGen := MVNormal(zeros, V)
	return func() *m.DenseMatrix {
		x := make([][]float64, n)
		for i := 0; i < n; i++ {
			x[i] = rowGen().Array()
		}
		X := m.MakeDenseMatrixStacked(x)
		S, _ := X.Transpose().TimesDense(X)
		return S
	}
}

// InverseWishartPDF is inverse wishart distribution's pdf
func InverseWishartPDF(n int, Ψ *m.DenseMatrix) func(B *m.DenseMatrix) float64 {
	p := Ψ.Rows()
	Ψdet := Ψ.Det()
	normalization := pow(Ψdet, -0.5*float64(n)) *
		pow(2, -0.5*float64(n*p)) /
		fn.Γ(float64(n)/2)
	return func(B *m.DenseMatrix) float64 {
		Bdet := B.Det()
		Binv, _ := B.Inverse()
		ΨBinv, _ := Ψ.Times(Binv)
		return normalization *
			pow(Bdet, -.5*float64(n+p+1)) *
			exp(-0.5*ΨBinv.Trace())
	}
}

// InverseWishartLnPDF is inverse wishart distribution's lnpdf
func InverseWishartLnPDF(n int, Ψ *m.DenseMatrix) func(W *m.DenseMatrix) float64 {
	p := Ψ.Rows()
	Ψdet := Ψ.Det()
	normalization := log(Ψdet)*-0.5*float64(n) +
		log(2)*-0.5*float64(n*p) -
		fn.LnΓ(float64(n)/2)
	return func(B *m.DenseMatrix) float64 {
		Bdet := B.Det()
		Binv, _ := B.Inverse()
		ΨBinv, _ := Ψ.Times(Binv)
		return normalization +
			log(Bdet)*-.5*float64(n+p+1) +
			-0.5*ΨBinv.Trace()
	}
}

// NextInverseWishart return metric in inverse wishart distribution
func NextInverseWishart(n int, V *m.DenseMatrix) *m.DenseMatrix {
	return InverseWishart(n, V)()
}

// InverseWishart is inverse wishart distribution function
func InverseWishart(n int, V *m.DenseMatrix) func() *m.DenseMatrix {
	p := V.Rows()
	zeros := m.Zeros(p, 1)
	rowGen := MVNormal(zeros, V)
	return func() *m.DenseMatrix {
		x := make([][]float64, n)
		for i := 0; i < n; i++ {
			x[i] = rowGen().Array()
		}
		X := m.MakeDenseMatrixStacked(x)
		S, _ := X.Transpose().TimesDense(X)
		Sinv, _ := S.Inverse()
		return Sinv
	}
}
