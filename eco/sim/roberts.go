// Roberts similarity
// Roberts 1986
// Algorithm inspired by R:labdsv

package sim

import (
	. "go-eco.googlecode.com/hg/eco"
	"math"
)

// Roberts similarity matrix, float data
func Roberts_S(data *Matrix) *Matrix {
	rows := data.R
	cols := data.C
	out := NewMatrix(rows, rows)

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
