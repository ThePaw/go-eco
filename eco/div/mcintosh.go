// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package div

// McIntosh diversity D and equitability E

import (
	"code.google.com/p/go-eco/eco/aux"
	"math"
)

// McIntoshDiv returns vector of McIntosh diversities. 
// McIntosh 1967. 
func McIntoshDiv(data *aux.Matrix) *aux.Vector {
	rows := data.R
	cols := data.C
	out := aux.NewVector(rows)

	for i := 0; i < rows; i++ {
		n := 0.0 // total number of all individuals in the sample
		u := 0.0 // sum of squares
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			n += x
			u += x * x
		}
		u = math.Sqrt(u)
		v := (n - u) / (n - math.Sqrt(n))
		out.Set(i, v)
	}
	return out
}

// McIntoshEq returns vector of McIntosh E equitabilities. 
// McIntosh 1967. 
func McIntoshEq(data *aux.Matrix) *aux.Vector {
	rows := data.R
	cols := data.C
	out := aux.NewVector(rows)

	for i := 0; i < rows; i++ {
		s := 0.0 // number of species
		n := 0.0 // total number of all individuals in the sample
		u := 0.0 // sum of squares
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0.0 {
				s++
				n += x
				u += x * x
			}

		}
		u = math.Sqrt(u)
		v := (n - u) / (n - (n / math.Sqrt(s)))
		out.Set(i, v)
	}
	return out
}
