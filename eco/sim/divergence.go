// Divergence dissimilarity matrix
// Ellis et al. (1993)

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	"math"
)

// Divergence dissimilarity matrix
// Ellis et al. (1993)

func DivergenceBool_D(data *DenseMatrix) *DenseMatrix {
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	out := Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = getABCD(data, i, j)
			v := (math.Sqrt(b+c) / math.Sqrt(a+b+c+d))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
