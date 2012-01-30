// Lance-Williams similarity
package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

func LanceWilliamsBool_D(data *DenseMatrix) *DenseMatrix {
	var (
		a, b, c float64
	)

	warnIfNotBool(data)

	rows := data.Rows()
	out := Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			a, b, c, _ = getABCD(data, i, j)
			v := (b+c) / (2 * (a+b+c))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
