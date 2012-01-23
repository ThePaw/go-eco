// Pearson rho correlations as similarity matrix
package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	"math"
)

// Pearson's ρ (rho) similarity matrix
func PearsonRho_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim *DenseMatrix
	)
	rows := data.Rows()
	cols := data.Cols()
	sim = Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		sim.Set(i, i, 1.0)
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
			s := sxy / math.Sqrt(sxx*syy)
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}

// Pearson's Φ similarity matrix
// Phi of Pearson, Gower & Legendre (1986), Yule (1912)
// !!! CHECK against L&L 1998 !!!
func PearsonPhiBool_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim        *DenseMatrix
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	sim = Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = getABCD(data, i, j)
			s := (a*d - b*c) / math.Sqrt((a+b)*(a+c)*(d+b)*(d+c))
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}
