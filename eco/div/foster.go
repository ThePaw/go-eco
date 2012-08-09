// Foster index of poverty

package div

import (
	"code.google.com/p/go-eco/eco/aux"
	"math"
)

// Foster index of poverty
// Foster, J. E. (1984). On Economic Poverty: A Survey of Aggregate Measures. Advances in Econometrics, 3, 215–251.
// Zheng, B. (1997). Aggregate Poverty Measures. Journal of Economic Surveys, 11, 123–162.
func Foster_D(data *aux.Matrix, k, parameter float64) *Vector {
	rows := data.R
	cols := data.C
	out := NewVector(rows)

	for i := 0; i < rows; i++ {
		n := 0.0
		v := 0.0

		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0.0 {
				n++
				if x < k {
					v += math.Pow((k-x)/k, parameter-1)
				}
			}
		}
		if n > 0 {
			v /= n
		}
		out.Set(i, v)
	}
	return out
}
