// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// McConnagh similarity matrix

import (
	"code.google.com/p/go-eco/eco/aux"
)

// McConnaghBool_S returns a McConnagh similarity matrix for boolean data. 
// Hubalek (1982). 
func McConnaghBool_S(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = aux.GetABCD(data, i, j)
			v := ((a * a) - (b * c)) / ((a + b) * (a + c))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
