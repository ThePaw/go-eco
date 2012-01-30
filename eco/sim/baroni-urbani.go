// Baroni-Urbani and Buser (dis)similarity matrix
// Baroni-Urbani & Buser (1976), Wolda (1981)

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	"math"
)

// Baroni-Urbani and Buser similarity matrix
func BaroniUrbaniBool_S(data *DenseMatrix) *DenseMatrix {
	var (
		out        *DenseMatrix
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	warnIfNotBool(data)

	rows := data.Rows()
	out = Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = getABCD(data, i, j)
			v := ((math.Sqrt(a * d)) + a) / ((math.Sqrt(a * d)) + b + c + a)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Baroni-Urbani and Buser dissimilarity matrix
// according to R:vegan
func BaroniUrbaniBool_D(data *DenseMatrix) *DenseMatrix {
	var (
		a, b, c, d float64
	)

	warnIfNotBool(data)

	rows := data.Rows()
	out := Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			a, b, c, d = getABCD(data, i, j)
			sqrtcd := math.Sqrt(float64(c * d))
			v := 1.0 - (sqrtcd+c)/(sqrtcd+a+b+c)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

