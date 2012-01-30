// Horn-Morisita distance and similarity
// Similarity is 1.00-d

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

// Horn-Morisita distance matrix
func HornMorisita_D(data *DenseMatrix) *DenseMatrix {
	rows := data.Rows()
	cols := data.Cols()
	out := Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sumXY := 0.0
			sumX := 0.0
			sumY := 0.0
			λx := 0.0
			λy := 0.0

			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				sumXY += x * y
				sumX += x
				sumY += y
				λx += x * x
				λy += y * y
			}

			v := 1 - 2*sumXY/(λx/sumX/sumX+λy/sumY/sumY)/sumX/sumY
			if v < 0 {
				v = 0.0
			}
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

func HornMorisitaBool_D(data *DenseMatrix) *DenseMatrix {
	return BrayCurtisBool_D(data)
}
