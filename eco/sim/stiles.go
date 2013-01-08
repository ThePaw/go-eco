// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Stiles similarity. 

import (
	"code.google.com/p/go-eco/eco/aux"
	"math"
)

// StilesBool_S returns a Stiles similarity matrix for boolean data. 
// Stiles (1946). 
func StilesBool_S(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = aux.GetABCD(data, i, j)
			t1 := a + b + c + d
			t2 := math.Abs(a*d - b*c)
			t3 := (a + b) * (a + c) * (b + d) * (c + d)
			v := math.Log(t1 * (t2 - t1/2) * (t2 - t1/2) / t3)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
