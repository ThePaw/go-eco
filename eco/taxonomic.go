// Taxonomic distance and similarity
// Dij = [ Σ ( Xki – Xkj )^2 / N] ^(1/2)
// Sneath, PHA & RR Sokal 1973 Numerical Taxonomy: the principles and practice of numerical classification. WH Freeman and Co., San Francisco.

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	. "math"
)

// Taxonomic distance matrix
func Taxonomic_D(data *DenseMatrix) *DenseMatrix {
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
			sum1 := 0.0; sum2 := 0.0; sum3 := 0.0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				sum1 += x * x
				sum2 += y * y
				sum3 += x * y
			}
			dis.Set(i, j, Sqrt(sum1 + sum2 - 2.0 * sum3))
			dis.Set(j, i, Sqrt(sum1 + sum2 - 2.0 * sum3))
		}

	}
	return dis
}

// Taxonomic similarity matrix
// If d denotes Taxonomic distance, similarity is s=1.00/(d+1), so that it is in [0, 1]
func Taxonomic_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim, dis *DenseMatrix
	)

	dis = Taxonomic_D(data)
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
