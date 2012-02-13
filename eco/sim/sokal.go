// Sokal - Sneath similarity matrices, different types
// Sokal & Sneath (1963)
// Legendre & Legendre 1998: 255, eqs. 7.3 - 7.6

package sim

import (
	. "go-eco.googlecode.com/hg/eco"
	"math"
)

// Sokal - Sneath similarity matrix #1
// Legendre & Legendre 1998: 255, eq. 7.3 (S3 index)
func SokalSneath1Bool_S(data *Matrix) *Matrix {
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = GetABCD(data, i, j)
			v := (2*a + 2*d) / (2*a + b + c + 2*d)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Sokal - Sneath similarity matrix #2
// Legendre & Legendre 1998: 255, eq. 7.4 (S4 index)
func SokalSneath2Bool_S(data *Matrix) *Matrix {
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = GetABCD(data, i, j)
			v := (a + d) / (b + c)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Sokal - Sneath similarity matrix #3
// Legendre & Legendre 1998: 255, eq. 7.5 (S5 index)
func SokalSneath3Bool_S(data *Matrix) *Matrix {
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = GetABCD(data, i, j)
			v := (a/(a+b) + a/(a+c) + d/(b+d) + d/(c+d)) / 4
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Sokal - Sneath similarity matrix #4
// Legendre & Legendre 1998: 255, eq. 7.6 (S6 index)
func SokalSneath4Bool_S(data *Matrix) *Matrix {
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = GetABCD(data, i, j)
			v := a / math.Sqrt((a+b)*(a+c)) * d / math.Sqrt((b+d)*(c+d))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Sokal - Sneath similarity matrix #5
// Sokal & Sneath (1963)  ### REF!pg, eq.
// Legendre & Legendre 1998: 255, eq. 7.13 (S10 index)
// sokal1 of R:simba
func SokalSneath5Bool_S(data *Matrix) *Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = GetABCD(data, i, j)
			v := a / (a + 2*(b+c))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Sokal - Sneath similarity matrix #6
// sokal2 of R:simba
func SokalSneath6Bool_S(data *Matrix) *Matrix {
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = GetABCD(data, i, j)
			v := (a * d) / math.Sqrt((a+b)*(a+c)*(d+b)*(d+c))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Sokal - Sneath similarity matrix #7
// sokal3 of R:simba
func SokalSneath7Bool_S(data *Matrix) *Matrix {
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = GetABCD(data, i, j)
			v := ((2 * a) + (2 * d)) / (a + d + (a + b + c + d))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Sokal - Sneath similarity matrix #8
// sokal4 of R:simba
func SokalSneath8Bool_S(data *Matrix) *Matrix {
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = GetABCD(data, i, j)
			v := ((2 * a) + (2 * d)) / (a + d + (a + b + c + d))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Sokal - Sneath similarity matrix #9
// Legendre & Legendre 1998: 255, eq. 7.16 (S13 index)
func SokalSneath9Bool_S(data *Matrix) *Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = GetABCD(data, i, j)
			v := (a/(a+b) + a/(a+c)) / 2
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
