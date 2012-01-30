// Harrison dissimilarity matrix
// Harrison et al. (1992), Koleff et al. (2003)

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	"math"
)

// Harrison dissimilarity matrix
func HarrisonBool_D(data *DenseMatrix) *DenseMatrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	out := Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = getABCD(data, i, j)
			v := math.Min(b,c) / (math.Max(b,c) + a)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
