// Weiher dissimilarity matrix
// Weiher & Boylen (1994)

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

// Weiher dissimilarity matrix
func WeiherBool_D(data *DenseMatrix) *DenseMatrix {
	var (
		b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	out := Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			_, b, c, _ = getABCD(data, i, j)
			v := b + c
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
