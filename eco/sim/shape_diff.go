// Shape difference distance matrix

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

func ShapeDiffBool_D(data *DenseMatrix) *DenseMatrix {
	var (
		a, b, c, d float64
	)

	warnIfNotBool(data)

	rows := data.Rows()
	out := Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			a, b, c, d = getABCD(data, i, j)
			p:=(a+b+c+d)
			v := (p*(b+c)-(b-c)*(b-c)) / (p*p)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
