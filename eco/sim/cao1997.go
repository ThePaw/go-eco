// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// CY distance Index
// Cao et al. 1997a

import (
	"code.google.com/p/go-eco/eco/aux"
	"math"
)

// Cao_D returns a Cao CYd distance matrix for count data. 
// Cao et al. 1997a
func Cao_D(data *aux.Matrix) *aux.Matrix {
	rows := data.R
	cols := data.C
	out := aux.NewMatrix(rows, rows)

	// check whether data are integers; if not, truncate them
	aux.TruncData(data)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {

			sum := 0.0

			n := cols
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				if x == 0 && y == 0 { // total number of species
					n--
				}
				x = math.Max(x, 0.1)
				y = math.Max(y, 0.1)
				sum += ((x+y)*math.Log10((x+y)/2) - x*math.Log10(y) - y*math.Log10(x)/(x+y))
			}

			v := sum / float64(n)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
