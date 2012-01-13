// Bray–Curtis distance and similarity
// d[jk] = (sum abs(x[ij]-x[ik])/(sum (x[ij]+x[ik]))
// Similarity is 1.00/(d+1), so that it is in [0, 1]
// Bray JR, Curtis JT (1957) An ordination of the upland forest communities in southern Wisconsin. Ecol. Monogr. 27:325-349.

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	. "math"
)

// Bray–Curtis distance matrix
func BrayCurtis_D(data *DenseMatrix) *DenseMatrix {
	var (
		dis *DenseMatrix
	)
	rows := data.Rows()
	cols := data.Cols()
	dis = Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		dis.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sum1 := 0.0
			sum2 := 0.0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				sum1 += Abs(x - y)
				sum2 += x + y
			}
			d := sum1 / sum2
			dis.Set(i, j, d)
			dis.Set(j, i, d)
		}
	}
	return dis
}

// Bray–Curtis similarity matrix
func BrayCurtis_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim, dis *DenseMatrix
	)

	dis = BrayCurtis_D(data)
	rows := data.Rows()
	sim = Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		sim.Set(i, i, 1.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			//			s := 1.00 / (dis.Get(i, j) + 1.0)
			s := 1.00 - dis.Get(i, j)
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}

func BrayCurtisBool_D(data *DenseMatrix) *DenseMatrix {
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
			d := float64(b+c) / (2.0 * float64(a+b+c)) // ???
			dis.Set(i, j, d)
			dis.Set(j, i, d)
		}
	}
	return dis
}
