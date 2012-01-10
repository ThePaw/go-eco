package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	"os"
	"fmt"
	"math"
)

func ReadMatrix() *DenseMatrix {
	var (
		rows, cols int
		x          float64
		data       *DenseMatrix
	)

	fmt.Scanf("%d", &rows)
	fmt.Scanf("%d", &cols)

	data = Zeros(rows, cols)
	for i := 0; i < rows; i++ {
		for j := i + 1; j < cols; j++ {
			fmt.Scanf("%f", &x)
			data.Set(i, j, x)
		}
	}
	return data
}

func truncData(data *DenseMatrix) {
	rows := data.Rows()
	cols := data.Cols()
	eps := 1e-3
	warning:=false

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			x:= data.Get(i, j)
			newX := math.Floor(data.Get(i, j))
			if ! warning {
				if x - newX > eps {
					warning=true
				}
			}
			data.Set(i, j, newX)	// truncate data, anyway
		}
	}
	if warning {
		fmt.Fprint(os.Stderr, "warning: data were not integers, so truncated\n")
	}
	return

}

func checkIfBool(data *DenseMatrix) {
	rows := data.Rows()
	cols := data.Cols()
	eps := 1e-3
	warning:=false

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			x:= data.Get(i, j)
			newX := math.Floor(data.Get(i, j))
			if ! warning {
				if x - newX > eps {
					warning=true
				}
			}
			data.Set(i, j, newX)	// truncate data, anyway
		}
	}
	if warning {
		fmt.Fprint(os.Stderr, "warning: data are 0/1, however, treated as boolean: 0 == false, otherwise true\n")
	}
	return

}


