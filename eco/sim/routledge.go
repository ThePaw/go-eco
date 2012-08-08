// Routledge dissimilarity matrices
// Routledge (1977), Magurran (1988), Wilson & Shmida (1984)

package sim

import (
	. "code.google.com/p/go-eco/eco"
	"math"
)

// Routledge dissimilarity matrix #1
// Routledge (1977), Magurran (1988)
func Routledge1Bool_D(data *Matrix) *Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = GetABCD(data, i, j)
			abc2 := (a + b + c) * (a + b + c)
			v := abc2/(abc2-2*b*c) - 1
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Routledge dissimilarity matrix #2
// Routledge (1977), Wilson & Shmida (1984)
func Routledge2Bool_D(data *Matrix) *Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = GetABCD(data, i, j)
			v := math.Log(2*a+b+c) - ((1 / (2*a + b + c)) * 2 * a * math.Log(2)) - ((1 / (2*a + b + c)) * ((a+b)*math.Log(a+b) + (a+c)*math.Log(a+c)))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Routledge dissimilarity matrix #3
// Routledge (1977)
func Routledge3Bool_D(data *Matrix) *Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = GetABCD(data, i, j)
			v := math.Log(2*a+b+c) - ((1 / (2*a + b + c)) * 2 * a * math.Log(2)) - ((1 / (2*a + b + c)) * ((a+b)*math.Log(a+b) + (a+c)*math.Log(a+c)))
			v = math.Exp(v) - 1
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
