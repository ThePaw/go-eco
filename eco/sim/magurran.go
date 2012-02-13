// Magurran dissimilarity matrix
// Magurran (1988)

package sim

import (
	. "go-eco.googlecode.com/hg/eco"
)

// Magurran dissimilarity matrix
func MagurranBool_D(data *Matrix) *Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = GetABCD(data, i, j)
			v := (2*a + b + c) * (1 - (a / (a + b + c)))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
