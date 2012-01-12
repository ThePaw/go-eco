// Sørensen distance and similarity

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

// Sørensen distance matrix, for boolean data
func SorensenBool_D(data *DenseMatrix) *DenseMatrix {
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
			// (A+B-2*J)/(A+B)
			d := (aa+bb-2*jj)/(aa+bb)
			dis.Set(i, j, d)
			dis.Set(j, i, d)
		}
	}
	return dis
}

