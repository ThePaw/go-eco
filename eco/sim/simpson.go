// Simpson similarity matrix
// Simpson (1960), Shi (1993)

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	"math"
)

// Simpson dissimilarity matrix #1
func Simpson1Bool_D(data *DenseMatrix) *DenseMatrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	dis := Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = getABCD(data, i, j)
			delta := math.Min(b, c) / (math.Min(b, c) + a)
			dis.Set(i, j, delta)
			dis.Set(j, i, delta)
		}
	}
	return dis
}

// Simpson similarity matrix #2
func Simpson2Bool_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim  *DenseMatrix
		a, b float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	sim = Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, _, _ = getABCD(data, i, j)
			s := a/a + b
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}
