// Morisita distance and similarity
// d[jk] = 1 - 2*sum(x[ij]*x[ik])/((lambda[j]+lambda[k]) * sum(x[ij])*sum(x[ik]))
// where lambda[j] = sum(x[ij]*(x[ij]-1))/sum(x[ij])*sum(x[ij]-1)
// Can only be used with integer data, and may still
// fail with unfortunate pairs of species occurring only once.
// Morisita (1959)

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

// Morisita distance matrix
// Morisita (1959)
func Morisita_D(data *DenseMatrix) *DenseMatrix {
	rows := data.Rows()
	cols := data.Cols()
	out := Zeros(rows, rows)

	// check whether data are integers; if not, truncate them
	truncData(data)

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
				λx += x * (x - 1.0)
				λy += y * (y - 1.0)

			}

			v := 1 - 2*sumXY/(λx/sumX/(sumX-1)+λy/sumY/(sumY-1))/sumX/sumY
			if v < 0 {
				v = 0.0
			}
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
