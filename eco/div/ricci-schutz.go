// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package div

// Ricci-Schutz inequality index (also called Pietra’s measure)

import (
	"code.google.com/p/go-eco/eco/aux"
	"math"
)

// RicciSchutzIneq returns vector of Ricci-Schutz inequalities. 
// F A Cowell: Measurement of Inequality, 2000, in A B Atkinson & F Bourguignon (Eds): Handbook of Income Distribution. Amsterdam.
// F A Cowell: Measuring Inequality, 1995 Prentice Hall/Harvester Wheatshef.
// Marshall & Olkin: Inequalities: Theory of Majorization and Its Applications, New York 1979 (Academic Press).
func RicciSchutzIneq(data *aux.Matrix) *aux.Vector {
	// Algorithm inspired by R:ineq
	rows := data.R
	cols := data.C
	out := aux.NewVector(rows)

	for i := 0; i < rows; i++ {
		s := 0.0 // number of species

		// calculate mean
		meanX := 0.0
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0.0 {
				s++
				meanX += x
			}
		}
		meanX /= s

		// calculate mean difference
		meanD := 0.0
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0.0 {
				meanD += math.Abs(x - meanX)
			}
		}
		meanD /= s
		v := meanD / (2 * meanX)
		out.Set(i, v)
	}
	return out
}
