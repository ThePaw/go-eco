// Eyraud dissimilarity matrix
// Eyraud (1936) in Shi (1993)
// Warning: it gives values near zero for both identical, and complementary data!!! STRANGE!
package sim

import (
	. "go-eco.googlecode.com/hg/eco"
)

// Eyraud dissimilarity matrix
func EyraudBool_D(data *Matrix) *Matrix {
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = GetABCD(data, i, j)
			v := (a - ((a + b) * (a + c))) / ((a + b) * (a + c) * (b + d) * (c + d))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
