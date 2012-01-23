// Legendre similarity matrix
// Gower & Legendre (1986), Russell/Rao in Ellis et al. (1993), Legendre & Legendre (1998)

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

// Legendre similarity matrix #1
// Gower & Legendre (1986), Russell/Rao in Ellis et al. (1993)
func Legendre1Bool_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim        *DenseMatrix
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	sim = Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = getABCD(data, i, j)
			s := a / (a + b + c + d)
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}

// Legendre similarity matrix #2
// Legendre & Legendre (1998)
func Legendre2Bool_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim     *DenseMatrix
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	sim = Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = getABCD(data, i, j)
			s := (3 * a) / ((3 * a) + b + c)
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}
