// Similarity ratio matrix
package sim

import (
	. "go-eco.googlecode.com/hg/eco"
)

func SimRatio_S(data *Matrix) *Matrix {
	rows := data.R
	cols := data.C
	out := NewMatrix(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 1.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sumXX := 0.0
			sumYY := 0.0
			sumXY := 0.0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				sumXX += x * x
				sumYY += y * y
				sumXY += x * y
			}
			v := sumXY / (sumXX + sumYY - sumXY)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
