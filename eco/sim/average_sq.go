// Squared average distance matrix
package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

// Squared average distance
func AverageSqBool_D(data *DenseMatrix) *DenseMatrix {
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	dis := Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		dis.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = getABCD(data, i, j)
			delta := (b + c) / (a+b+c+d)
			dis.Set(i, j, delta)
			dis.Set(j, i, delta)
		}
	}
	return dis
}
