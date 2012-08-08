// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package sim
// Jaccard similarity matrix

import (
	. "go-eco.googlecode.com/hg/eco"
)

// JaccardBool_S returns a Jaccard similarity matrix for boolean data. 
// Legendre & Legendre 1998: 256, eq. 7.10 (S7 index). 
func JaccardBool_S(data *Matrix) *Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = GetABCD(data, i, j)
			v := a / (a + b + c)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
