// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Baroni-Urbani and Buser (dis)similarity matrix
// Baroni-Urbani & Buser (1976), Wolda (1981)

import (
	"code.google.com/p/go-eco/eco/aux"
	"math"
)

// BaroniUrbaniBool_S returns a Baroni-Urbani and Buser  similarity matrix for floating-point data. 
func BaroniUrbaniBool_S(data *aux.Matrix) *aux.Matrix {
	var (
		out        *aux.Matrix
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	aux.WarnIfNotBool(data)

	rows := data.R
	out = aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = aux.GetABCD(data, i, j)
			v := ((math.Sqrt(a * d)) + a) / ((math.Sqrt(a * d)) + b + c + a)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// BaroniUrbaniBool_D returns a Baroni-Urbani and Buser dissimilarity matrix for floating-point data. 
func BaroniUrbaniBool_D(data *aux.Matrix) *aux.Matrix {
// according to R:vegan
	var (
		a, b, c, d float64
	)

	aux.WarnIfNotBool(data)

	rows := data.R
	out := aux.NewMatrix(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			a, b, c, d = aux.GetABCD(data, i, j)
			sqrtcd := math.Sqrt(float64(c * d))
			v := 1.0 - (sqrtcd+c)/(sqrtcd+a+b+c)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
