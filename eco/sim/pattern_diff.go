// Pattern difference distance matrix
package sim

import (
	. "go-eco.googlecode.com/hg/eco"
)

func PatternDiffBool_D(data *Matrix) *Matrix {
	var (
		a, b, c, d float64
	)

	WarnIfNotBool(data)

	rows := data.R
	out := NewMatrix(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			a, b, c, d = GetABCD(data, i, j)
			p := (a + b + c + d)
			v := (b * c) / (p * p)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
