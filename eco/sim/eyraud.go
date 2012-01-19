// Eyraud similarity matrix
// Eyraud (1936) in Shi (1993)

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

// Eyraud similarity matrix
func EyraudBool_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim           *DenseMatrix
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	sim = Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = getABCD(data, i, j)
			s:= (a - ((a+b)*(a+c))) / ((a+b)*(a+c)*(b+d)*(c+d))
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}

