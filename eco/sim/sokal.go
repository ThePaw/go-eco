// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Sokal - Sneath similarity matrices, different types
// Sokal & Sneath (1963)
// Legendre & Legendre 1998: 255, eqs. 7.3 - 7.6

import (
	"code.google.com/p/go-eco/eco/aux"
	"math"
)

// SokalSneath1Bool_S returns a Sokal - Sneath similarity matrix #1 for boolean data. 
// Legendre & Legendre 1998: 255, eq. 7.3 (S3 index). 
func SokalSneath1Bool_S(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = aux.GetABCD(data, i, j)
			v := (2*a + 2*d) / (2*a + b + c + 2*d)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// SokalSneath2Bool_S returns a Sokal - Sneath similarity matrix #2 for boolean data. 
// Legendre & Legendre 1998: 255, eq. 7.4 (S4 index). 
func SokalSneath2Bool_S(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = aux.GetABCD(data, i, j)
			v := (a + d) / (b + c)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// SokalSneath3Bool_S returns a Sokal - Sneath similarity matrix #3 for boolean data. 
// Legendre & Legendre 1998: 255, eq. 7.5 (S5 index). 
func SokalSneath3Bool_S(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = aux.GetABCD(data, i, j)
			v := (a/(a+b) + a/(a+c) + d/(b+d) + d/(c+d)) / 4
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// SokalSneath4Bool_S returns a Sokal - Sneath similarity matrix #4 for boolean data. 
// Legendre & Legendre 1998: 255, eq. 7.6 (S6 index). 
func SokalSneath4Bool_S(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = aux.GetABCD(data, i, j)
			v := a / math.Sqrt((a+b)*(a+c)) * d / math.Sqrt((b+d)*(c+d))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// SokalSneath5Bool_S returns a Sokal - Sneath similarity matrix #5 for boolean data. 
// Sokal & Sneath (1963)  ### REF!pg, eq. 
// Legendre & Legendre 1998: 255, eq. 7.13 (S10 index). 
func SokalSneath5Bool_S(data *aux.Matrix) *aux.Matrix {
	// sokal1 of R:simba
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = aux.GetABCD(data, i, j)
			v := a / (a + 2*(b+c))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// SokalSneath6Bool_S returns a Sokal - Sneath similarity matrix #6 for boolean data. 
func SokalSneath6Bool_S(data *aux.Matrix) *aux.Matrix {
	// sokal2 of R:simba
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = aux.GetABCD(data, i, j)
			v := (a * d) / math.Sqrt((a+b)*(a+c)*(d+b)*(d+c))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// SokalSneath7Bool_S returns a Sokal - Sneath similarity matrix #7 for boolean data. 
func SokalSneath7Bool_S(data *aux.Matrix) *aux.Matrix {
	// sokal3 of R:simba
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = aux.GetABCD(data, i, j)
			v := ((2 * a) + (2 * d)) / (a + d + (a + b + c + d))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// SokalSneath8Bool_S returns a Sokal - Sneath similarity matrix #8 for boolean data. 
func SokalSneath8Bool_S(data *aux.Matrix) *aux.Matrix {
	// sokal4 of R:simba
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = aux.GetABCD(data, i, j)
			v := ((2 * a) + (2 * d)) / (a + d + (a + b + c + d))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// SokalSneath9Bool_S returns a Sokal - Sneath similarity matrix #9 for boolean data. 
// Legendre & Legendre 1998: 255, eq. 7.16 (S13 index). 
func SokalSneath9Bool_S(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = aux.GetABCD(data, i, j)
			v := (a/(a+b) + a/(a+c)) / 2
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
