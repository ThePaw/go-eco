// Robinson's similarity function

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	"math"
)

func robinson_S(data *DenseMatrix) *DenseMatrix {
	var (
		i, j, k      int
		sum, rowsum  float64
		sim, percent *DenseMatrix
	)

	sim = Zeros(data.Rows(), data.Rows()) // square similarity matrix row vs. row
	percent = Zeros(data.Rows(), data.Cols())
	i = 0
	j = 0
	k = 0
	sum = 0.0
	rowsum = 0.0

	// Set diagonal to 200
	for i = 0; i < data.Rows(); i++ {
		sim.Set(i, i, 200.0)
	}

	// calculate percentages
	for i = 0; i < data.Rows(); i++ {

		rowsum = 0
		//for (j = 0; j < data->cols; ++j)
		for j = 0; j < data.Cols(); j++ {

			rowsum += data.Get(i, j)
		}
		for j = 0; j < data.Cols(); j++ {
			//	percent->x[i][j] = data->x[i][j] * 100.0 / rowsum;
			percent.Set(i, j, data.Get(i, j)*100.0/rowsum)

		}
	}

	for j = 0; j < data.Rows(); j++ {
		for k = j + 1; k < data.Rows(); k++ {
			sum = 0
			for i = 0; i < data.Cols(); i++ {
				//	sum += math.Abs(percent->x[j][i] - percent->x[k][i]);
				sum += math.Abs(percent.Get(j, i) - percent.Get(k, i))
			}
			//			sim->x[j][k] = sim->x[k][j] = 200 - sum;

			sim.Set(j, k, 200-sum)
			sim.Set(k, j, 200-sum)
		}
	}
	return sim
}
