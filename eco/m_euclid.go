// Mean Euclidean distance and similarity

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

// Mean Euclidean distance matrix
func MeanEuclid_D(data *DenseMatrix) *DenseMatrix {
	var (
		dis *DenseMatrix
	)

	dis = Euclid_D(data)
	rows := dis.Rows()
	for i := 0; i < rows; i++ {
		for j := i + 1; j < dis.Cols(); j++ {
			d := dis.Get(i, j) / float64(rows)
			dis.Set(i, j, d)
			dis.Set(j, i, d)
		}
	}
	return dis
}

// Mean Euclidean similarity matrix
// If d denotes Mean Euclidean distance, similarity is s=1.00/(d+1), so that it is in [0, 1]
func MeanEuclid_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim, dis *DenseMatrix
	)

	dis = MeanEuclid_D(data)
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
