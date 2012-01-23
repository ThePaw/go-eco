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
	dis := Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		dis.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			_, b, c, _ = getABCD(data, i, j)
			dist := 2.0 * (b+c) * math.Log(2.0)
			dis.Set(i, j, dist)
			dis.Set(j, i, dist)
		}
	}
	return dis
}
