package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	"os"
	"fmt"
	"math"
)

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

