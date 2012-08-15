// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package div

// Smith and Wilson's evenness index 1-D. 
// Smith & Wilson, 1996. 

import (
	"code.google.com/p/go-eco/eco/aux"
	. "code.google.com/p/go-eco/eco/rich"
	"math"
)

// Smith and Wilson's evenness index 
// needs to be verified !
// SmithWilson1Eq returns vector of Smith and Wilson's evenness indices 1-D. 
func SmithWilson1Eq(data *aux.Matrix, which byte, small bool) *aux.Vector {
	rows := data.R
	cols := data.C
	dd := SimpsonDiv(data, which, small)
	ss := SObs(data)
	out := aux.NewVector(cols)

	for i := 0; i < rows; i++ {
		s := ss.Get(i)
		d := dd.Get(i)
		v := (1 - d) / (1 - 1/s)
		out.Set(i, v)
	}
	return out
}

// SmithWilson2Eq returns vector of Smith and Wilson's evenness indices B. 
func SmithWilson2Eq(data *aux.Matrix) *aux.Vector {
	rows := data.R
	cols := data.C
	out := aux.NewVector(cols)

	for i := 0; i < rows; i++ {
		s := 0.0 // number of species
		sum1 := 0.0
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0.0 {
				s++
				sum1 += math.Log(x)
			}
		}
		sum1 /= s
		sum2 := 0.0
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0.0 {
				y := math.Log(x) - sum1
				sum2 += y * y
			}
		}
		v := 1.0 - (2.0 / (math.Pi * math.Atan(sum2)))
		out.Set(i, v)
	}
	return out
}
