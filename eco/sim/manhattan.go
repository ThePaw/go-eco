// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Manhattan distance
// Also known as rectilinear distance, Minkowski's L1 distance, taxicab metric, or city-block distance (metric). 

import (
	"code.google.com/p/go-eco/eco/aux"
	. "math"
)

// Manhattan_D returns a Manhattan distance matrix for floating-point data. 
// Legendre & Legendre (1998): 282, eq. 7.45 (D7 index). 
func Manhattan_D(data *aux.Matrix) *aux.Matrix {
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
				sum += Abs(x - y)
			}
			out.Set(i, j, sum)
			out.Set(j, i, sum)
		}
	}
	return out
}

// ManhattanBool_D returns a Boolean Manhattan dissimilarity matrix for boolean data. 
func ManhattanBool_D(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = aux.GetABCD(data, i, j)
			v := (b + c) / (a + b + c + d)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
