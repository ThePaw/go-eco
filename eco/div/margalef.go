// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package div

// Margalef's D  diversity index

import (
	"code.google.com/p/go-eco/eco/aux"
	"math"
)

// MargalefDiv returns vector of Margalef diversities. 
func MargalefDiv(data *aux.Matrix) *aux.Vector {
	rows := data.R
	cols := data.C
	out := aux.NewVector(rows)

	for i := 0; i < rows; i++ {
		s := 0.0 // number of species
		n := 0.0 // total number of all individuals in the sample
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			n += x
			if x > 0.0 {
				s++
			}
		}
		v := (s - 1) / math.Log(n)
		out.Set(i, v)
	}
	return out
}
