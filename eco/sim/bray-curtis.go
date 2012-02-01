// Bray–Curtis distance
// d[jk] = (sum abs(x[ij]-x[ik])/(sum (x[ij]+x[ik]))
// Bray JR, Curtis JT (1957) An ordination of the upland forest communities in southern Wisconsin. Ecol. Monogr. 27:325-349.

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	. "math"
)

// Bray–Curtis distance matrix
func BrayCurtis_D(data *DenseMatrix) *DenseMatrix {
	rows := data.Rows()
	cols := data.Cols()
	out := Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sum1 := 0.0
			sum2 := 0.0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				sum1 += Abs(x - y)
				sum2 += x + y
			}
			v := sum1 / sum2
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

func BrayCurtisBool_D(data *DenseMatrix) *DenseMatrix {
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
			v := (b + c) / (2.0 * (a + b + c)) // ???
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
