// Legendre similarity matrix
// Gower & Legendre (1986), Russell/Rao in Ellis et al. (1993), Legendre & Legendre (1998)

package sim

import (
	. "code.google.com/p/go-eco/eco"
)

// Legendre similarity matrix #1
// Gower & Legendre (1986), Russell/Rao in Ellis et al. (1993); Russel & Rao (1940)
// Legendre & Legendre (1998): 257, eq. 7.14 (S11 index)

func Legendre1Bool_S(data *Matrix) *Matrix {
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = GetABCD(data, i, j)
			v := a / (a + b + c + d)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Legendre similarity matrix #2
// Legendre & Legendre (1998)
func Legendre2Bool_S(data *Matrix) *Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = GetABCD(data, i, j)
			v := (3 * a) / ((3 * a) + b + c)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
