// Taxonomic distance
// Dij = [ Σ ( Xki – Xkj )^2 / N] ^(1/2)
// Sneath, PHA & RR Sokal 1973 Numerical Taxonomy: the principles and practice of numerical classification. WH Freeman and Co., San Francisco.

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	. "math"
)

// Taxonomic distance matrix
func Taxonomic_D(data *DenseMatrix) *DenseMatrix {
	rows := data.Rows()
	cols := data.Cols()
	out := Zeros(rows, rows) // square distance matrix row vs. row

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sum1 := 0.0
			sum2 := 0.0
			sum3 := 0.0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				sum1 += x * x
				sum2 += y * y
				sum3 += x * y
			}
			v := Sqrt(sum1 + sum2 - 2.0*sum3)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}

	}
	return out
}

// Scaled taxonomic distance matrix
func TaxonomicSc_D(data *DenseMatrix) *DenseMatrix {
	out := Taxonomic_D(data)
	rows := data.Rows()

	// find maximum value
	max := 0.0
	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			x := out.Get(i, j)
			if max < x {
				max = x
			}
		}
	}

	// scale
	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			v := out.Get(i, j) / max
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
