// Růžička distance and similarity

package sim

import (
	. "code.google.com/p/go-eco/eco"
)

// Růžička distance matrix
func Ruzicka_D(data *Matrix) *Matrix {
	var (
		aa, bb, jj float64
		out        *Matrix
	)

	rows := data.R
	out = NewMatrix(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			aa, bb, jj, _ = GetABJPmin(data, i, j)
			// (A+B-2*J)/(A+B-J)
			v := (aa + bb - 2*jj) / (aa + bb - jj)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
