// Chebychev distance
// Chebychev distance is a special case of the Minkowski metric, where p = ∞

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	. "math"
)

// Chebychev distance matrix
func Chebychev_D(data *DenseMatrix) *DenseMatrix {
	rows := data.Rows()
	cols := data.Cols()
	out := Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			v := 0.0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				v = Max(v, Abs(x-y))

			}
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

