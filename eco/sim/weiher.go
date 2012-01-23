// Weiher similarity matrix
// Weiher & Boylen (1994)

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

// Weiher similarity matrix
func WeiherBool_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim  *DenseMatrix
		b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	sim = Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			_, b, c, _ = getABCD(data, i, j)
			s := b + c
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}
