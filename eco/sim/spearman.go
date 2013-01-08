// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

import (
	"code.google.com/p/go-eco/eco/aux"
)

// SpearmanRho_S returns a Spearman's ρ (rho)  similarity matrix for floating-point data. 
func SpearmanRho_S(data *aux.Matrix) *aux.Matrix {
	rows := data.R
	cols := data.C
	out := aux.NewMatrix(rows, rows)
	ranks := aux.NewMatrix(rows, cols)

	// ToDo: check for ties

	// calculate ranks
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// count scores lower than this
			count := 0
			for k := 0; k < cols; k++ {
				if data.Get(i, k) <= data.Get(i, j) {
					count++
				}
			}
			ranks.Set(i, j, float64(count))
		}
	}

	for i := 0; i < rows; i++ {
		out.Set(i, i, 1.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sumd2 := 0.0
			for k := 0; k < cols; k++ {
				sumd2 += (ranks.Get(i, k) - ranks.Get(j, k)) * (ranks.Get(i, k) - ranks.Get(j, k))
			}
			v := 1.0 - 6.0*sumd2/float64(cols*cols*cols-cols)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
