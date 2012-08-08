// Peirce similarity matrix
// Peirce (1884)

package sim

import (
	. "code.google.com/p/go-eco/eco"
)

// Peirce similarity matrix
func PeirceBool_S(data *Matrix) *Matrix {
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = GetABCD(data, i, j)
			v := (a*d - b*c) / ((a + c) * (b + d))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
