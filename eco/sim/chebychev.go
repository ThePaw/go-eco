// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Chebychev distance
// Chebychev distance is a special case of the Minkowski metric, where p = ∞

import (
	"code.google.com/p/go-eco/eco/aux"
	. "math"
)

// Chebychev_D returns Chebychev distance matrix for floating-point data. 
// Chebychev distance is a special case of the Minkowski metric, where p = ∞.
func Chebychev_D(data *aux.Matrix) *aux.Matrix {
	rows := data.R
	cols := data.C
	out := aux.NewMatrix(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			v := 0.0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				v = Max(v, Abs(x-y))

			}
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
