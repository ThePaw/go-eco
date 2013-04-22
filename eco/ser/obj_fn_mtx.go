// Copyright 2012 The Gt Authors. All rights reserved. See the LICENSE file.

package ser

// Objective (loss and gain) functions for mxn data matrices. 

import (
	"math"
	"code.google.com/p/go-fn/fn"
)

// MooreStressLoss returns the Moore Stress criterion (Niermann 2005:42, Eq. 1, 2).
func MooreStressLoss(mtx Matrix64, rowPerm, colPerm IntVector) float64 {
	r, c := mtx.Dims()
	if !(rowPerm.Len() == r && colPerm.Len() == c) {
		panic("bad dimensions")
	}
	stress := 0.0
	for i := 1; i <= r; i++ {
		for j := 1; j <= c; j++ {
			for l := imax(1, i-1); l <= imin(r, i+1); l++ {
				for m := imax(1, j-1); m <= imin(c, j+1); m++ {
					val := mtx[rowPerm[i-1]][colPerm[j-1]] - mtx[rowPerm[l-1]][colPerm[m-1]]
					val *= val
					stress += val
				}
			}
		}
	}
	return stress
}

// VonNeumannStressLoss returns the Moore Stress criterion (Niermann 2005:42).
func VonNeumannStressLoss(mtx Matrix64, rowPerm, colPerm IntVector) float64 {
	r, c := mtx.Dims()
	if !(rowPerm.Len() == r && colPerm.Len() == c) {
		panic("bad dimensions")
	}
	stress := 0.0
	for i := 1; i <= r; i++ {
		for j := 1; j <= c; j++ {
			for l := imax(1, i-1); l <= imin(r, i+1); l++ {
				for m := imax(1, j-1); m <= imin(c, j+1); m++ {
					if l == i || m == j {
						val := mtx[rowPerm[i-1]][colPerm[j-1]] - mtx[rowPerm[l-1]][colPerm[m-1]]
						val *= val
						stress += val
					}
				}
			}
		}
	}
	return stress
}

// MEffGain returns the measure of Effectiveness (McCormick 1972).
func MEffGain(mtx Matrix64, rowPerm, colPerm IntVector) float64 {
	var x0, x1, x2, x3, x4 float64
	rows, cols := mtx.Dims()

	if !(rowPerm.Len() == rows && colPerm.Len() == cols) {
		panic("bad dimensions")
	}
	gain := 0.0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			x0 = mtx[rowPerm[i]][colPerm[j]]
			if j-1 < 0 {
				x1 = 0
			} else {
				x1 = mtx[rowPerm[i]][colPerm[j-1]]
			}
			if j+1 > cols-1 {
				x2 = 0
			} else {
				x2 = mtx[rowPerm[i]][colPerm[j+1]]
			}
			if i-1 < 0 {
				x3 = 0
			} else {
				x3 = mtx[rowPerm[i-1]][colPerm[j]]
			}

			if i+1 > rows-1 {
				x4 = 0
			} else {

				x4 = mtx[rowPerm[i+1]][colPerm[j]]
			}
			gain += x0 * (x1 + x2 + x3 + x4)
		}
	}
	return gain / 2
}

// MirrorLoss computes energy E(p) of the permuted matrix according to Miklos (2005:3400), Eqs. 2, 3.
func MirrorLoss(mtx Matrix64, rowPerm, colPerm IntVector) float64 {
	rows, cols := mtx.Dims()
	if !(rowPerm.Len() == rows && colPerm.Len() == cols) {
		panic("bad dimensions")
	}

	av := 0.0
	for i := 0; i < rows; i++ {
		for k := 0; k < cols; k++ {
			for l := 0; l < cols; l++ {
				if k < l {
					av += math.Abs(mtx[rowPerm[i]][colPerm[k]] - mtx[rowPerm[i]][colPerm[l]])
				}
			}
		}
	}

	for k := 0; k < cols; k++ {
		for i := 0; i < rows; i++ {
			for j := 0; j < rows; j++ {
				if i < j {
					av += math.Abs(mtx[rowPerm[i]][colPerm[k]] - mtx[rowPerm[j]][colPerm[k]])
				}
			}
		}
	}

	denom := float64(rows)*fn.BinomCoeff(int64(cols), 2) + float64(cols)*fn.BinomCoeff(int64(rows), 2)
	av /= denom

	loss := 0.0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols-1; j++ {
			loss += math.Abs(mtx[rowPerm[i]][colPerm[j]] - mtx[rowPerm[i]][colPerm[j+1]])
		}
	}

	for i := 0; i < rows-1; i++ {
		for j := 0; j < cols; j++ {
			loss += math.Abs(mtx[rowPerm[i]][colPerm[j]] - mtx[rowPerm[i+1]][colPerm[j]])
		}
	}

	for j := 0; j < cols; j++ {
		loss += (math.Abs(mtx[rowPerm[0]][colPerm[j]]-mtx[rowPerm[1]][colPerm[j]]) + (math.Abs(mtx[rowPerm[rows-1]][colPerm[j]] - mtx[rowPerm[rows]][colPerm[j]])))
	}

	for i := 0; i <rows; i++ {
		loss += (math.Abs(mtx[rowPerm[i]][colPerm[0]]-mtx[rowPerm[i]][colPerm[1]]) + (math.Abs(mtx[rowPerm[i]][colPerm[cols-1]] - mtx[rowPerm[i]][colPerm[cols]])))
	}
	loss /= av
	return loss
}

// PsiLoss computes energy ψ(p) of the permuted similarity matrix according to Podani (1994); see  Miklos (2005), Eq. 4.
func PsiLoss(mtx Matrix64, rowPerm, colPerm IntVector) float64 {
	rows, cols := mtx.Dims()
	if !(rowPerm.Len() == rows && colPerm.Len() == cols) {
		panic("bad dimensions")
	}
	loss := 0.0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			x := mtx[rowPerm[i]][colPerm[j]]
			a := math.Abs(float64(cols*(i+1))/float64(rows) - float64(j+1))
			b := math.Abs(float64(rows*(j+1))/float64(cols) - float64(i+1))
			loss += x*a + b
		}
	}
	return loss
}


