// Canberra distance and similarity
// Lance G. N. and Williams W. T. (1967) Mixed data classificatory programs. 1. Agglomerative systems. Aust. Comput. J. 1, 82-85. 

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	. "math"
)

// Canberra distance matrix
func Canberra_D(data *DenseMatrix) *DenseMatrix {
	var (
		dis *DenseMatrix
	)
	rows := data.Rows()
	cols := data.Cols()
	dis = Zeros(rows, rows) // square distance matrix row vs. row

	for i := 0; i < rows; i++ {
		dis.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sum := 0.0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				sum += Abs((x - y) / (x + y));
			}
			dis.Set(i, j, sum)
			dis.Set(j, i, sum)
		}
	}
	return dis
}

// Canberra similarity matrix
// If d denotes Canberra distance, similarity is s=1.00/(d+1), so that it is in [0, 1]
func Canberra_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim, dis *DenseMatrix
	)

	dis = Canberra_D(data)
	rows := data.Rows()
	sim = Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		sim.Set(i, i, 1.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			x := dis.Get(i, j) + 1.0
			sim.Set(i, j, 1.00/x)
			sim.Set(j, i, 1.00/x)
		}
	}
	return sim
}

// Scaled Canberra distance matrix
// Reference needed!
func CanberraSc_D(data *DenseMatrix) *DenseMatrix {
	var (
		dis *DenseMatrix
	)
	rows := data.Rows()
	cols := data.Cols()
	dis = Zeros(rows, rows) // square distance matrix row vs. row

	for i := 0; i < rows; i++ {
		dis.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sum := 0.0
			count := 0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)

				if (x != 0 || y != 0) {
					count++;
					sum += Abs((x - y) / (x + y));
				}
			}
			dis.Set(i, j, sum/float64(count))
			dis.Set(j, i, sum/float64(count))
		}
	}
	return dis
}

// Scaled Canberra similarity matrix
// If d denotes Scaled Canberra distance, similarity is s=1.00/(d+1), so that it is in [0, 1]
func CanberraSc_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim, dis *DenseMatrix
	)

	dis = CanberraSc_D(data)
	rows := data.Rows()
	sim = Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		sim.Set(i, i, 1.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			x := dis.Get(i, j) + 1.0
			sim.Set(i, j, 1.00/x)
			sim.Set(j, i, 1.00/x)
		}
	}
	return sim
}
