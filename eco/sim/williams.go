// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Williams dissimilarity matrix
// Williams (1996), Koleff et al. (2003)

import (
	. "code.google.com/p/go-eco/eco"
	"math"
)

// Williams1Bool_D returns a Williams dissimilarity matrix #1 for boolean data. 
// Williams (1996), Koleff et al. (2003)
func Williams1Bool_D(data *Matrix) *Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = GetABCD(data, i, j)
			v := math.Min(b, c) / (a + b + c)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Williams2Bool_D returns a Williams dissimilarity matrix #2 for boolean data. 
// Williams (1996), Koleff et al. (2003)
func Williams2Bool_D(data *Matrix) *Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = GetABCD(data, i, j)
			//			v := (2*b*c + 1) / (math.Pow(a+b+c, 2) - (a + b + c))
			v := ((b * c) + 1) / ((((a + b + c) * (a + b + c)) - (a + b + c)) / 2)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
