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
		sim           *DenseMatrix
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	sim = Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = getABCD(data, i, j)
			s:= ((math.Sqrt(a*d))+a) / ((math.Sqrt(a*d))+b+c+a)
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}

// Baroni-Urbani and Buser dissimilarity matrix
// according to R:vegan
func BaroniUrbaniBool_D(data *DenseMatrix) *DenseMatrix {
	var (
		dis        *DenseMatrix
		a, b, c, d float64
	)

	rows := data.Rows()
	cols := data.Cols()
	dis = Zeros(rows, rows)
	a = 0
	b = 0
	c = 0
	d = 0

	checkIfBool(data)

	for i := 0; i < rows; i++ {
		dis.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)

				switch {
				case x != 0 && y != 0:
					a++
				case x != 0 && y == 0:
					b++
				case x == 0 && y != 0:
					c++
				case x == 0 && y == 0:
					d++
				}

			}
			sqrtcd := math.Sqrt(float64(c * d))
			d := 1.0 - (sqrtcd+c)/(sqrtcd+a+b+c)
			dis.Set(i, j, d)
			dis.Set(j, i, d)
		}
	}
	return dis
}

