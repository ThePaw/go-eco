// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Lance-Williams distance

import (
	"code.google.com/p/go-eco/eco/aux"
)

// LanceWilliamsBool_D returns a Lance-Williams distance matrix for boolean data. 
func LanceWilliamsBool_D(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c float64
	)

	aux.WarnIfNotBool(data)

	rows := data.R
	out := aux.NewMatrix(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			a, b, c, _ = aux.GetABCD(data, i, j)
			v := (b + c) / (2 * (a + b + c))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
