package smp

// Sampling models.
// Performs sampling of species populations given population densities and sample size (sampling intensity).

import (
	"math"
	"code.google.com/p/go-eco/eco/aux"
	"code.google.com/p/probab/dst"
)

// betapara estimates the parameters a, b of beta distribution from expected proportion (pi), binomial denominator (m), and shape parameter (tau2). 
func betapara(pi, m, tau2 float64) (a, b float64) {
	// Solution (hopefully correct) of Exercise 4.17 of McCullagh & Nelder 1989, helped by Moore, Appl Stat 36, 8-14; 1987.
	t1 := tau2 * m
	t2 := t1 - m - tau2 + 1
	t3 := 1 / (1 + t1 - tau2)
	t4 := t2 * t3
	a = -t4 * pi
	b = t4 * (pi - 1)
	return
}

// PoissonSampler implements Poisson in-place sampling of the matrix of species abundances.
func PoissonSampler(mtx *aux.Matrix, sEffort float64) {
	var y float64
	nSamp := mtx.R
	nSpec := mtx.C
	for i := 0; i < nSamp; i++ {
		for j := 0; j < nSpec; j++ {
			x := sEffort * mtx.Get(i, j)
			if x > 0 {
				y = float64(dst.PoissonNext(x))
			} else {
				y = 0
			}
			mtx.Set(i, j, y)
		}
	}
}

// MultinomialSampler implements Multinomial in-place sampling of the matrix of species abundances.
func MultinomialSampler(mtx *aux.Matrix, sSize int64) {
	nSamp := mtx.R
	nSpec := mtx.C
	for i := 0; i < nSamp; i++ { // for every sample

		// recalculate species abundances to proportions
		θ := make([]float64, nSpec)
		sum := 0.0
		for j := 0; j < nSpec; j++ {
			sum += mtx.Get(i, j)
		}
		for j := 0; j < nSpec; j++ {
			θ[j] = mtx.Get(i, j) / sum
		}

		y := dst.MultinomialNext(θ, sSize)
		for j := 0; j < nSpec; j++ {
			mtx.Set(i, j, float64(y[j]))
		}
	}
}

// BetaBinomialSampler implements BetaBinomial in-place sampling of the matrix of species abundances.
func BetaBinomialSampler(mtx *aux.Matrix, sSize int64, shape []float64) {
	nSamp := mtx.R
	nSpec := mtx.C
	for i := 0; i < nSamp; i++ { // for every sample

		// recalculate species abundances to proportions
		θ := make([]float64, nSpec)
		sum := 0.0
		for j := 0; j < nSpec; j++ {
			sum += mtx.Get(i, j)
		}
		for j := 0; j < nSpec; j++ {
			θ[j] = mtx.Get(i, j) / sum
		}

		y := dst.MultinomialNext(θ, sSize)

		for j := 0; j < nSpec; j++ {
			x := float64(y[j])
			resp := 0.0
			if x > 0.0 {
				a, b := betapara(x, float64(sSize), shape[j])
				x = dst.BetaNext(a, b)
				resp = float64(dst.BinomialNext(sSize, x))
			}
			mtx.Set(i, j, resp)
		}
	}
}

// NegBinomSampler implements Poisson in-place sampling of the matrix of species abundances.
func NegBinomSampler(mtx *aux.Matrix, sSize int64, shape []float64) {
	nSamp := mtx.R
	nSpec := mtx.C
	for i := 0; i < nSamp; i++ {
		for j := 0; j < nSpec; j++ {
			x := float64(sSize) * mtx.Get(i, j)
			if x > 0 {
				// from ranlib.c: gengam(a, r) == dst.GammaNext(r, 1/a)
				x = dst.GammaNext(1/shape[j], shape[j]*x)
			} else {
				x = 0
			}
			mtx.Set(i, j, x)
		}
	}
}

// TruncSampler just truncates the matrix of species abundances to integer values.
func TruncSampler(mtx *aux.Matrix) {
	nSamp := mtx.R
	nSpec := mtx.C
	for i := 0; i < nSamp; i++ {
		for j := 0; j < nSpec; j++ {
			x := math.Floor(mtx.Get(i, j))
			if x < 0 {
				x = 0
			}
			mtx.Set(i, j, x)
		}
	}
}
