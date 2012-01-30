package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

func SizeDiffBool_D(data *DenseMatrix) *DenseMatrix {
	var (
		out        *DenseMatrix
		a, b, c, d int64
	)

	rows := data.Rows()
	cols := data.Cols()
	out = Zeros(rows, rows)
	a = 0
	b = 0
	c = 0
	d = 0

	warnIfNotBool(data)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
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
			v := float64(b+c) * float64(b+c) / float64(cols*cols)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
