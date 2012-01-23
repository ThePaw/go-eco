package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

func PatternDiffBool_D(data *DenseMatrix) *DenseMatrix {
	var (
		a, b, c, d float64
	)

	warnIfNotBool(data)

	rows := data.Rows()
	dis := Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		dis.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			a, b, c, d = getABCD(data, i, j)
			p:=(a+b+c+d)
			dist := (b*c) / (p*p)
			dis.Set(i, j, dist)
			dis.Set(j, i, dist)
		}
	}
	return dis
}
