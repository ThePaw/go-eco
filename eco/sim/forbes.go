// Forbes similarity matrix
// Forbes (1925), Shi (1993)

package sim

import (
	. "go-eco.googlecode.com/hg/eco"
	"math"
)

// Forbes similarity matrix
func ForbesBool_S(data *Matrix) *Matrix {
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = GetABCD(data, i, j)
			v := (a*(a+b+c+d) - (2 * math.Max(a+b, a+c))) / (((a + b + c + d) * math.Min(a+b, a+c)) - (2 * math.Max(a+b, a+c)))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
