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
	dis := Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		dis.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sum := 0.0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				sum += Abs(x - y)
			}
			dis.Set(i, j, sum)
			dis.Set(j, i, sum)
		}
	}
	return dis
}

// Boolean Manhattan dissimilarity
func ManhattanBool_D(data *DenseMatrix) *DenseMatrix {
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	dis := Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = getABCD(data, i, j)
			delta := (b + c) / (a + b + c + d)
			dis.Set(i, j, delta)
			dis.Set(j, i, delta)
		}
	}
	return dis
}
