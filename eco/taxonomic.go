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
			d := Sqrt(sum1 + sum2 - 2.0*sum3)
			dis.Set(i, j, d)
			dis.Set(j, i, d)
		}

	}
	return dis
}


// Scaled taxonomic distance matrix
func TaxonomicSc_D(data *DenseMatrix) *DenseMatrix {
	var (
		dis *DenseMatrix
	)

	dis = Taxonomic_D(data)
	rows := data.Rows()

	// find maximum value
	max := 0.0
	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			x:= dis.Get(i, j)
			if max < x {
				max = x
			}
		}
	}

	// scale
	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			d:= dis.Get(i, j)/max
			dis.Set(i, j, d)
			dis.Set(j, i, d)
		}
	}
	return dis
}

// Scaled taxonomic  similarity matrix
func Taxonomic_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim, dis *DenseMatrix
	)

	dis = TaxonomicSc_D(data)
	rows := data.Rows()
	sim = Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		sim.Set(i, i, 1.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			s := 1.00 - dis.Get(i, j)
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}


