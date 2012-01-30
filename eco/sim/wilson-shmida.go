// Wilson - Shmida dissimilarity

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

// Wilson - Shmida dissimilarity matrix
func WilsonShmidaBool_D(data *DenseMatrix) *DenseMatrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	out := Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = getABCD(data, i, j)
			v := (b + c) / (2*a + b + c)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
