// Hamann similarity matrix
// Holley JW, Guilford JP 1964 A note on the G index of agreement. Educational and Psychological Measurement, 24(7):749-753.

package sim

import (
	. "go-eco.googlecode.com/hg/eco"
)

// Hamann similarity matrix
// Legendre & Legendre 1998: 256, eq. 7.7. 
// S9 index of Gower & Legendre (1986)
// S6 index of R:ade4:dist.binary

func HamannBool_S(data *Matrix) *Matrix {
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	WarnIfNotBool(data)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = GetABCD(data, i, j)
			v := (a + d - b - c) / (a + b + c + d)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
