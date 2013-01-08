// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Drennan distance and similarity

import (
	"code.google.com/p/go-eco/eco/aux"
)

// Drennan_D returns a Drennan distance matrix for floating-point data. 
// Drennan, R.D. 1976 A refinement of chronological seriation using nonmetric 
// multidimensional scaling. American antiquity, 41: 290-302. 
// Marquardt, W.H. 1978 Archaeological seriation. In: Schiffer, M.B.(ed.) 
// Advances in Archaeological Method and Theory. Academic Press, N.Y., p.284. 
func Drennan_D(data *aux.Matrix) *aux.Matrix {
	rows := data.R
	cols := data.C
	percent := aux.NewMatrix(rows, cols) // percentages
	out := aux.NewMatrix(rows, rows)     // distances

	for i := 0; i < rows; i++ {
		rowsum := 0.0
		for j := i + 1; j < cols; j++ {
			rowsum += data.Get(i, j)
		}
		for j := i + 1; j < cols; j++ {
			percent.Set(i, j, data.Get(i, j)*100.0/rowsum)
		}
	}

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sum := 0.0
			for k := 0; k < cols; k++ {
				x := percent.Get(i, k)
				y := percent.Get(j, k)
				sum += (x - y)
			}
			v := sum / 200.0
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
