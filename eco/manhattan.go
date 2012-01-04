// Manhattan distance and similarity
// Also known as rectilinear distance, Minkowski's L1 distance, taxi cab metric, or city block distance. 

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	. "math"
)

func Manhattan_D(data *DenseMatrix) *DenseMatrix {
	var (
		sum float64
		dis *DenseMatrix
	)

	dis = Zeros(data.Rows(), data.Rows()) // square similarity matrix row vs. row
	sum = 0.0

	for i := 0; i < data.Rows(); i++ {
		dis.Set(i, i, 0.0)
	}

	for i := 0; i < data.Rows(); i++ {
		for j := i + 1; j < data.Rows(); j++ {
			sum = 0
			for k := 0; k < data.Cols(); k++ {
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
	sim = Zeros(data.Rows(), data.Rows()) // square similarity matrix row vs. row

	for i := 0; i < data.Rows(); i++ {
		sim.Set(i, i, 1.0)
	}

	for i := 0; i < data.Rows(); i++ {
		for j := i + 1; j < data.Rows(); j++ {
			x := dis.Get(i, j) + 1.0
			sim.Set(i, j, 1.00/x)
			sim.Set(j, i, 1.00/x)
		}
	}
	return sim
}
