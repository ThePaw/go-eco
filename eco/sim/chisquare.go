// Chi - Squared similarity matrix
// Yule & Kendall (1950)

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

// Chi - Squared similarity matrix
func ChiSquaredBool_S(data *DenseMatrix) *DenseMatrix {
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	sim := Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = getABCD(data, i, j)
			t1 := a + b + c + d
			t2 := (a*d - b*c)
			t3 := (a + b) * (a + c) * (b + d) * (c + d)
			// ((a+b+c+d)*((a*d) - (b*c))^2) / ((a+b)*(a+c)*(b+d)*(c+d))
			s := t1 * t2 * t2 / t3
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}
