// Chebychev distance and similarity
// Chebychev distance is a special case of the Minkowski metric, where p = ∞

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	. "math"
)

// Chebychev distance matrix
func Chebychev_D(data *DenseMatrix) *DenseMatrix {
	var (
		dis *DenseMatrix
	)

	rows := data.Rows()
	cols := data.Cols()
	dis = Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		dis.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			d := 0.0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				d = Max(d, Abs(x-y))

			}
			dis.Set(i, j, d)
			dis.Set(j, i, d)
		}
	}
	return dis
}

// Chebychev similarity matrix
// If d denotes Chebychev distance, similarity is s=1.00/(d+1), so that it is in [0, 1]
func MinkowskiInfiniteOrder_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim, dis *DenseMatrix
	)

	dis = Chebychev_D(data)
	rows := data.Rows()
	sim = Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		sim.Set(i, i, 1.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			s := 1.00 / (dis.Get(i, j) + 1.0)
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}
