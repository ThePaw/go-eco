// Eyraud dissimilarity matrix
// Eyraud (1936) in Shi (1993)
// Warning: it gives values near zero for both identical, and complementary data!!! STRANGE!
package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

// Eyraud dissimilarity matrix
func EyraudBool_D(data *DenseMatrix) *DenseMatrix {
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	dis := Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = getABCD(data, i, j)
			v := (a - ((a + b) * (a + c))) / ((a + b) * (a + c) * (b + d) * (c + d))
			dis.Set(i, j, v)
			dis.Set(j, i, v)
		}
	}
	return dis
}
