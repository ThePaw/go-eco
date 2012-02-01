// Minkowski distance

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	. "math"
)

// Minkowski distance matrix
// Legendre & Legendre (1998): 281, eq. 7.44 (D6 index)
func Minkowski_D(power int, data *DenseMatrix) *DenseMatrix {
	rows := data.Rows()
	cols := data.Cols()
	out := Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sum := 0.0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				sum += Pow(Abs(x-y), float64(power))
			}
			v := Pow(sum, 1/float64(power))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
