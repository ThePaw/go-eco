// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Squared average distance matrix

import (
	"code.google.com/p/go-eco/eco/aux"
)

// AverageSqBool_D returns a Squared average distance matrix for boolean data. 
func AverageSqBool_D(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = aux.GetABCD(data, i, j)
			v := (b + c) / (a + b + c + d)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
