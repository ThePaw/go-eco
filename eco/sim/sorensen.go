// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Sørensen similarity and distance
// Soerensen (1948)

import (
	"code.google.com/p/go-eco/eco/aux"
)

// SorensenBool_S returns a Sørensen similarity matrix for boolean data. 
// Legendre & Legendre (1998): 256, eq. 7.11  (S8 index). 
func SorensenBool_S(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = aux.GetABCD(data, i, j)
			v := 2 * a / (2*a + b + c)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// SorensenBool_D returns a Sørensen distance matrix for boolean data. 
func SorensenBool_D(data *aux.Matrix) *aux.Matrix {
	var (
		aa, bb, jj float64
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	aux.WarnIfNotBool(data)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			aa, bb, jj, _ = aux.GetABJPbool(data, i, j)
			// (A+B-2*J)/(A+B)
			v := (aa + bb - 2*jj) / (aa + bb)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
