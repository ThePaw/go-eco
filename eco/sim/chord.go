// Chord distance
// Orloci (1967b)
// Legendre & Legendre (1998): 279, eq. 7.37 (D3 index)

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	"math"
)

// Chord distance matrix, float data
func Chord_D(data *DenseMatrix) *DenseMatrix {
	rows := data.Rows()
	cols := data.Cols()
	out := Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
				sumXY :=0
				sumXX :=0
				sumYY :=0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				sumXY += x * y
				sumXX += x*x
				sumYY += y*y
			}
			v := math.Sqrt(2*(1 - (sumXY/math.Sqrt(sumXX*sumYY)))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}


