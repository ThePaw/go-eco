// Jaccard distance and similarity


// For vectors x and y the "quadratic" terms are J = sum(x*y), A = sum(x^2), B = sum(y^2), and "minimum" terms are J = sum(pmin(x,y)), A = sum(x) and B = sum(y), and "binary" terms are either of these after transforming data into binary form (shared number of species, and number of species for each row). 
// in 'abcd' notation a = J, b = A-J, c = B-J, d = P-A-B+J. 

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

func JaccardBool_D(data *DenseMatrix) *DenseMatrix {
	var (
		aa, bb, jj float64
		dis                    *DenseMatrix
	)

	rows := data.Rows()
	dis = Zeros(rows, rows)
	checkIfBool(data)

	for i := 0; i < rows; i++ {
		dis.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			aa, bb, jj, _ = getABJPbool(data, i, j)
			d := (aa+bb-2*jj)/(aa+bb-jj)
			dis.Set(i, j, d)
			dis.Set(j, i, d)
		}
	}
	return dis
}

