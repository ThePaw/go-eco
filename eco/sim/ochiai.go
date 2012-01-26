// Ochiai distance and similarity
// Ochiai (1957)

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	"math"
)

// Ochiai similarity matrix
// Ochiai (1957)
// Legendre & Legendre (1998): 258, eq. 7.17 (S14 index)
func OchiaiBool_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim     *DenseMatrix
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	sim = Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = getABCD(data, i, j)
			s := a / math.Sqrt((a+b)*(a+c))
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}

// Ochiai distance matrix, for boolean data (according to R: vegan)
func OchiaiBool_D(data *DenseMatrix) *DenseMatrix {
	var (
		aa, bb, jj float64
		dis        *DenseMatrix
	)

	rows := data.Rows()
	dis = Zeros(rows, rows)
	warnIfNotBool(data)

	for i := 0; i < rows; i++ {
		dis.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			aa, bb, jj, _ = getABJPbool(data, i, j)
			// 1-J/sqrt(A*B)
			d := 1.0 - jj/math.Sqrt(aa*bb)
			dis.Set(i, j, d)
			dis.Set(j, i, d)
		}
	}
	return dis
}
