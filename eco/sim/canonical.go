// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Canonical distance
// == Euclidean

import (
	"code.google.com/p/go-eco/eco/aux"
	"math"
)

// Canonical_D returns a Canonical distance distance matrix for floating-point data. 
func Canonical_D(data *aux.Matrix) *aux.Matrix {
	// Algorithm from R:ade4
	rows := data.R
	cols := data.C
	out := aux.NewMatrix(rows, rows)

	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			sum := 0.0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				sum += (x - y) * (x - y)
			}
			v := math.Sqrt(sum)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
