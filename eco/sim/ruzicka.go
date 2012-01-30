// Růžička distance and similarity

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

// Růžička distance matrix
func Ruzicka_D(data *DenseMatrix) *DenseMatrix {
	var (
		aa, bb, jj float64
		out        *DenseMatrix
	)

	rows := data.Rows()
	out = Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			aa, bb, jj, _ = getABJPmin(data, i, j)
			// (A+B-2*J)/(A+B-J)
			v := (aa + bb - 2*jj) / (aa + bb - jj)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
