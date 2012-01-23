// Harrison dissimilarity matrix
// Harrison et al. (1992), Koleff et al. (2003)

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

// Harrison dissimilarity matrix
func HarrisonBool_D(data *DenseMatrix) *DenseMatrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	dis := Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = getABCD(data, i, j)
			v := (2*a + b + c) * (1 - (a / (a + b + c)))
			dis.Set(i, j, v)
			dis.Set(j, i, v)
		}
	}
	return dis
}
