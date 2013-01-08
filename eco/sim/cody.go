// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Cody dissimilarity matrix

import (
	"code.google.com/p/go-eco/eco/aux"
)

// CodyBool_D returns a Cody dissimilarity matrix for boolean data.
// Cody (1993). 
func CodyBool_D(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = aux.GetABCD(data, i, j)
			v := 1 - ((a * (2*a + b + c)) / (2 * (a + b) * (a + c)))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
