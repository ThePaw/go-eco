// Dice's (dis)similarity matrix

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

/*
func DiceBool_D(data *DenseMatrix) *DenseMatrix {
	var (
		dis        *DenseMatrix
		a, b, c, d int64
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
			dist := 1.0 - 2.0 * float64(a) / float64(b + c)
			dis.Set(i, j, dist)
			dis.Set(j, i, dist)
		}
	}
	return dis
}
*/

// Dice's dissimilarity
// it is not a proper distance metric as it does not possess the property of triangle inequality
// Dice = 2*Jaccard / (1 + Jaccard) 
func DiceBool_D(data *DenseMatrix) *DenseMatrix {
	var (
		aa, bb, jj float64
		dis        *DenseMatrix
	)

	rows := data.Rows()
	dis = Zeros(rows, rows)
	checkIfBool(data)

	for i := 0; i < rows; i++ {
		dis.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			aa, bb, jj, _ = getABJPquad(data, i, j) // quadratic terms
			// 1-2*J/(A*B)
			d := 1.0 - 2.0*jj/(aa*bb)
			dis.Set(i, j, d)
			dis.Set(j, i, d)
		}
	}
	return dis
}

func DiceBool_S(data *DenseMatrix) *DenseMatrix {
	dis := DiceBool_D(data)
	//1-D
	return sFromD(dis, 0)
}
