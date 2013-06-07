package ser

// Sampling models.
// Performs sampling of species populations given population densities and sample size (sampling intensity).

import (
	"code.google.com/p/probab/dst"
	"math"
)

// betapara estimates the parameters a, b of beta distribution from expected proportion (pi), binomial denominator (m), and shape parameter (tau2). 
func betapara(pi, m, tau2 float64) (a, b float64) {
	// Solution (hopefully correct) of Exercise 4.17 of McCullagh & Nelder 1989, helped by Moore, Appl Stat 36, 8-14; 1987.
	// Thanks to Jari Oksanen.
	t1 := tau2 * m
	t2 := t1 - m - tau2 + 1
	t3 := 1 / (1 + t1 - tau2)
	t4 := t2 * t3
	a = -t4 * pi
	b = t4 * (pi - 1)
	return
}

// PoissonSampler implements Poisson sampling of the matrix of species abundances.
func PoissonSampler(mtx Matrix64, sEffort float64) IntMatrix {
	const eps = 1e-6
	var y float64
	nSamp := mtx.Rows()
	nSpec := mtx.Cols()
	out := NewIntMatrix(nSamp, nSpec)
	for i := 0; i < nSamp; i++ {
		for j := 0; j < nSpec; j++ {
			x := sEffort * mtx[i][j]
			if x > eps { // do not sample very low x (to speed up)
				y = float64(dst.PoissonNext(x))
			} else {
				y = 0
			}
			out[i][j] = int(y)
		}
	}
	return out
}

// MultinomialSampler implements Multinomial in-place sampling of the matrix of species abundances.
func MultinomialSampler(mtx Matrix64, sSize int64) IntMatrix {
	nSamp := mtx.Rows()
	nSpec := mtx.Cols()
	out := NewIntMatrix(nSamp, nSpec)
	for i := 0; i < nSamp; i++ { // for every sample

		// recalculate species abundances to proportions
		θ := make([]float64, nSpec)
		sum := 0.0
		for j := 0; j < nSpec; j++ {
			sum += mtx[i][j]
		}
		for j := 0; j < nSpec; j++ {
			θ[j] = mtx[i][j] / sum
		}

		x := dst.MultinomialNext(θ, sSize)
		for j := 0; j < nSpec; j++ {
			out[i][j] = int(x[j])
		}
	}
	return out
}

// BetaBinomialSampler implements BetaBinomial sampling of the matrix of species abundances.
func BetaBinomialSampler(mtx Matrix64, sSize int64, shape []float64) IntMatrix {
	nSamp := mtx.Rows()
	nSpec := mtx.Cols()
	out := NewIntMatrix(nSamp, nSpec)
	for i := 0; i < nSamp; i++ { // for every sample

		// recalculate species abundances to proportions
		θ := make([]float64, nSpec)
		sum := 0.0
		for j := 0; j < nSpec; j++ {
			sum += mtx[i][j]
		}
		for j := 0; j < nSpec; j++ {
			θ[j] = mtx[i][j] / sum
		}

		y := dst.MultinomialNext(θ, sSize)

		for j := 0; j < nSpec; j++ {
			x := float64(y[j])
			if x > 0.0 {
				a, b := betapara(x, float64(sSize), shape[j])
				x = dst.BetaNext(a, b)
				x = float64(dst.BinomialNext(sSize, x))
			}
			out[i][j] = int(x)
		}
	}
	return out
}

// NegBinomSampler implements Poisson sampling of the matrix of species abundances.
func NegBinomSampler(mtx Matrix64, sSize int64, shape []float64) IntMatrix {
	nSamp := mtx.Rows()
	nSpec := mtx.Cols()
	out := NewIntMatrix(nSamp, nSpec)
	for i := 0; i < nSamp; i++ {
		for j := 0; j < nSpec; j++ {
			x := float64(sSize) * mtx[i][j]
			if x > 0 {
				// from ranlib.c: gengam(a, r) == dst.GammaNext(r, 1/a)
				x = dst.GammaNext(1/shape[j], shape[j]*x)
			} else {
				x = 0
			}
			out[i][j] = int(x)
		}
	}
	return out
}

// TruncSampler just truncates the matrix of species abundances to integer values.
func TruncSampler(mtx Matrix64) IntMatrix {
	nSamp := mtx.Rows()
	nSpec := mtx.Cols()
	out := NewIntMatrix(nSamp, nSpec)
	for i := 0; i < nSamp; i++ {
		for j := 0; j < nSpec; j++ {
			x := math.Floor(mtx[i][j])
			if x < 0 {
				x = 0
			}
			out[i][j] = int(x)
		}
	}
	return out
}

// TruncSampler2 just truncates the matrix of species abundances to integer values. REMOVE after rewriting type Sampler...
func TruncSampler2(mtx Matrix64, sEffort float64) IntMatrix {
	nSamp := mtx.Rows()
	nSpec := mtx.Cols()
	out := NewIntMatrix(nSamp, nSpec)
	for i := 0; i < nSamp; i++ {
		for j := 0; j < nSpec; j++ {
			x := math.Floor(mtx[i][j])
			if x < 0 {
				x = 0
			}
			out[i][j] = int(x)
		}
	}
	return out
}
