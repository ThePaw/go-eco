// Shannon similarity matrix
package sim

import (
	. "go-eco.googlecode.com/hg/eco"
	"math"
)

func ShannonBool_D(data *Matrix) *Matrix {
	var (
		b, c float64
	)

	WarnIfNotBool(data)

	rows := data.R
	out := NewMatrix(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			_, b, c, _ = GetABCD(data, i, j)
			v := 2.0 * (b + c) * math.Log(2.0)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
