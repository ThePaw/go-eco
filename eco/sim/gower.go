// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Gower distance and similarity
// Gower is like Manhattan, but data are standardized to range 0..1
// for rows and distance is divided by the number of pairs with both non-missing values. 
// 
// dis[jk] = (1/M) sum (abs(x[ij]-x[ik])/(max(x[i])-min(x[i]))
// where M is the number of columns (excluding missing values) 
// 

import (
	"code.google.com/p/go-eco/eco/aux"
	. "math"
)

// Gower_D returns a Gower distance matrix for floating-point data. 
// Gower, J. C. (1971), “A general coefficient of similarity and some of its properties”. Biometrics, 27, 623–637.
// Kaufman, L. and Rousseeuw, P.J. (1990), Finding Groups in Data: An Introduction to Cluster Analysis. Wiley, New York. 
func Gower_D(data *aux.Matrix) *aux.Matrix {
	const missing float64 = -999 //code for missing values

	rows := data.R
	cols := data.C
	out := aux.NewMatrix(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sum := 0.0
			count := 0
			maxx := 0.0
			minx := 0.0

			// columns are considered as interval-scaled variables and 
			// d_ijk = abs(x_ik - x_jk) / R_k
			// where R_k is the range of the kth variable. 

			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				maxx = Max(x, maxx)
				maxx = Max(y, maxx)
				minx = Min(x, minx)
				minx = Min(y, minx)
				if x != missing && y != missing {
					r := maxx - minx
					v := Abs(x-y) / r
					sum += v
					count++
				}
			}
			v := sum / float64(count)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// GowerOrd_D returns a Gower distance matrix for rank ordered variables. 
func GowerOrd_D(data *aux.Matrix, kr bool) *aux.Matrix {
	// If kr == true, the extension of the Gower's dissimilarity measure proposed by Kaufman and Rousseeuw (1990) is used. 
	// Otherwise, the original Gower's (1971) dissimilarity is considered. 
	const missing float64 = -999 //code for missing values

	rows := data.R
	cols := data.C
	out := aux.NewMatrix(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sum := 0.0
			count := 0
			maxx := 0.0
			minx := 0.0

			// columns are considered as categorical ordinal variables and the values are substituted 
			// with the corresponding position index, r_ik in the factor levels. 
			// These position indexes (that are different from the output of the R function rank) are transformed in the following manner: 
			// z_ik = (r_ik - 1)/(max(r_ik) - 1)
			// These new values, z_ik, are treated as observations of an interval scaled variable. 

			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				maxx = Max(x, maxx)
				maxx = Max(y, maxx)
				minx = Min(x, minx)
				minx = Min(y, minx)
				if x != missing && y != missing {
					r := maxx - 1
					if kr {
						x = (x - 1) / r
						y = (y - 1) / r
					}
					v := Abs(x-y) / r
					sum += v
					count++
				}
			}
			v := sum / float64(count)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// GowerBool_D returns a Gower distance matrix for boolean data.
func GowerBool_D(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c, d float64
	)

	aux.WarnIfNotBool(data)

	rows := data.R
	out := aux.NewMatrix(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			a, b, c, d = aux.GetABCD(data, i, j)
			v := (b + c) / (a + b + c + d)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// GowerZBool_D returns a Gower-Z distance matrix for boolean data.
func GowerZBool_D(data *aux.Matrix) *aux.Matrix {
	// Citation needed
	var (
		a, b, c, _ float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	aux.WarnIfNotBool(data)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			a, b, c, _ = aux.GetABCD(data, i, j)
			v := (b + c) / (a + b + c)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// GowerBool_S returns a Gower similarity matrix for boolean data.
func GowerBool_S(data *aux.Matrix) *aux.Matrix {
	// Gower & Legendre (1986)
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = aux.GetABCD(data, i, j)
			v := (a - (b + c) + d) / (a + b + c + d)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
