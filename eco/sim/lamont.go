// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Lamont similarity matrix
// Lamont and Grant (1979)

import (
	"code.google.com/p/go-eco/eco/aux"
)

// LamontBool_S returns a Lamont similarity matrix for boolean data. 
func LamontBool_S(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = aux.GetABCD(data, i, j)
			v := a / (2*a + b + c)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
