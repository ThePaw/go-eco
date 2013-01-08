// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Binomial distance
// d[jk] = sum(x[ij]*log(x[ij]/n[i]) + x[ik]*log(x[ik]/n[i]) - n[i]*log(1/2))/n[i] 
// where n[i] = x[ij] + x[ik]
// Binomial index is derived from Binomial deviance under null hypothesis that the two compared communities are equal. It should be able to handle variable sample sizes. The index does not have a fixed upper limit, but can vary among sites with no shared species. For further discussion, see Anderson & Millar (2004). 
// Anderson, M.J. and Millar, R.B. (2004). Spatial variation and effects of habitat on temperate reef fish assemblages in northeastern New Zealand. Journal of Experimental Marine Biology and Ecology 305, 191–221. 

import (
	"code.google.com/p/go-eco/eco/aux"
	"math"
)

// Binomial_D returns a Binomial distance matrix for floating-point data. 
func Binomial_D(data *aux.Matrix) *aux.Matrix {
	rows := data.R
	cols := data.C
	out := aux.NewMatrix(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sum := 0.0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				n := x + y
				sum += (x*math.Log(x/n) + y*math.Log(y/n) - n*math.Log(0.5)) / n
			}
			v := sum
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// BinomialBool_D returns a Binomial distance matrix for boolean data. 
func BinomialBool_D(data *aux.Matrix) *aux.Matrix {
	var (
		b, c float64
	)

	aux.WarnIfNotBool(data)

	rows := data.R
	out := aux.NewMatrix(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			_, b, c, _ = aux.GetABCD(data, i, j)
			v := math.Log(2.0) * (b + c)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
