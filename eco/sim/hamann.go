// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Hamann similarity matrix

import (
	"code.google.com/p/go-eco/eco/aux"
)

// HamannBool_S returns a Hamann similarity matrix for boolean data.
// Holley JW, Guilford JP 1964 A note on the G index of agreement. Educational and Psychological Measurement, 24(7):749-753.
func HamannBool_S(data *aux.Matrix) *aux.Matrix {
	// Legendre & Legendre 1998: 256, eq. 7.7. 
	// S9 index of Gower & Legendre (1986)
	// S6 index of R:ade4:dist.binary
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	aux.WarnIfNotBool(data)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = aux.GetABCD(data, i, j)
			v := (a + d - b - c) / (a + b + c + d)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
