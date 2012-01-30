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
		out        *DenseMatrix
	)

	rows := data.Rows()
	out = Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			aa, bb, jj, _ = getABJPquad(data, i, j)
			// 1-J/sqrt(A*B)
			v := 1.0 - jj/math.Sqrt(aa*bb)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
