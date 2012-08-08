// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Cody dissimilarity matrix
// Cody (1993)

import (
	. "go-eco.googlecode.com/hg/eco"
)

// CodyBool_D returns a Cody dissimilarity matrix for boolean data.
func CodyBool_D(data *Matrix) *Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = GetABCD(data, i, j)
			v := 1 - ((a * (2*a + b + c)) / (2 * (a + b) * (a + c)))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
