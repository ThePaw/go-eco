// Similarity ratio matrix
package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

func SimRatio_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim *DenseMatrix
	)
	rows := data.Rows()
	cols := data.Cols()
	sim = Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		sim.Set(i, i, 1.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sum1 := 0.0
			sum2 := 0.0
			sum3 := 0.0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				sum1 += x * x
				sum2 += y * y
				sum3 += x * y
			}
			s := sum3 / (sum1 + sum2 - sum3)
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}
