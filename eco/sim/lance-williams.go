// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package sim
// Lance-Williams distance

import (
	. "go-eco.googlecode.com/hg/eco"
)

// LanceWilliamsBool_D returns a Lance-Williams distance matrix for boolean data. 
func LanceWilliamsBool_D(data *Matrix) *Matrix {
	var (
		a, b, c float64
	)

	WarnIfNotBool(data)

	rows := data.R
	out := NewMatrix(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			a, b, c, _ = GetABCD(data, i, j)
			v := (b + c) / (2 * (a + b + c))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
