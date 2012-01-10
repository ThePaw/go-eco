// Morisita distance and similarity
// d[jk] = 1 - 2*sum(x[ij]*x[ik])/((lambda[j]+lambda[k]) * sum(x[ij])*sum(x[ik]))
// where lambda[j] = sum(x[ij]*(x[ij]-1))/sum(x[ij])*sum(x[ij]-1)
// Can only be used with integer data, and may still
// fail with unfortunate pairs of species occurring only once.
// Similarity is 1.00-d

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

// Morisita distance matrix
func Morisita_D(data *DenseMatrix) *DenseMatrix {
	var (
		dis *DenseMatrix
	)

	rows := data.Rows()
	cols := data.Cols()
	dis = Zeros(rows, rows)
	// check whether data are integers; if not, truncate them
	truncData(data)

	for i := 0; i < rows; i++ {
		dis.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {

			sumXY :=0.0
			sumX := 0.0
			sumY := 0.0
			λx := 0.0
			λy := 0.0

			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				sumXY += x*y
				sumX += x
				sumY += y
				λx += x*(x - 1.0)
				λy += y*(y - 1.0)

			}

			d := 1 - 2*sumXY/(λx/sumX/(sumX-1) + λy/sumY/(sumY-1))/sumX/sumY;
			if d < 0 {
				d=0.0
			}
			dis.Set(i, j, d)
			dis.Set(j, i, d)
		}
	}
	return dis
}

// Morisita similarity matrix
func Morisita_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim, dis *DenseMatrix
	)

	dis = Morisita_D(data)
	rows := data.Rows()
	sim = Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		sim.Set(i, i, 1.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			s := 1.00 - dis.Get(i, j)
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}
