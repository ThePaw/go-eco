// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package ser

// Matrix energy functions.

import (
	"math"
)

// Psi computes energy of the permuted similarity matrix according to Podani (1994); see Miklos 2005:3401, Eq. 4.
func Psi(matrix IntMatrix, perm IntVector) float64 {
	var (
		eE   float64
		s, o int
	)
	rows, cols := matrix.Dims()
	eE = 0
	for s = 0; s < rows; s++ {
		for o = 0; o < cols; o++ {
			x := float64(matrix[perm[s]][perm[o]])
			a := math.Abs(float64(cols*(s+1))/float64(rows) - float64(o+1))
			b := math.Abs(float64(rows*(o+1))/float64(cols) - float64(s+1))
			eE += x * (a + b)
		}
	}
	return eE
}

// Psi2 computes energy of the permuted *data* matrix according to Cejchan (unpublished) to seriate Q - matrix.
func Psi2(matrix IntMatrix, rowPerm, colPerm IntVector) float64 {
	var (
		eE, mod      float64
		s, o, posMod int
	)

	rows, cols := matrix.Dims()
	eE = 0
	for o = 0; o < cols; o++ {

		// for every species (column) find its modal value
		mod = 0
		for s = 0; s < rows; s++ {
			x := float64(matrix[rowPerm[s]][colPerm[o]])
			if mod < x {
				mod = x
				posMod = s
			}
		}
		// and use it to calc contribution to energy
		for s = 0; s < rows; s++ {
			x := float64(matrix[rowPerm[s]][colPerm[o]])
			d := math.Abs(float64(s - posMod))
			eE += x * d
		}
	}
	return eE
}

// Psi3 computes energy of the permuted *data* matrix according to Cejchan (unpublished) to seriate Q - matrix.
func Psi3(matrix IntMatrix, rowPerm, colPerm IntVector) float64 {
	var (
		eE, mod, ex  float64
		s, o, posMod int
	)

	rows, cols := matrix.Dims()
	ex = 3

	eE = 0
	for o = 0; o < cols; o++ {

		// for every species (column) find its modal value
		mod = 0
		for s = 0; s < rows; s++ {
			x := float64(matrix[rowPerm[s]][colPerm[o]])
			if mod < x {
				mod = x
				posMod = s
			}
		}
		// and use it to calc contribution to energy
		for s = 0; s < rows; s++ {
			x := float64(matrix[rowPerm[s]][colPerm[o]])
			d := math.Abs(math.Pow(float64(s-posMod), ex)) // like in Minkowski metric
			eE += x * d
		}
	}
	//	eE = math.Pow(eE, 1/ex)

	return eE
}
