// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Colwell - Coddington - Gaston dissimilarity
// Colwell & Coddington (1948), Gaston et al. (2001)

import (
	. "go-eco.googlecode.com/hg/eco"
)

// CoCoGastonBool_D returns Colwell - Coddington - Gaston et al.  dissimilarity matrix for boolean data.
func CoCoGastonBool_D(data *Matrix) *Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = GetABCD(data, i, j)
			v := (b + c) / (a + b + c)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
