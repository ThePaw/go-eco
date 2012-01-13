// Růžička distance and similarity

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

// Růžička distance matrix
func Ruzicka_D(data *DenseMatrix) *DenseMatrix {
	var (
		aa, bb, jj float64
		dis        *DenseMatrix
	)

	rows := data.Rows()
	dis = Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		dis.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			aa, bb, jj, _ = getABJPmin(data, i, j)
			// (A+B-2*J)/(A+B-J)
			d := (aa + bb - 2*jj) / (aa + bb - jj)
			dis.Set(i, j, d)
			dis.Set(j, i, d)
		}
	}
	return dis
}
