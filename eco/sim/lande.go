// Lande dissimilarity

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

// Lande dissimilarity matrix
func LandeBool_D(data *DenseMatrix) *DenseMatrix {
	var (
		b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	dis := Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			_, b, c, _ = getABCD(data, i, j)
			v := (b + c)/2
			dis.Set(i, j, v)
			dis.Set(j, i, v)
		}
	}
	return dis
}
