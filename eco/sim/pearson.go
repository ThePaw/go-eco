// Pearson rho correlations as similarity matrix
package sim

import (
	. "go-eco.googlecode.com/hg/eco"
	"math"
)

// Pearson's ρ (rho) similarity matrix
func PearsonRho_S(data *Matrix) *Matrix {
	rows := data.R
	cols := data.C
	out := NewMatrix(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 1.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sxx := 0.0
			syy := 0.0
			sxy := 0.0
			xmean := 0.0
			ymean := 0.0

			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				xmean += x
				ymean += y
			}
			xmean /= float64(cols)
			ymean /= float64(cols)

			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				sxx += x - xmean
				syy += y - ymean
				sxy += (x - xmean) * (y - ymean)
			}
			v := sxy / math.Sqrt(sxx*syy)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Pearson's Φ similarity matrix
// Phi of Pearson, Gower & Legendre (1986), Yule (1912)
// !!! CHECK against L&L 1998 !!!
func PearsonPhiBool_S(data *Matrix) *Matrix {
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = GetABCD(data, i, j)
			v := (a*d - b*c) / math.Sqrt((a+b)*(a+c)*(d+b)*(d+c))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
