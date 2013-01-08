// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Lande dissimilarity

import (
	"code.google.com/p/go-eco/eco/aux"
)

// LandeBool_D returns a Lande dissimilarity matrix for boolean data. 
func LandeBool_D(data *aux.Matrix) *aux.Matrix {
	var (
		b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			_, b, c, _ = aux.GetABCD(data, i, j)
			v := (b + c) / 2
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
