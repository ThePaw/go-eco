// Hamann similarity matrix
// Holley JW, Guilford JP 1964 A note on the G index of agreement. Educational and Psychological Measurement, 24(7):749-753.
// Legendre & Legendre 1998: 256, eq. 7.7. 

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

// Hamann similarity matrix
func HamannBool_S(data *DenseMatrix) *DenseMatrix {
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	warnIfNotBool(data)

	rows := data.Rows()
	out := Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = getABCD(data, i, j)
			v := (a + d - b - c) / (a + b + c + d)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
