// Lande similarity

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

// Lande similarity matrix
func LandeBool_S(data *DenseMatrix, which byte) *DenseMatrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	sim := Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = getABCD(data, i, j)
			s:= (b+c)/(2*a+b+c) 
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}

