// Stiles similarity matrix
// Stiles (1946)

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	"math"
)

// Stiles similarity matrix
func StilesBool_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim        *DenseMatrix
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	sim = Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = getABCD(data, i, j)
			t1 := a + b + c + d
			t2 := math.Abs(a*d - b*c)
			t3 := (a + b) * (a + c) * (b + d) * (c + d)
			s := math.Log(t1 * (t2 - t1/2) * (t2 - t1/2) / t3)
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}

//log(((a+b+c+d) * (( abs((a*d)-(b*c)) - ( (a+b+c+d) / 2))^2) / ((a+b)*(a+c) *(b+d)*(c+d))))##
