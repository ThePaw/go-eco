// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

// Watts index of poverty. 

package div

import (
	"code.google.com/p/go-eco/eco/aux"
	"math"
)

// WattsIneq returns vector of Watts indices of poverty. 
// Foster, J. E. (1984). On Economic Poverty: A Survey of Aggregate Measures. Advances in Econometrics, 3, 215–251.
// Zheng, B. (1997). Aggregate Poverty Measures. Journal of Economic Surveys, 11, 123–162.
func WattsIneq(data *aux.Matrix, k float64) *aux.Vector {
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
					v += math.Log(k / x)
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
