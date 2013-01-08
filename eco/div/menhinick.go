// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package div

// Menhinick diversity and equitability

import (
	"code.google.com/p/go-eco/eco/aux"
	"math"
)

// MenhinickDiv returns vector of Menhinick diversities. 
// Menhinick 1967
func MenhinickDiv(data *aux.Matrix) *aux.Vector {
	rows := data.R
	cols := data.C
	out := aux.NewVector(rows)

	for i := 0; i < rows; i++ {
		n := 0.0 // total number of all individuals in the sample
		s := 0.0 // number of species
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0.0 {
				s++
				n += x
			}
		}
		v := s / math.Sqrt(n)
		out.Set(i, v)
	}
	return out
}
