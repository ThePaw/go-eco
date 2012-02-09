// Braun–Blanquet similarity matrix
// Braun-Blanquet 1932; Magurran 2004.

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	"math"
)

// Braun–Blanquet similarity
func BraunBlanquetBool_S(data *DenseMatrix) *DenseMatrix {
	var (
		a, b, c float64
	)

	warnIfNotBool(data)

	rows := data.Rows()
	out := Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = getABCD(data, i, j)
			v := a/math.Max(b+a, c+a)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

