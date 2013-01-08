// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Peirce similarity matrix

import (
	"code.google.com/p/go-eco/eco/aux"
)

// PeirceBool_S returns a Peirce similarity matrix for boolean data. 
// Peirce (1884) ???
func PeirceBool_S(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = aux.GetABCD(data, i, j)
			v := (a*d - b*c) / ((a + c) * (b + d))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
