// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Ochiai distance and similarity
// Ochiai (1957)

import (
	"code.google.com/p/go-eco/eco/aux"
	"math"
)

// OchiaiBool_S returns a Ochiai similarity matrix for boolean data. 
// Ochiai (1957); 
// Legendre & Legendre (1998): 258, eq. 7.17 (S14 index). 
func OchiaiBool_S(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = aux.GetABCD(data, i, j)
			v := a / math.Sqrt((a+b)*(a+c))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// OchiaiBool_D returns a Ochiai distance matrix for boolean data. 
// According to R: vegan. 
func OchiaiBool_D(data *aux.Matrix) *aux.Matrix {
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
			// 1-J/sqrt(A*B)
			v := 1.0 - jj/math.Sqrt(aa*bb)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
