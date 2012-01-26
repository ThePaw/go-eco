// Steinhaus similarity
// Motyka (1947)
// Legendre & Legendre (1998): 265, eq. 7.24 (S17 index)
// for count or interval data

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	"math"
)

// Steinhaus similarity matrix
func Steinhaus_S(data *DenseMatrix) *DenseMatrix {
	rows := data.Rows()
	cols := data.Cols()
	out := Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sumX := 0.0
			sumY := 0.0
			sumMin := 0.0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				sumX += x
				sumY += y
				sumMin += math.Min(x, y)
			}
			v := 2*sumMin(sumX+sumY)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}


