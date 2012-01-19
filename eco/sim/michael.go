// Michael similarity matrix
// Michael (1920), Shi (1993)

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

// Michael similarity matrix
func MichaelBool_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim           *DenseMatrix
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	sim = Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = getABCD(data, i, j)
			s:= (4*((a*d) - (b*c))) / ((a+d)*(a+d) + (b+c)*(b+c))
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}

