// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Braun–Blanquet similarity matrix
// Braun-Blanquet 1932; Magurran 2004.

import (
	. "go-eco.googlecode.com/hg/eco"
	"math"
)

// Braun–Blanquet similarity
func BraunBlanquetBool_S(data *Matrix) *Matrix {
	var (
		a, b, c float64
	)

	WarnIfNotBool(data)

	rows := data.R
	out := NewMatrix(rows, rows)

	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = GetABCD(data, i, j)
			v := a / math.Max(b+a, c+a)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
