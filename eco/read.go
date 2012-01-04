package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	"math"
)


func ReadMatrix.....
var(
	data *DenseMatrix
)


	fmt.Scanf("%d", &rows)  ...Fscanln
	fmt.Scanf("%d", &cols)

	data = Zeros(rows, cols)
	for i := 0; i < data.Rows(); i++	{
		for j := i + 1; j < data.Rows(); j++ {
			sum = 0
			fmt.Scanf("%f", &x)
			data.Set(i, j, x)
		}
	}

