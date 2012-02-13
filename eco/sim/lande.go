// Lande dissimilarity

package sim

import (
	. "go-eco.googlecode.com/hg/eco"
)

// Lande dissimilarity matrix
func LandeBool_D(data *Matrix) *Matrix {
	var (
		b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			_, b, c, _ = GetABCD(data, i, j)
			v := (b + c) / 2
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
