// Euclidean distance and similarity

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	"math"
)

func MeanEuclid_D(data *DenseMatrix)  *DenseMatrix {
	var (
		sum float64
		dis *DenseMatrix
	)

	dis = Euclid_D(data)
	sum = 0.0

	for i := 0; i < dis.Rows(); i++	{
		for j := i + 1; j < dis.Cols(); j++ {
			x = dis.Get(i, j)
			dis.Set(i, j, x/dis.Rows())
			dis.Set(j, i, x/dis.Rows())
		}
	}
	return dis
}


