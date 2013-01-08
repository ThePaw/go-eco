// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Michael similarity matrix

import (
	"code.google.com/p/go-eco/eco/aux"
)

// MichaelBool_S returns a Michael distance matrix, for boolean data
// Michael (1920), Shi (1993). 
func MichaelBool_S(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = aux.GetABCD(data, i, j)
			v := (4 * ((a * d) - (b * c))) / ((a+d)*(a+d) + (b+c)*(b+c))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
