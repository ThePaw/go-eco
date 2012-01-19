// Peirce similarity matrix
// Peirce (1884)

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

// Peirce similarity matrix
func PeirceBool_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim           *DenseMatrix
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	sim = Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = getABCD(data, i, j)
			s:= (a*d - b*c) / ((a+c)*(b+d))
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}

