// Forbes similarity matrix
// Forbes (1925), Shi (1993)

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	"math"
)

// Forbes similarity matrix
func ForbesBool_S(data *DenseMatrix) *DenseMatrix {
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	sim := Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = getABCD(data, i, j)
			s := (a*(a+b+c+d) - (2 * math.Max(a+b, a+c))) / (((a + b + c + d) * math.Min(a+b, a+c)) - (2 * math.Max(a+b, a+c)))
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}
