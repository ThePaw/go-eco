// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Chi - Squared similarity matrix
// Yule & Kendall (1950)

import (
	"code.google.com/p/go-eco/eco/aux"
)

// ChiSquaredBool_S returns Chi - Squared similarity matrix for boolean data. 
func ChiSquaredBool_S(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = aux.GetABCD(data, i, j)
			t1 := a + b + c + d
			t2 := (a*d - b*c)
			t3 := (a + b) * (a + c) * (b + d) * (c + d)
			// ((a+b+c+d)*((a*d) - (b*c))^2) / ((a+b)*(a+c)*(b+d)*(c+d))
			v := t1 * t2 * t2 / t3
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
