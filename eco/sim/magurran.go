// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Magurran dissimilarity. 
// Magurran (1988). 

import (
	"code.google.com/p/go-eco/eco/aux"
)

// MagurranBool_D returns a Magurran dissimilarity matrix for boolean data. 
// Magurran (1988). 
func MagurranBool_D(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = aux.GetABCD(data, i, j)
			v := (2*a + b + c) * (1 - (a / (a + b + c)))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
