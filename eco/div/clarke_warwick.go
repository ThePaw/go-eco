// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package div

// Clarke & Warwick  Δ⁺ and Λ⁺

import (
	"code.google.com/p/go-eco/eco/aux"
)

// VarTD returns Clarke & Warwick Variation in taxonomic distinctnesses Λ⁺. 
// Clarke & Warwick (2001), Eq. 3, 4
func VarTD(data *aux.Matrix, weight *aux.Matrix) *aux.Vector {
	rows := data.R
	cols := data.C
	out := aux.NewVector(rows)
	ω := AvTD(data, weight)
	s := float64(cols)

	for i := 0; i < rows; i++ { // for every sample
		ω2 := ω.Get(i)
		ω2 *= ω2
		sum := 0.0
		for k := 0; k < cols; k++ {
			for l := 0; l < cols; l++ { // allow for asymmetric weights
				if k != l {
					a := data.Get(i, k)
					b := data.Get(i, l)
					if a+b == 2 { // both species present
						w := weight.Get(k, l)
						sum += w * w
					}
				}
			}
		}
		out.Set(i, sum/(s*(s-1))-ω2)
	}
	return out
}

// AvTD returns a Clarke & Warwick Average taxonomic distinctness Δ⁺  vector for boolean data. 
// Clarke & Warwick (2001), Eq. 2
func AvTD(data *aux.Matrix, weight *aux.Matrix) *aux.Vector {
	rows := data.R
	cols := data.C
	out := aux.NewVector(rows)

	if weight.R != weight.C {
		panic("bad weight matrix")
	}

	if weight.C != data.C {
		panic("data and weight matrices do not correspond")
	}

	s := float64(cols)
	for i := 0; i < rows; i++ {
		sum := 0.0
		for k := 0; k < cols; k++ {
			for l := 0; l < cols; l++ { // allow for asymmetric weights
				if k != l {
					a := data.Get(i, k)
					b := data.Get(i, l)
					if a+b == 2 { // both species present
						sum += weight.Get(k, l)
					}
				}
			}
		}
		out.Set(i, 2*sum/(s*(s-1)))
	}
	return out
}
