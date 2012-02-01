// Shannon similarity matrix
package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	"math"
)

func ShannonBool_D(data *DenseMatrix) *DenseMatrix {
	var (
		b, c float64
	)

	warnIfNotBool(data)

	rows := data.Rows()
	out := Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			_, b, c, _ = getABCD(data, i, j)
			v := 2.0 * (b + c) * math.Log(2.0)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
