// Dice's similarity and dissimilarity matrix
// Dice (1945), Wolda (1981)

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	"math"
)

func DiceBool_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim           *DenseMatrix
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	sim = Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = getABCD(data, i, j)
			s:= a / (math.Min(b+a, c+a)) 
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}

// Dice's dissimilarity
// it is not a proper distance metric as it does not possess the property of triangle inequality
// Dice = 2*Jaccard / (1 + Jaccard)
// Formula from R:vegan 
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

