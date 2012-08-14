// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

// Roberts similarity

package sim

import (
	"code.google.com/p/go-eco/eco/aux"
	"math"
)

// Roberts_S returns a Roberts similarity matrix for floating-point data. 
// Roberts 1986. 
func Roberts_S(data *aux.Matrix) *aux.Matrix {
// Algorithm inspired by R:labdsv. 
	rows := data.R
	cols := data.C
	out := aux.NewMatrix(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			numer := 0.0
			denom := 0.0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				min := math.Min(x, y)
				max := math.Max(x, y)
				numer += (x + y) * (min / max)
				denom += x + y
			}
			v := numer / denom
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
