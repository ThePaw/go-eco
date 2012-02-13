// Renkonen dissimilarity and similarity matrix
package sim

import (
	. "go-eco.googlecode.com/hg/eco"
	"math"
)

// Renkonen dissimilarity
func Renkonen_D(data *Matrix) *Matrix {

	// recalculate data to proportions
	RecalcToProp(data)

	rows := data.R
	cols := data.C
	out := NewMatrix(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 1.0)
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
			v := 1/sum - 1
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
