// Baroni-Urbani and Buser (dis)similarity matrix
// Baroni-Urbani C, Buser MW (1976) Similarity of binary data. Syst. Zool 25:251-259.

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	"math"
)

// Baroni-Urbani and Buser dissimilarity matrix
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
			sqrtcd := math.Sqrt(float64(c*d))
			d := 1.0 - (sqrtcd + c) / (sqrtcd + a + b + c)
			dis.Set(i, j, d)
			dis.Set(j, i, d)
		}
	}
	return dis
}

// Baroni-Urbani and Buser similarity matrix
func BaroniUrbaniBool_S(data *DenseMatrix) *DenseMatrix {
	dis := BaroniUrbaniBool_D(data)
	//1-D
	return sFromD(dis, 0)
}
