// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Geodesic distance
// Legendre & Legendre (1998): 280, eq. 7.39 (D4 index)

import (
	. "code.google.com/p/go-eco/eco"
	"math"
)

// Geodesic_D returns a Geodesic distance matrix for floating-point data. 
func Geodesic_D(data *Matrix) *Matrix {
	rows := data.R
	cols := data.C
	out := NewMatrix(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sumXY := 0.0
			sumXX := 0.0
			sumYY := 0.0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				sumXY += x * y
				sumXX += x * x
				sumYY += y * y
			}
			v := math.Sqrt(2 * (1 - (sumXY / math.Sqrt(sumXX*sumYY))))
			v = math.Acos(1 - v*v/2)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
