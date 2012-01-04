package eco

import (
	. "fmt"
	. "gomatrix.googlecode.com/hg/matrix"
)

func ReadMatrix() *DenseMatrix {
	var (
		rows, cols int
		x          float64
		data       *DenseMatrix
	)

	Scanf("%d", &rows)
	Scanf("%d", &cols)

	data = Zeros(rows, cols)
	for i := 0; i < rows; i++ {
		for j := i + 1; j < cols; j++ {
			Scanf("%f", &x)
			data.Set(i, j, x)
		}
	}
	return data
}
