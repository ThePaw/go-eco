// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Kendall tau correlations as similarity matrix
// Let (x1, y1), (x2, y2), …, (xn, yn) be a set of joint observations from two random variables X and Y respectively, such that all the values of (xi) and (yi) are unique. Any pair of observations (xi, yi) and (xj, yj) are said to be concordant if the ranks for both elements agree: that is, if both xi > xj and yi > yj or if both xi < xj and yi < yj. They are said to be discordant, if xi > xj and yi < yj or if xi < xj and yi > yj. If xi = xj or yi = yj, the pair is neither concordant nor discordant.
// The Kendall τ coefficient is defined as:
// tau = number of concordant pairs- number of discordant pairs) / (n*(n-1)/2)

import (
	"code.google.com/p/go-eco/eco/aux"
)

// KendallTau_S returns a Kendall tau correlations as similarity matrix for floating-point data. 
func KendallTau_S(data *aux.Matrix) *aux.Matrix {
	rows := data.R
	cols := data.C
	out := aux.NewMatrix(rows, rows)
	ranks := aux.NewMatrix(rows, rows)

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
		out.Set(i, i, 1.0)
	}
	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			con := 0
			dis := 0
			for k := 0; k < cols; k++ {
				for l := k + 1; l < cols; l++ {
					if (ranks.Get(i, k) < ranks.Get(i, l) && ranks.Get(j, k) < ranks.Get(j, l)) || (ranks.Get(i, k) > ranks.Get(i, l)) && (ranks.Get(j, k) > ranks.Get(j, l)) {
						con++
					} else if (ranks.Get(i, k) < ranks.Get(i, l) && ranks.Get(j, k) > ranks.Get(j, l)) || (ranks.Get(i, k) > ranks.Get(i, l) && ranks.Get(j, k) < ranks.Get(j, l)) {
						dis++
					}
				}
			}
			v := float64(con-dis) / (float64(cols*(cols-1)) / 2.0)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
