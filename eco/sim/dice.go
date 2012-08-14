// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Dice's similarity and dissimilarity matrix

import (
	"code.google.com/p/go-eco/eco/aux"
	"math"
)

// DiceBool_S returns a Dice's similarity matrix for boolean data.
// Dice (1945), Wolda (1981). 
func DiceBool_S(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = aux.GetABCD(data, i, j)
			v := a / (math.Min(b+a, c+a))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// DiceBool_D returns a Dice's dissimilarity matrix for boolean data.
func DiceBool_D(data *aux.Matrix) *aux.Matrix {
	// Dice's dissimilarity
	// it is not a proper distance metric as it does not possess the property of triangle inequality
	// Dice = 2*Jaccard / (1 + Jaccard)
	// Formula from R:vegan 
	var (
		aa, bb, jj float64
		out        *aux.Matrix
	)

	rows := data.R
	out = aux.NewMatrix(rows, rows)
	aux.WarnIfNotBool(data)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			aa, bb, jj, _ = aux.GetABJPquad(data, i, j) // quadratic terms
			// 1-2*J/(A*B)
			v := 1.0 - 2.0*jj/(aa*bb)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
