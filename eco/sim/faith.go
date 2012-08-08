// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Faith similarity

import (
	. "go-eco.googlecode.com/hg/eco"
)

// FaithBool_S returns a Faith similarity matrix for boolean data  (S26 index in Legendre & Legendre, 1998). 
func FaithBool_S(data *Matrix) *Matrix {
	// Faith (1983)
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = GetABCD(data, i, j)
			v := (a + d/2) / (a + b + c + d)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
