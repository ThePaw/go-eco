// Harte similarity matrix
// Harte & Kinzig (1997), Koleff et al. (2003)

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

// Harte similarity matrix
func HarteBool_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim     *DenseMatrix
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	sim = Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = getABCD(data, i, j)
			s := 1 - (2 * a / (2*a + b + c))
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}
