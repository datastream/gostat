package stat

// GeometricPMF is geometric distribution's pmf
func GeometricPMF(ρ float64) func(i int64) float64 {
	return func(n int64) float64 { return ρ * pow(ρ, float64(n)) }
}

// GeometricLnPMF is geometric distribution's  lnpmf
func GeometricLnPMF(ρ float64) func(i int64) float64 {
	return func(n int64) float64 { return log(ρ) + float64(n)*log(ρ) }
}

// NextGeometric => # of NextBernoulli(ρ) failures before one success
func NextGeometric(ρ float64) int64 {
	if NextBernoulli(ρ) == 1 {
		return 1 + NextGeometric(ρ)
	}
	return 0
}

// Geometric is geometric distribution funtion
func Geometric(ρ float64) func() int64 { return func() int64 { return NextGeometric(ρ) } }
