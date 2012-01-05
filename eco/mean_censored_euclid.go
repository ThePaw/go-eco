// Mean Censored Euclidean distance and similarity
//

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	. "math"
)

// Mean Censored Euclidean distance matrix
func MeanCensoredEuclid_D(data *DenseMatrix) *DenseMatrix {
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
			sum := 0.0
			nonzero := 0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				sum += (x - y) * (x - y)
				if x != 0.0 || y != 0.0 {
					nonzero++
				}
			}
			d := Sqrt(sum / float64(nonzero))
			dis.Set(i, j, d)
			dis.Set(j, i, d)
		}
	}
	return dis
}

// Mean Censored Euclidean similarity matrix
// If d denotes Mean Censored Euclidean distance, similarity is s=1.00/(d+1), so that it is in [0, 1]
func MeanCensoredEuclid_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim, dis *DenseMatrix
	)

	dis = MeanCensoredEuclid_D(data)
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
