// Yule similarity matrix
// Legendre & Legendre 1998: 256, eq. 7.8. 

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

// Yule similarity matrix
func YuleBool_S(data *DenseMatrix) *DenseMatrix {
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	out := Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = getABCD(data, i, j)
			v := (a*d - b*c) / (a*d + b*c)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
