// Williams similarity matrix
// Williams (1996), Koleff et al. (2003)

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	"math"
)

// Williams similarity matrix #1
func Williams1Bool_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim           *DenseMatrix
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	sim = Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = getABCD(data, i, j)
			s:= math.Min(b,c) / (a+b+c) 
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}

// Williams similarity matrix #2
func Williams2Bool_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim           *DenseMatrix
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	sim = Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = getABCD(data, i, j)
			s:= (2*b*c+1) / (math.Pow(a+b+c, 2) - (a+b+c)) 
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}

