// Williams dissimilarity matrix
// Williams (1996), Koleff et al. (2003)

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	"math"
)

// Williams dissimilarity matrix #1
func Williams1Bool_D(data *DenseMatrix) *DenseMatrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	dis := Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = getABCD(data, i, j)
			delta := math.Min(b, c) / (a + b + c)
			dis.Set(i, j, delta)
			dis.Set(j, i, delta)
		}
	}
	return dis
}

// Williams dissimilarity matrix #2
func Williams2Bool_D(data *DenseMatrix) *DenseMatrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	dis := Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = getABCD(data, i, j)
//			delta := (2*b*c + 1) / (math.Pow(a+b+c, 2) - (a + b + c))
			delta := ((b*c)+1) / ((((a+b+c)*(a+b+c)) - (a+b+c)) / 2)
			dis.Set(i, j, delta)
			dis.Set(j, i, delta)
		}
	}
	return dis
}
