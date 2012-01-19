// Harrison similarity matrix
// Harrison et al. (1992), Koleff et al. (2003)

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

// Harrison similarity matrix
func HarrisonBool_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim           *DenseMatrix
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	sim = Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = getABCD(data, i, j)
			s:= (2*a+b+c)*(1-(a/(a+b+c)))
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}

