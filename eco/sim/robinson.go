// Robinson's similarity function

package sim

import (
	. "go-eco.googlecode.com/hg/eco"
	. "math"
)

func robinson_S(data *Matrix) *Matrix {
	rows := data.R
	cols := data.C
	out := NewMatrix(data.R, data.R) // square similarity matrix row vs. row
	percent := NewMatrix(data.R, data.C)

	// Set diagonal to 200
	for i := 0; i < data.R; i++ {
		out.Set(i, i, 200.0)
	}

	// calculate percentages
	for i := 0; i < data.R; i++ {
		rowsum := 0.0
		for j := 0; j < data.C; j++ {

			rowsum += data.Get(i, j)
		}
		for j := 0; j < data.C; j++ {
			percent.Set(i, j, data.Get(i, j)*100.0/rowsum)

		}
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sum := 0.0
			for k := 0; k < cols; k++ {
				x := percent.Get(i, k)
				y := percent.Get(j, k)
				sum += Abs(x - y)
			}
			v := 200.0 - sum
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
