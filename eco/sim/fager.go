// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

import (
	"code.google.com/p/go-eco/eco/aux"
	"math"
)

// FagerBool_S returns a Fager similarity matrix for boolean data.
// Fager (1957), Shi (1993). 
func FagerBool_S(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = aux.GetABCD(data, i, j)
			v := (a / math.Sqrt(math.Min(a+b, a+c)*math.Max(a+b, a+c))) - (1 / (2 * math.Sqrt(math.Min(a+b, a+c))))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
