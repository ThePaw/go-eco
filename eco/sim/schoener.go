// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Schoener’s (1968) measure of niche overlap. 
// Schoener, T. W. 1968. Anolis lizards of Bimini: resource partitioning in a complex fauna. Ecology 49:704–726.

import (
	"code.google.com/p/go-eco/eco/aux"
	"math"
)

// Schoener_S returns a Schoener’s (1968) measure of niche overlap matrix for floating-point data. 
func Schoener_S(data *aux.Matrix) *aux.Matrix {
	rows := data.R
	cols := data.C
	out := aux.NewMatrix(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sum := 0.0
			for k := 0; k < cols; k++ {
				x := math.Sqrt(data.Get(i, k))
				y := math.Sqrt(data.Get(j, k))

				sum += math.Abs(x - y)
			}
			v := 1 - 0.5*sum
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
