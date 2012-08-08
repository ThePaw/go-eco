// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Lande dissimilarity

import (
	. "go-eco.googlecode.com/hg/eco"
)

// Lande dissimilarity matrix
// LandeBool_D returns a Lande dissimilarity matrix for boolean data. 
func LandeBool_D(data *Matrix) *Matrix {
	var (
		b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			_, b, c, _ = GetABCD(data, i, j)
			v := (b + c) / 2
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
