// Cody similarity matrix
// Cody (1993)

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

// Cody similarity matrix
func CodyBool_S(data *DenseMatrix) *DenseMatrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	sim := Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = getABCD(data, i, j)
			s := 1 - ((a * (2*a + b + c)) / (2 * (a + b) * (a + c)))
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}
