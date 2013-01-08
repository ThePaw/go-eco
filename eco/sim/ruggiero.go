// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Ruggiero similarity matrix

import (
	"code.google.com/p/go-eco/eco/aux"
)

// RuggieroBool_S returns a Ruggiero similarity matrix for boolean data. 
// Ruggiero et al. (1998), Koleff et al. (2003). 
func RuggieroBool_S(data *aux.Matrix) *aux.Matrix {
	var (
		a, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, _, c, _ = aux.GetABCD(data, i, j)
			v := a / (a + c)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
