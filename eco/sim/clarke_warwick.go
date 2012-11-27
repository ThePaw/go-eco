// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Average taxonomic distinctness Δ⁺
// Clarke & Warwick (1998)
// Clarke & Warwick (2001), Eq. 2

import (
	"code.google.com/p/go-eco/eco/aux"
)

// AvTDBool_D returns a Clarke & Warwick Average taxonomic distinctness matrix for boolean data. 
func AvTDBool_D(data *aux.Matrix, weight *aux.Matrix) *aux.Matrix {
	rows := data.R
	cols := data.C
	out := aux.NewMatrix(rows, rows)

	if weight.R != weight.C {
		panic("bad weight matrix")
	}

	if weight.C != data.C {
		panic("data and weight matrices do not correspond")
	}

	s := float64(cols)
	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sum := 0.0
			for k := 0; k < cols; k++ {
				for l := 0; l < cols; l++ {	// allow for asymmetric weights
					if k != l {
					a := data.Get(i, k)
					b := data.Get(j, k)
					c := data.Get(i, l)
					d := data.Get(j, l)
					if a+b+c+d == 4 { // both species present in both samples
						sum += weight.Get(k, l)
					}
					}
				}
			}
			out.Set(i, j, 2*sum/(s*(s-1)))
			out.Set(j, i, 2*sum/(s*(s-1)))
		}
	}
	return out
}
