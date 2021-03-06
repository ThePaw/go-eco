// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Braun–Blanquet similarity matrix

import (
	"code.google.com/p/go-eco/eco/aux"
	"math"
)

// BraunBlanquetBool_S returns a Braun–Blanquet similarity matrix for boolean data. 
// Braun-Blanquet 1932; Magurran 2004.
func BraunBlanquetBool_S(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c float64
	)

	aux.WarnIfNotBool(data)

	rows := data.R
	out := aux.NewMatrix(rows, rows)

	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = aux.GetABCD(data, i, j)
			v := a / math.Max(b+a, c+a)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
