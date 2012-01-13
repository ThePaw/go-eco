// Cosine complement distance and similarity

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	"math"
)

// Cosine complement distance matrix, for boolean data
func Cosine_D(data *DenseMatrix) *DenseMatrix {
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
			aa, bb, jj, _ = getABJPquad(data, i, j)
			// 1-J/sqrt(A*B)
			d := 1.0 - jj/math.Sqrt(aa*bb)
			dis.Set(i, j, d)
			dis.Set(j, i, d)
		}
	}
	return dis
}
