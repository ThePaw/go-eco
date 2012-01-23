// Kulczynski similarity and distance matrix
// Oosting (1956), Southwood (1978)

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	. "math"
)

// Kulczynski similarity matrix #1
func Kulczynski1Bool_S(data *DenseMatrix) *DenseMatrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	sim := Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = getABCD(data, i, j)
			s := a / (b + c)
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}

// Kulczynski similarity matrix #2
func Kulczynski2Bool_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim     *DenseMatrix
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	sim = Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = getABCD(data, i, j)
			s := ((a / 2) * ((2 * a) + b + c)) / ((a + b) * (a + c))
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}

// Kulczynski distance matrix
func Kulczynski_D(data *DenseMatrix) *DenseMatrix {
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
			sumMin := 0.0
			sumX := 0.0
			sumY := 0.0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				sumMin += Min(x, y)
				sumX += x
				sumY += x
			}
			d := 1 - 0.5*(sumMin/sumX+sumMin/sumY)
			dis.Set(i, j, d)
			dis.Set(j, i, d)
		}
	}
	return dis
}
