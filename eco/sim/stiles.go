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
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	out := Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = getABCD(data, i, j)
			t1 := a + b + c + d
			t2 := math.Abs(a*d - b*c)
			t3 := (a + b) * (a + c) * (b + d) * (c + d)
			v := math.Log(t1 * (t2 - t1/2) * (t2 - t1/2) / t3)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
