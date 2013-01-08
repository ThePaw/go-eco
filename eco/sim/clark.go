// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Clark dissimilarity

import (
	"code.google.com/p/go-eco/eco/aux"
	"math"
)

// Clark_D returns Clark dissimilarity matrix for count or interval data. 
// Clark (1952). 
// Legendre & Legendre (1998): 283, eq. 7.51 (D11 index). 
func Clark_D(data *aux.Matrix) *aux.Matrix {
	rows := data.R
	cols := data.C
	out := aux.NewMatrix(rows, rows)

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
