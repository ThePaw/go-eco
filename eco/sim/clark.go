// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package sim
// Clark dissimilarity

import (
	. "go-eco.googlecode.com/hg/eco"
	"math"
)

// Clark_D returns Clark dissimilarity matrix for count or interval data. 
func Clark_D(data *Matrix) *Matrix {
// Clark (1952)
// Legendre & Legendre (1998): 283, eq. 7.51 (D11 index)
	rows := data.R
	cols := data.C
	out := NewMatrix(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sum := 0.0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				t := (x - y) / (x + y)
				sum += t * t
			}
			v := math.Sqrt(sum / float64(cols))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
