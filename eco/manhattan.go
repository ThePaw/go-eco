// Manhattan distance and similarity
// Also known as rectilinear distance, Minkowski's L1 distance, taxi cab metric, or city block distance. 

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	. "math"
)

func Manhattan_D(data *DenseMatrix) *DenseMatrix {
	var (
		dis *DenseMatrix
	)

	rows := data.Rows()
	cols := data.Cols()
	dis = Zeros(rows, rows) // square similarity matrix row vs. row

	for i := 0; i < rows; i++ {
		dis.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sum := 0.0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				sum += Abs(x - y)
			}
			dis.Set(i, j, sum)
			dis.Set(j, i, sum)
		}
	}
	return dis
}

func Manhattan_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim, dis *DenseMatrix
	)

	dis = Manhattan_D(data)
	rows := data.Rows()
	sim = Zeros(rows, rows) // square similarity matrix row vs. row

	for i := 0; i < rows; i++ {
		sim.Set(i, i, 1.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			x := dis.Get(i, j) + 1.0
			sim.Set(i, j, 1.00/x)
			sim.Set(j, i, 1.00/x)
		}
	}
	return sim
}
