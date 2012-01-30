// Rogers - Tanimoto similarity matrix
// Rogers & Tanimoto (1960), Gower & Legendre (1986)
// Legendre & Legendre 1998: 255

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

// Rogers - Tanimoto similarity matrix
func RogersTanimotoBool_S(data *DenseMatrix) *DenseMatrix {
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	out := Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = getABCD(data, i, j)
			v := (a + d) / (a + 2*(b+c) + d)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
