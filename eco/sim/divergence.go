// Divergence dissimilarity matrix
// Ellis et al. (1993)

package sim

import (
	. "go-eco.googlecode.com/hg/eco"
	"math"
)

// Divergence dissimilarity matrix
// Ellis et al. (1993)

func DivergenceBool_D(data *Matrix) *Matrix {
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = GetABCD(data, i, j)
			v := (math.Sqrt(b+c) / math.Sqrt(a+b+c+d))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
