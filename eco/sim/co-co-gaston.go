// Colwell - Coddington - Gaston similarity
// Colwell & Coddington (1948), Gaston et al. (2001)

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

// // Colwell - Coddington - Gaston et al.  similarity matrix
func CoCoGastonBool_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim     *DenseMatrix
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	sim = Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = getABCD(data, i, j)
			s := ((a + b + c) / ((2*a + b + c) / 2)) - 1
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}
