package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

func PatternDiffBool_D(data *DenseMatrix) *DenseMatrix {
	var (
		dis        *DenseMatrix
		a, b, c, d int64
	)

	rows := data.Rows()
	cols := data.Cols()
	dis = Zeros(rows, rows)
	a = 0
	b = 0
	c = 0
	d = 0

	checkIfBool(data)

	for i := 0; i < rows; i++ {
		dis.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)

				switch {
				case x != 0 && y != 0:
					a++
				case x != 0 && y == 0:
					b++
				case x == 0 && y != 0:
					c++
				case x == 0 && y == 0:
					d++
				}

			}
			d := float64(b * c) / float64(cols * cols)
			dis.Set(i, j, d)
			dis.Set(j, i, d)
		}
	}
	return dis
}


