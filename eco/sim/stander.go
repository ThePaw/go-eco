// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Stander similarity matrix

import (
	"code.google.com/p/go-eco/eco/aux"
	"math"
)

// Stander_S returns a Stander similarity matrix for proportions. 
func Stander_S(data *aux.Matrix) *aux.Matrix {
	// Stander J. M. (1970) Diversity and similarity of benthic fauna off the coast of Oregon. 
	// M.S. thesis, Oregon State University, Corvallis, Oregon, USA.

	// recalculate data to proportions
	aux.RecalcToProp(data)

	rows := data.R
	cols := data.C
	out := aux.NewMatrix(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 1.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sumxy := 0.0
			sumxx := 0.0
			sumyy := 0.0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				sumxy += x*y
				sumxx += x*x
				sumyy += y*y
			}
			v := sumxy/math.Sqrt(sumxx*sumyy)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
