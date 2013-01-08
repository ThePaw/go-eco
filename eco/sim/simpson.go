// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Simpson similarity
// Simpson (1960), Shi (1993)

import (
	"code.google.com/p/go-eco/eco/aux"
	"math"
)

// Simpson1Bool_D returns a Simpson dissimilarity matrix #1 for boolean data. 
func Simpson1Bool_D(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = aux.GetABCD(data, i, j)
			v := math.Min(b, c) / (math.Min(b, c) + a)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Simpson similarity matrix #2
// Simpson2Bool_S returns a Simpson similarity matrix #2 for floating-point data. 
func Simpson2Bool_S(data *aux.Matrix) *aux.Matrix {
	var (
		a, b float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, _, _ = aux.GetABCD(data, i, j)
			v := a/a + b
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
