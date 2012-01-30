// Simple matching coefficient similarity matrix. 
// Sokal RR, Michener CD 1958 A statistical method for evaluating systematic relationship. University of Kansas Science Bulletin, 38:1409-1438. 
// Legendre & Legendre 1998: 255, eq. 7.1 (S1 index). 

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

// Simple matching coefficient similarity matrix
func SimpleMatchingBool_S(data *DenseMatrix) *DenseMatrix {
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	out := Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = getABCD(data, i, j)
			v := (a + d) / (a + b + c + d)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
