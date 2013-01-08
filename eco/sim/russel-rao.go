// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Russell-Rao similarity matrix 
// Russel & Rao (1940)
// = legendre of R:simba
// Gower & Legendre (1986), Russell/Rao in Ellis et al. (1993)
// Legendre & Legendre (1998): 257, eq. 7.14, S11

import (
	"code.google.com/p/go-eco/eco/aux"
)

// RusselRaoBool_S returns a Russell-Rao similarity matrix for boolean data. 
// Legendre & Legendre (1998): 257, eq. 7.14 (S11 index). 
func RusselRaoBool_S(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = aux.GetABCD(data, i, j)
			v := a / (a + b + c + d)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
