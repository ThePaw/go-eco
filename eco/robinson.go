// Robinson's similarity function

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	. "math"
)

func robinson_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim, percent *DenseMatrix
	)

	rows := data.Rows()
	cols := data.Cols()
	sim = Zeros(data.Rows(), data.Rows()) // square similarity matrix row vs. row
	percent = Zeros(data.Rows(), data.Cols())

	// Set diagonal to 200
	for i := 0; i < data.Rows(); i++ {
		sim.Set(i, i, 200.0)
	}

	// calculate percentages
	for i := 0; i < data.Rows(); i++ {
		rowsum := 0.0
		for j := 0; j < data.Cols(); j++ {

			rowsum += data.Get(i, j)
		}
		for j := 0; j < data.Cols(); j++ {
			percent.Set(i, j, data.Get(i, j)*100.0/rowsum)

		}
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sum := 0.0
			for k := 0; k < cols; k++ {
				x := percent.Get(i, k)
				y := percent.Get(j, k)
				sum += Abs(x - y)
			}
			s := 200.0 - sum
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}
