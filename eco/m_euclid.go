// Euclidean distance and similarity

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

func MeanEuclid_D(data *DenseMatrix) *DenseMatrix {
	var (
		dis *DenseMatrix
	)

	dis = Euclid_D(data)
	rows := dis.Rows()
	for i := 0; i < rows; i++ {
		for j := i + 1; j < dis.Cols(); j++ {
			x := dis.Get(i, j)
			dis.Set(i, j, x/float64(rows))
			dis.Set(j, i, x/float64(rows))
		}
	}
	return dis
}
