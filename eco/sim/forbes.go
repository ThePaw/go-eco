// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Forbes similarity matrix

import (
	"code.google.com/p/go-eco/eco/aux"
	"math"
)

// ForbesBool_S returns a Forbes similarity matrix for boolean data. 
// Forbes (1925), Shi (1993). 
func ForbesBool_S(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = aux.GetABCD(data, i, j)
			v := (a*(a+b+c+d) - (2 * math.Max(a+b, a+c))) / (((a + b + c + d) * math.Min(a+b, a+c)) - (2 * math.Max(a+b, a+c)))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
