// Harte dissimilarity matrix
// Harte & Kinzig (1997), Koleff et al. (2003)

package sim

import (
	. "go-eco.googlecode.com/hg/eco"
)

// Harte dissimilarity matrix
func HarteBool_D(data *Matrix) *Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = GetABCD(data, i, j)
			v := 1 - (2 * a / (2*a + b + c))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
