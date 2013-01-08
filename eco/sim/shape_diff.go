// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Shape difference distance matrix

import (
	"code.google.com/p/go-eco/eco/aux"
)

// ShapeDiffBool_D returns a Shape difference distance matrix for boolean data. 
func ShapeDiffBool_D(data *aux.Matrix) *aux.Matrix {
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
			p := (a + b + c + d)
			v := (p*(b+c) - (b-c)*(b-c)) / (p * p)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
