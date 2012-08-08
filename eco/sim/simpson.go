// Simpson similarity matrix
// Simpson (1960), Shi (1993)

package sim

import (
	. "code.google.com/p/go-eco/eco"
	"math"
)

// Simpson dissimilarity matrix #1
func Simpson1Bool_D(data *Matrix) *Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = GetABCD(data, i, j)
			v := math.Min(b, c) / (math.Min(b, c) + a)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Simpson similarity matrix #2
func Simpson2Bool_S(data *Matrix) *Matrix {
	var (
		a, b float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, _, _ = GetABCD(data, i, j)
			v := a/a + b
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
