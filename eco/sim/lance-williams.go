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
	dis := Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		dis.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			a, b, c, _ = getABCD(data, i, j)
			dist := (b+c) / (2 * (a+b+c))
			dis.Set(i, j, dist)
			dis.Set(j, i, dist)
		}
	}
	return dis
}
