// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Pearson rho correlations as similarity matrix

import (
	. "code.google.com/p/go-eco/eco"
	"math"
)

// PearsonRho_S returns a Pearson's ρ (rho)  similarity matrix for floating-point data. 
func PearsonRho_S(data *Matrix) *Matrix {
	rows := data.R
	cols := data.C
	out := NewMatrix(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 1.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sxx := 0.0
			syy := 0.0
			sxy := 0.0
			xmean := 0.0
			ymean := 0.0

			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				xmean += x
				ymean += y
			}
			xmean /= float64(cols)
			ymean /= float64(cols)

			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				sxx += x - xmean
				syy += y - ymean
				sxy += (x - xmean) * (y - ymean)
			}
			v := sxy / math.Sqrt(sxx*syy)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// PearsonPhiBool_S returns a Pearson's Φ  similarity matrix for boolean data. 
// Phi of Pearson, Gower & Legendre (1986), Yule (1912).
func PearsonPhiBool_S(data *Matrix) *Matrix {
	// !!! CHECK against L&L 1998 !!!
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = GetABCD(data, i, j)
			v := (a*d - b*c) / math.Sqrt((a+b)*(a+c)*(d+b)*(d+c))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
