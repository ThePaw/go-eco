// Lamont similarity matrix
// Lamont and Grant (1979)

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

// Lamont similarity matrix
func LamontBool_S(data *DenseMatrix) *DenseMatrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	sim := Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = getABCD(data, i, j)
			v := a / (2*a + b + c)
			sim.Set(i, j, v)
			sim.Set(j, i, v)
		}
	}
	return sim
}
