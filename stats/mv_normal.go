package stat

import (
	matrix "github.com/skelterjohn/go.matrix"
)

// MVNormalPDF is move normal distribution's pdf
func MVNormalPDF(μ *matrix.DenseMatrix, Σ *matrix.DenseMatrix) func(x *matrix.DenseMatrix) float64 {
	p := μ.Rows()
	backμ := μ.DenseMatrix()
	backμ.Scale(-1)

	Σdet := Σ.Det()
	ΣdetRt := sqrt(Σdet)
	Σinv, _ := Σ.Inverse()

	normalization := pow(2*π, -float64(p)/2) / ΣdetRt

	return func(x *matrix.DenseMatrix) float64 {
		δ, _ := x.PlusDense(backμ)
		tmp := δ.Transpose()
		tmp, _ = tmp.TimesDense(Σinv)
		tmp, _ = tmp.TimesDense(δ)
		f := tmp.Get(0, 0)
		return normalization * exp(-f/2)
	}
}

// NextMVNormal return random value in move normal distribution
func NextMVNormal(μ *matrix.DenseMatrix, Σ *matrix.DenseMatrix) *matrix.DenseMatrix {
	n := μ.Rows()
	x := matrix.Zeros(n, 1)
	for i := 0; i < n; i++ {
		x.Set(i, 0, NextNormal(0, 1))
	}
	C, err := Σ.Cholesky()
	Cx, err := C.TimesDense(x)
	μCx, err := μ.PlusDense(Cx)
	if err != nil {
		panic(err)
	}
	return μCx
}

// MVNormal is move normal distribution
func MVNormal(μ *matrix.DenseMatrix, Σ *matrix.DenseMatrix) func() *matrix.DenseMatrix {
	C, _ := Σ.Cholesky()
	n := μ.Rows()
	return func() *matrix.DenseMatrix {
		x := matrix.Zeros(n, 1)
		for i := 0; i < n; i++ {
			x.Set(i, 0, NextNormal(0, 1))
		}
		Cx, _ := C.TimesDense(x)
		MCx, _ := μ.PlusDense(Cx)
		return MCx
	}
}
