// Harrison dissimilarity matrix
// Harrison et al. (1992), Koleff et al. (2003)

package sim

import (
	. "go-eco.googlecode.com/hg/eco"
	"math"
)

// Harrison dissimilarity matrix
func HarrisonBool_D(data *Matrix) *Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = GetABCD(data, i, j)
			v := math.Min(b, c) / (math.Max(b, c) + a)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
