// Pearson??? rho correlations as similarity matrix ???????
package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

func SpearmanRho_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim, ranks *DenseMatrix
	)
	rows := data.Rows()
	cols := data.Cols()
	sim = Zeros(rows, rows)
	ranks = Zeros(rows, cols)

	// ToDo: check for ties

	// calculate ranks
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// count scores lower than this
			count := 0
			for k := 0; k < cols; k++ {
				if data.Get(i, k) <= data.Get(i, j) {
					count++
				}
			}
			ranks.Set(i, j, float64(count))
		}
	}

	for i := 0; i < rows; i++ {
		sim.Set(i, i, 1.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sumd2 := 0.0
			for k := 0; k < cols; k++ {
				sumd2 += (ranks.Get(i, k) - ranks.Get(j, k)) * (ranks.Get(i, k) - ranks.Get(j, k))
			}
			s := 1.0 - 6.0*sumd2/float64(cols*cols*cols-cols)
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}
