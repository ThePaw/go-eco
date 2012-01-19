// Czekanowski dissimilarity and similarity matrix
package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

// Czekanowski dissimilarity
func Czekanowski_D(data *DenseMatrix) *DenseMatrix {
	rows := data.Rows()
	cols := data.Cols()
	dis := Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		dis.Set(i, i, 1.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sum1 := 0.0
			sum2 := 0.0
			sum3 := 0.0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				sum1 += x * x
				sum2 += y * y
				sum3 += x * y
			}
			d := 1 - 200.0*sum1/(sum2+sum3)
			dis.Set(i, j, d)
			dis.Set(j, i, d)
		}
	}
	return dis
}

// Czekanowski similarity
func Czekanowski_S(data *DenseMatrix) *DenseMatrix {
	dis := Czekanowski_D(data)
	//1-D
	return sFromD(dis, 0)
}
