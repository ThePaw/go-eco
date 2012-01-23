// Sokal - Sneath similarity matrices, 5 different types
// Sokal & Sneath (1963)
// Legendre & Legendre 1998: 255, eqs. 7.3 - 7.6

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	"math"
)

// Sokal - Sneath similarity matrix #1
// Legendre & Legendre 1998: 255, eq. 7.3 (S3 index)
func SokalSneath1Bool_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim        *DenseMatrix
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	sim = Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = getABCD(data, i, j)
			s := (2*a + 2*d) / (2*a + b + c + 2*d)
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}

// Sokal - Sneath similarity matrix #2
// Legendre & Legendre 1998: 255, eq. 7.4 (S4 index)
func SokalSneath2Bool_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim        *DenseMatrix
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	sim = Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = getABCD(data, i, j)
			s := (a + d) / (b + c)
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}

// Sokal - Sneath similarity matrix #3
// Legendre & Legendre 1998: 255, eq. 7.5 (S5 index)
func SokalSneath3Bool_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim        *DenseMatrix
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	sim = Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = getABCD(data, i, j)
			s := (a/(a+b) + a/(a+c) + d/(b+d) + d/(c+d)) / 4
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}

// Sokal - Sneath similarity matrix #4
// Legendre & Legendre 1998: 255, eq. 7.6 (S6 index)
func SokalSneath4Bool_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim        *DenseMatrix
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	sim = Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = getABCD(data, i, j)
			s := a / math.Sqrt((a+b)*(a+c)) * d / math.Sqrt((b+d)*(c+d))
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}

// Sokal - Sneath similarity matrix #5
// Is this the same as SokalSneath1Bool ??
// Sokal & Sneath (1963)  ### REF!pg, eq.
func SokalSneath5Bool_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim     *DenseMatrix
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	sim = Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = getABCD(data, i, j)
			s := a / (a + 2*(b+c))
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}
