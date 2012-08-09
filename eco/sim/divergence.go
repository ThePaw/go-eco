// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Divergence dissimilarity matrix
// Ellis et al. (1993)

import (
	"code.google.com/p/go-eco/eco/aux"
	"math"
)

// DivergenceBool_D returns a Divergence dissimilarity matrix for boolean data.
func DivergenceBool_D(data *aux.Matrix) *aux.Matrix {
	// Divergence dissimilarity matrix
	// Ellis et al. (1993)
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = aux.GetABCD(data, i, j)
			v := (math.Sqrt(b+c) / math.Sqrt(a+b+c+d))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
