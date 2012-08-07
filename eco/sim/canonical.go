// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package sim
// Canonical distance
// == Euclidean

import (
	. "go-eco.googlecode.com/hg/eco"
	"math"
)

// Canonical distance matrix, float data
// Algorithm from R:ade4
func Canonical_D(data *Matrix) *Matrix {
	rows := data.R
	cols := data.C
	out := NewMatrix(rows, rows)

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
