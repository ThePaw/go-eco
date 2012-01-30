// Manhattan distance
// Also known as rectilinear distance, Minkowski's L1 distance, taxicab metric, or city-block distance (metric). 

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	. "math"
)

// Manhattan distance
// Legendre & Legendre (1998): 282, eq. 7.45 (D7 index)
func Manhattan_D(data *DenseMatrix) *DenseMatrix {
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
				sum += Abs(x - y)
			}
			out.Set(i, j, sum)
			out.Set(j, i, sum)
		}
	}
	return out
}

// Boolean Manhattan dissimilarity
func ManhattanBool_D(data *DenseMatrix) *DenseMatrix {
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	out := Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = getABCD(data, i, j)
			v := (b + c) / (a + b + c + d)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
