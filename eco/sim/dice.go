// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Dice's similarity and dissimilarity matrix
// Dice (1945), Wolda (1981)

import (
	. "go-eco.googlecode.com/hg/eco"
	"math"
)

// DiceBool_S returns a Dice's similarity matrix for boolean data.
func DiceBool_S(data *Matrix) *Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = GetABCD(data, i, j)
			v := a / (math.Min(b+a, c+a))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// DiceBool_D returns a Dice's dissimilarity matrix for boolean data.
func DiceBool_D(data *Matrix) *Matrix {
	// Dice's dissimilarity
	// it is not a proper distance metric as it does not possess the property of triangle inequality
	// Dice = 2*Jaccard / (1 + Jaccard)
	// Formula from R:vegan 
	var (
		aa, bb, jj float64
		out        *Matrix
	)

	rows := data.R
	out = NewMatrix(rows, rows)
	WarnIfNotBool(data)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			aa, bb, jj, _ = GetABJPquad(data, i, j) // quadratic terms
			// 1-2*J/(A*B)
			v := 1.0 - 2.0*jj/(aa*bb)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
