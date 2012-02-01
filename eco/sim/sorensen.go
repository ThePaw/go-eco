// Sørensen similarity and distance
// Soerensen (1948)

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

// Sørensen similarity matrix, for boolean data
// Legendre & Legendre (1998): 256, eq. 7.11  (S8 index)
func SorensenBool_S(data *DenseMatrix) *DenseMatrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	out := Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = getABCD(data, i, j)
			v := 2 * a / (2*a + b + c)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Sørensen distance matrix, for boolean data
func SorensenBool_D(data *DenseMatrix) *DenseMatrix {
	var (
		aa, bb, jj float64
	)

	rows := data.Rows()
	out := Zeros(rows, rows)
	warnIfNotBool(data)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			aa, bb, jj, _ = getABJPbool(data, i, j)
			// (A+B-2*J)/(A+B)
			v := (aa + bb - 2*jj) / (aa + bb)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

/*
// Sørensen distance matrix, for quantitative data
func Sorensen_D(data *DenseMatrix) *DenseMatrix {
	return Czekanowski_D(data)
}
*/
