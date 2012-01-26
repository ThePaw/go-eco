// Faith similarity

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	"math"
)

// Faith similarity matrix
// Faith (1983)
// Legendre & Legendre (1998): 258, eq. 7.18 (S26 index)
func FaithBool_S(data *DenseMatrix) *DenseMatrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	sim := Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = getABCD(data, i, j)
			s := (a + d/2)/ (a+b+c+d)
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}


