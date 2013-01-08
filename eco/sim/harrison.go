// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Harrison dissimilarity matrix

import (
	"code.google.com/p/go-eco/eco/aux"
	"math"
)

// HarrisonBool_D returns a Harrison dissimilarity matrix for boolean data.
// Harrison et al. (1992), Koleff et al. (2003). 
func HarrisonBool_D(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = aux.GetABCD(data, i, j)
			v := math.Min(b, c) / (math.Max(b, c) + a)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
