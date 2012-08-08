// Russell-Rao similarity matrix 
// Russel & Rao (1940)
// = legendre of R:simba
// Gower & Legendre (1986), Russell/Rao in Ellis et al. (1993)
// Legendre & Legendre (1998): 257, eq. 7.14, S11

package sim

import (
	. "code.google.com/p/go-eco/eco"
)

// Russell-Rao similarity matrix 
// Legendre & Legendre (1998): 257, eq. 7.14 (S11 index)
func RusselRaoBool_S(data *Matrix) *Matrix {
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
