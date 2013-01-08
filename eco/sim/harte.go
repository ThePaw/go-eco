// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Harte dissimilarity matrix

import (
	"code.google.com/p/go-eco/eco/aux"
)

// HarteBool_D returns a Harte dissimilarity matrix for floating-point data. 
// Harte & Kinzig (1997), Koleff et al. (2003)
func HarteBool_D(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = aux.GetABCD(data, i, j)
			v := 1 - (2 * a / (2*a + b + c))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
