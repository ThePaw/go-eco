// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package div

// Entropy inequality index. 

import (
	"code.google.com/p/go-eco/eco/aux"
	"math"
)

// EntropyIneq returns vector of Entropy inequality indices TT. 
// F A Cowell: Measurement of Inequality, 2000, in A B Atkinson & F Bourguignon (Eds): Handbook of Income Distribution. Amsterdam.
// F A Cowell: Measuring Inequality, 1995 Prentice Hall/Harvester Wheatshef.
// Marshall & Olkin: Inequalities: Theory of Majorization and Its Applications, New York 1979 (Academic Press).
func EntropyIneq(data *aux.Matrix, p float64) *aux.Vector {
	// Algorithm inspired by R:ineq
	var out *aux.Vector
	rows := data.R
	cols := data.C

	if p == 0 {
		out = TheilIneq(data, 1)
	} else if p == 1 {
		out = TheilIneq(data, 0)
	} else {
		out = aux.NewVector(rows)
		for i := 0; i < rows; i++ {
			// calculate mean and count number of species
			meanX := 0.0
			s := 0.0 // number of species
			for j := 0; j < cols; j++ {
				x := data.Get(i, j)
				if x > 0.0 {
					s++
					meanX += x
				}
			}
			meanX /= s

			// calculate mean term
			mean2 := 0.0
			for j := 0; j < cols; j++ {
				x := data.Get(i, j)
				if x > 0.0 {
					mean2 += math.Pow(x/meanX, p) - 1
				}
			}
			mean2 /= s
			v := mean2 / (p * (p - 1))
			out.Set(i, v)
		}
	}
	return out
}
