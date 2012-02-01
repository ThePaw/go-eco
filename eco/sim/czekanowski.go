// Czekanowski dissimilarity
package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	"math"
)

// Czekanowski dissimilarity matrix #1
// Czekanowski (1909)
// Legendre & Legendre (1998): 282, eq. 7.46 (D8 index)
func Czekanowski1_D(data *DenseMatrix) *DenseMatrix {
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
				sum += math.Abs(x - y)
			}
			v := sum / float64(cols)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Czekanowski dissimilarity matrix #2
// Reference needed !
func Czekanowski2_D(data *DenseMatrix) *DenseMatrix {
	rows := data.Rows()
	cols := data.Cols()
	out := Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 1.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sumXX := 0.0
			sumYY := 0.0
			sumXY := 0.0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				sumXX += x * x
				sumYY += y * y
				sumXY += x * y
			}
			v := 1 - 200.0*sumXX/(sumYY+sumXY)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
