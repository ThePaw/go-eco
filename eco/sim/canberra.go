// Canberra distance and similarity
// Lance G. N. and Williams W. T. (1967) Mixed data classificatory programs. 1. Agglomerative systems. Aust. Comput. J. 1, 82-85. 

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	. "math"
)

// Canberra distance matrix
func Canberra_D(data *DenseMatrix) *DenseMatrix {
	rows := data.Rows()
	cols := data.Cols()
	out := Zeros(rows, rows) // square distance matrix row vs. row

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sum := 0.0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				sum += Abs((x - y) / (x + y))
			}
			out.Set(i, j, sum)
			out.Set(j, i, sum)
		}
	}
	return out
}

// Scaled Canberra distance matrix
// Reference needed!
func CanberraSc_D(data *DenseMatrix) *DenseMatrix {
	rows := data.Rows()
	cols := data.Cols()
	out := Zeros(rows, rows) // square distance matrix row vs. row

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sum := 0.0
			count := 0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)

				if x != 0 || y != 0 {
					count++
					sum += Abs((x - y) / (x + y))
				}
			}
			v := sum / float64(count)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

func CanberraBool_D(data *DenseMatrix) *DenseMatrix {
	//same as GowerZBool()
	return GowerZBool_D(data)
}
