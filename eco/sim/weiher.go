// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Weiher dissimilarity matrix
// Weiher & Boylen (1994)

import (
	"code.google.com/p/go-eco/eco/aux"
)

// WeiherBool_D returns a Weiher dissimilarity matrix for boolean data. 
// Weiher & Boylen (1994)
func WeiherBool_D(data *aux.Matrix) *aux.Matrix {
	var (
		b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			_, b, c, _ = aux.GetABCD(data, i, j)
			v := b + c
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
