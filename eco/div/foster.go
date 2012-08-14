// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package div

// Foster index of poverty. 

import (
	"code.google.com/p/go-eco/eco/aux"
	"math"
)

// FosterIneq returns vector of Foster indices of poverty. 
// Foster, J. E. (1984). On Economic Poverty: A Survey of Aggregate Measures. Advances in Econometrics, 3, 215–251.
// Zheng, B. (1997). Aggregate Poverty Measures. Journal of Economic Surveys, 11, 123–162.
func FosterIneq(data *aux.Matrix, k, parameter float64) *aux.Vector {
	rows := data.R
	cols := data.C
	out := aux.NewVector(rows)

	for i := 0; i < rows; i++ {
		n := 0.0
		v := 0.0

		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0.0 {
				n++
				if x < k {
					v += math.Pow((k-x)/k, parameter-1)
				}
			}
		}
		if n > 0 {
			v /= n
		}
		out.Set(i, v)
	}
	return out
}
