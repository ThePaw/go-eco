// Faith similarity

package sim

import (
	. "go-eco.googlecode.com/hg/eco"
)

// Faith similarity matrix
// Faith (1983)
// Legendre & Legendre (1998): 258, eq. 7.18 (S26 index)
func FaithBool_S(data *Matrix) *Matrix {
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = GetABCD(data, i, j)
			v := (a + d/2) / (a + b + c + d)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
