// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package div

// Herfindahl index of concentration. 

import (
	"code.google.com/p/go-eco/eco/aux"
	"math"
)

// HerfindahlIneq returns vector of Herfindahl inequalities. 
// F A Cowell: Measurement of Inequality, 2000, in A B Atkinson & F Bourguignon (Eds): Handbook of Income Distribution. Amsterdam.
// F A Cowell: Measuring Inequality, 1995 Prentice Hall/Harvester Wheatshef.
// M Hall & N Tidemann: Measures of Concentration, 1967, JASA 62, 162-168.
func HerfindahlIneq(data *aux.Matrix, m float64) *aux.Vector {
	rows := data.R
	cols := data.C
	out := aux.NewVector(rows)

	for i := 0; i < rows; i++ {
		s := 0.0    // number of species
		sumX := 0.0 // total number of all individuals in the sample

		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0.0 {
				s++
				sumX += x
			}
		}

		v := 0.0
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0.0 {
				y := x / sumX
				y = math.Pow(y, m+1)
				y = x * math.Log(y)
				v += y
			}
		}
		v = math.Pow(v, 1/s)
		out.Set(i, v)
	}
	return out
}
