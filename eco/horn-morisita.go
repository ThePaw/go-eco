// Horn-Morisita distance and similarity
// Similarity is 1.00-d

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

// Horn-Morisita distance matrix
func HornMorisita_D(data *DenseMatrix) *DenseMatrix {
	var (
		dis *DenseMatrix
	)

	rows := data.Rows()
	cols := data.Cols()
	dis = Zeros(rows, rows)

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
				λx += x*x
				λy += y*y
			}

			d := 1 - 2*sumXY/(λx/sumX/sumX + λy/sumY/sumY)/sumX/sumY;
			if d < 0 {
				d=0.0
			}
			dis.Set(i, j, d)
			dis.Set(j, i, d)
		}
	}
	return dis
}

// Horn-Morisita similarity matrix
func HornMorisita_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim, dis *DenseMatrix
	)

	dis = HornMorisita_D(data)
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
