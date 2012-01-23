// Cody dissimilarity matrix
// Cody (1993)

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

// Cody dissimilarity matrix
func CodyBool_D(data *DenseMatrix) *DenseMatrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	dis := Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = getABCD(data, i, j)
			v := 1 - ((a * (2*a + b + c)) / (2 * (a + b) * (a + c)))
			dis.Set(i, j, v)
			dis.Set(j, i, v)
		}
	}
	return dis
}
