// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Margaleff similarity matrix

import (
	"code.google.com/p/go-eco/eco/aux"
)

// MargaleffBool_S returns a Margaleff similarity matrix for boolean data. 
// Clifford & Stevenson (1975). 
func MargaleffBool_S(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = aux.GetABCD(data, i, j)
			v := (a * (a + b + c + d)) / ((a + b) * (a + c))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
