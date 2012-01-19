// Johnson similarity matrix
// Johnson (1971), Johnson (1967)

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

// Johnson similarity matrix #1
// Johnson (1971)
func Johnson1Bool_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim           *DenseMatrix
		a, b float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	sim = Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, _, _ = getABCD(data, i, j)
			s:= a / (2*b)
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}

// Johnson similarity matrix #2
// Johnson (1967)
func Johnson2Bool_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim           *DenseMatrix
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	sim = Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = getABCD(data, i, j)
			s:= (a/(a+b)) + (a/(a+c))
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}



