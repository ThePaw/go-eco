// Weiher dissimilarity matrix
// Weiher & Boylen (1994)

package sim

import (
	. "go-eco.googlecode.com/hg/eco"
)

// Weiher dissimilarity matrix
func WeiherBool_D(data *Matrix) *Matrix {
	var (
		b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			_, b, c, _ = GetABCD(data, i, j)
			v := b + c
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
