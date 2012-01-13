// Renkonen dissimilarity and similarity matrix
package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	"math"
)

// Renkonen dissimilarity
func Renkonen_D(data *DenseMatrix) *DenseMatrix {
	
	// recalculate data to proportions
	recalcToProp(data) 

	rows := data.Rows()
	cols := data.Cols()
	dis := Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		dis.Set(i, i, 1.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sum := 0.0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				sum += math.Min(x, y)
			}
			// original formula: Σ min(p1, p2)
			// d = 1/s - 1
			d := 1/sum - 1
			dis.Set(i, j, d)
			dis.Set(j, i, d)
		}
	}
	return dis
}

// Renkonen similarity
func Renkonen_S(data *DenseMatrix) *DenseMatrix {
	dis := Renkonen_D(data)
	// s = 1/(D+1) type
	return sFromD(dis, 1)
}

