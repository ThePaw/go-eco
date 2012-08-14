// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Canberra distance and similarity
// Lance G. N. and Williams W. T. (1967) Mixed data classificatory programs. 1. Agglomerative systems. Aust. Comput. J. 1, 82-85. 

import (
	"code.google.com/p/go-eco/eco/aux"
	. "math"
)

// Canberra_D returns a Canberra distance matrix for floating-point data. 
func Canberra_D(data *aux.Matrix) *aux.Matrix {
	rows := data.R
	cols := data.C
	out := aux.NewMatrix(rows, rows) // square distance matrix row vs. row

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sum := 0.0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				sum += Abs((x - y) / (x + y))
			}
			out.Set(i, j, sum)
			out.Set(j, i, sum)
		}
	}
	return out
}

// CanberraSc_D returns a Scaled Canberra distance matrix for floating-point data. 
func CanberraSc_D(data *aux.Matrix) *aux.Matrix {
	// Reference needed!
	rows := data.R
	cols := data.C
	out := aux.NewMatrix(rows, rows) // square distance matrix row vs. row

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sum := 0.0
			count := 0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)

				if x != 0 || y != 0 {
					count++
					sum += Abs((x - y) / (x + y))
				}
			}
			v := sum / float64(count)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// CanberraBool_D returns a Canberra distance matrix for boolean data. It equals to GowerZBool.
func CanberraBool_D(data *aux.Matrix) *aux.Matrix {
	//same as GowerZBool()
	return GowerZBool_D(data)
}
