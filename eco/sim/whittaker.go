// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Whittaker dissimilarity

import (
	. "go-eco.googlecode.com/hg/eco"
	"math"
)

// WhittakerBool_D returns a Whittaker dissimilarity matrix for boolean data. 
// Whittaker (1960), Magurran (1988)
func WhittakerBool_D(data *Matrix) *Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = GetABCD(data, i, j)
			v := ((a + b + c) / ((2*a + b + c) / 2)) - 1
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Whittaker_D returns a Whittaker distance matrix for count or interval data. 
// Whittaker (1952); 
// Legendre & Legendre (1998): 282, eq. 7.47 (D9 index). 
func Whittaker_D(data *Matrix) *Matrix {
	rows := data.R
	cols := data.C
	out := NewMatrix(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sumX := 0.0
			sumY := 0.0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				sumX += x
				sumY += y
			}
			sum := 0.0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				sum += math.Abs(x/sumX - y/sumY)
			}
			v := sum / 2
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
