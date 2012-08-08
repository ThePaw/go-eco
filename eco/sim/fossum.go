// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Fossum similarity matrix
// Holliday et al. (2002), Ellis et al. (1993)

import (
	. "code.google.com/p/go-eco/eco"
)

// FossumBool_S returns a Fossum similarity matrix for boolean data.
func FossumBool_S(data *Matrix) *Matrix {
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = GetABCD(data, i, j)
			v := ((a + b + c + d) * (-1 * ((a / 2) * (a / 2)))) / ((a + b) * (a + c))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
