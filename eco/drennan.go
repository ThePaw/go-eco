// Drennan distance and similarity
// Marquardt, W.H. 1978 Archaeological seriation. In: Schiffer, M.B.(ed.)
// Advances in Archaeological Method and Theory. Academic Press, N.Y., p.284.
// Drennan, R.D. 1976 A refinement of chronological seriation using nonmetric
// multidimensional scaling. American antiquity, 41: 290-302.

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

// Drennan distance matrix
func Drennan_D(data *DenseMatrix) *DenseMatrix {
	var (
		dis, percent *DenseMatrix
	)
	rows := data.Rows()
	cols := data.Cols()
	percent = Zeros(rows, cols) // percentages
	dis = Zeros(rows, rows)     // distances

	for i := 0; i < rows; i++ {
		rowsum := 0.0
		for j := i + 1; j < cols; j++ {
			rowsum += data.Get(i, j)
		}
		for j := i + 1; j < cols; j++ {
			percent.Set(i, j, data.Get(i, j)*100.0/rowsum)
		}
	}

	for i := 0; i < rows; i++ {
		dis.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sum := 0.0
			for k := 0; k < cols; k++ {
				x := percent.Get(i, k)
				y := percent.Get(j, k)
				sum += (x - y)
			}
			d := sum / 200.0
			dis.Set(i, j, d)
			dis.Set(j, i, d)
		}
	}
	return dis
}

// Drennanean similarity matrix
func Drennan_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim, dis *DenseMatrix
	)

	dis = Drennan_D(data)
	rows := data.Rows()
	sim = Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		sim.Set(i, i, 1.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			s := 1.00 / (dis.Get(i, j) + 1.0)
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}
