// Lennon similarity matrix
// Lennon et al. (2001), Koleff et al. (2003)

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	"math"
)

// Lennon similarity matrix #1
func Lennon1Bool_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim           *DenseMatrix
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	sim = Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = getABCD(data, i, j)
			s:= (2 * math.Abs(b-c)) / (2*a + b +c)
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}
// Lennon similarity matrix #2
func Lennon2Bool_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim           *DenseMatrix
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	sim = Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = getABCD(data, i, j)
			s:= 1 - (math.Log((2*a + b + c)/(a + b + c)) / math.Log(2))
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}

