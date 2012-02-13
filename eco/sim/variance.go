package sim

import (
	. "go-eco.googlecode.com/hg/eco"
)

func VarianceBool_D(data *Matrix) *Matrix {
	var (
		out        *Matrix
		a, b, c, d int64
	)

	rows := data.R
	cols := data.C
	out = NewMatrix(rows, rows)
	a = 0
	b = 0
	c = 0
	d = 0

	WarnIfNotBool(data)

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
			v := float64(b+c) / (4.0 * float64(cols))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
