// Lennon dissimilarity matrix
// Lennon et al. (2001), Koleff et al. (2003)

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	"math"
)

// Lennon dissimilarity matrix #1
func Lennon1Bool_D(data *DenseMatrix) *DenseMatrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	out := Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = getABCD(data, i, j)
			v := (2 * math.Abs(b-c)) / (2*a + b + c)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Lennon dissimilarity matrix #2
func Lennon2Bool_D(data *DenseMatrix) *DenseMatrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	out := Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = getABCD(data, i, j)
			v := 1 - (math.Log((2*a+b+c)/(a+b+c)) / math.Log(2))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
