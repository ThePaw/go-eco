// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Colwell - Coddington - Gaston dissimilarity

import (
	"code.google.com/p/go-eco/eco/aux"
)

// CoCoGastonBool_D returns Colwell - Coddington - Gaston et al.  dissimilarity matrix for boolean data.
// Colwell & Coddington (1948), Gaston et al. (2001). 
func CoCoGastonBool_D(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = aux.GetABCD(data, i, j)
			v := (b + c) / (a + b + c)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
