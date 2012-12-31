// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Faith similarity

import (
	"code.google.com/p/go-eco/eco/aux"
)

// FaithBool_S returns a Faith similarity matrix for boolean data  (S26 index in Legendre & Legendre, 1998). 
// Faith (1983). 
func FaithBool_S(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = aux.GetABCD(data, i, j)
			v := (a + d/2) / (a + b + c + d)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
