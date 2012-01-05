// Bray - Curtis distance and similarity
//

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	. "math"
)

// Bray - Curtis distance matrix
func BrayCurtis_D(data *DenseMatrix) *DenseMatrix {
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
			sum1 := 0.0
			sum2 := 0.0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				sum1 += Abs(x - y)
				sum2 += x + y
			}
			d := sum1 / sum2
			dis.Set(i, j, d)
			dis.Set(j, i, d)
		}
	}
	return dis
}

// Bray - Curtis similarity matrix
// If d denotes Bray - Curtis distance, similarity is s=1.00/(d+1), so that it is in [0, 1]
func BrayCurtis_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim, dis *DenseMatrix
	)

	dis = BrayCurtis_D(data)
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
