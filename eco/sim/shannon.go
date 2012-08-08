// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Shannon distance matrix

import (
	. "code.google.com/p/go-eco/eco"
	"math"
)

// ShannonBool_D returns a Shannon distance matrix for boolean data. 
func ShannonBool_D(data *Matrix) *Matrix {
	var (
		b, c float64
	)

	WarnIfNotBool(data)

	rows := data.R
	out := NewMatrix(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			_, b, c, _ = GetABCD(data, i, j)
			v := 2.0 * (b + c) * math.Log(2.0)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
