// Robinson's similarity function

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	. "math"
)

func robinson_S(data *DenseMatrix) *DenseMatrix {
	rows := data.Rows()
	cols := data.Cols()
	out := Zeros(data.Rows(), data.Rows()) // square similarity matrix row vs. row
	percent := Zeros(data.Rows(), data.Cols())

	// Set diagonal to 200
	for i := 0; i < data.Rows(); i++ {
		out.Set(i, i, 200.0)
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
			v := 200.0 - sum
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
