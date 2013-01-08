// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package div

// Theil inequality index
// The Theil index is a measure of economic inequality. It has also been used to measure the lack of racial diversity. The basic Theil index TT is the same as redundancy in information theory which is the maximum possible entropy of the data minus the observed entropy. It is a special case of the generalized entropy index. It can be viewed as a measure of redundancy, lack of diversity, isolation, segregation, inequality, non-randomness, and compressibility. It was proposed by econometrician Henri Theil, a successor of Jan Tinbergen at the Erasmus University Rotterdam.
// The indices measure an entropic "distance" the population is away from the "ideal" egalitarian state of everyone having the same income. The numerical result is in terms of negative entropy so that a higher number indicates more order that is further away from the "ideal" of maximum disorder. Formulating the index to represent negative entropy instead of entropy allows it to be a measure of inequality rather than equality.

import (
	"code.google.com/p/go-eco/eco/aux"
	"math"
)

// TheilIneq returns vector of Theil inequality indices TT. 
// F A Cowell: Measurement of Inequality, 2000, in A B Atkinson & F Bourguignon (Eds): Handbook of Income Distribution. Amsterdam. 
// F A Cowell: Measuring Inequality, 1995 Prentice Hall/Harvester Wheatshef. 
// Marshall & Olkin: Inequalities: Theory of Majorization and Its Applications, New York 1979 (Academic Press). 
func TheilIneq(data *aux.Matrix, m int64) *aux.Vector {
	// Algorithm inspired by R:ineq
	rows := data.R
	cols := data.C
	out := aux.NewVector(rows)

	for i := 0; i < rows; i++ {
		s := 0.0    // number of species
		sumX := 0.0 // total number of all individuals in the sample

		// calculate mean and mean log
		meanX := 0.0
		meanLnX := 0.0
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0.0 {
				s++
				sumX += x
				meanLnX += math.Log(x)
			}
		}
		meanX = sumX / s
		meanLnX /= s

		v := 0.0
		if m == 0 {
			for j := 0; j < cols; j++ {
				x := data.Get(i, j)
				if x > 0.0 {
					y := x / meanX
					y = x * math.Log(y)
					v += y
				}
			}
			v /= sumX
		} else {
			v = math.Exp(meanLnX) / meanX
			v = -math.Log(v)
		}
		out.Set(i, v)
	}
	return out
}

// Theil L inequality index 
// TheilIneq returns vector of Theil inequality indices TL. 
// TL is also known as the MLD (mean log deviation) because it gives the standard deviation of ln(x)
// http://en.wikipedia.org/wiki/Theil_index
func TheilL_D(data *aux.Matrix, m int64) *aux.Vector {
	rows := data.R
	cols := data.C
	out := aux.NewVector(rows)

	for i := 0; i < rows; i++ {
		s := 0.0 // number of species
		// calculate mean and mean log
		meanX := 0.0
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0.0 {
				s++
				meanX += x
			}
		}
		meanX /= s

		v := 0.0
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0.0 {
				v += meanX / x
			}
		}
		v /= s
		out.Set(i, v)
	}
	return out
}

// TheilSIneq returns vector of Theil S (symmetric) inequality indices TS. 
// Sometimes the average of TT and TL  is used, which has the advantage of being "symmetric" like the Gini, Hoover, and Coulter indices. "Symmetric" means it gives the same result for x as it does for 1/x. 
// http://en.wikipedia.org/wiki/Theil_index. 
func TheilSIneq(data *aux.Matrix, m int64) *aux.Vector {
	rows := data.R
	cols := data.C
	out := aux.NewVector(rows)

	for i := 0; i < rows; i++ {
		s := 0.0    // number of species
		sumX := 0.0 // total number of all individuals in the sample

		// calculate mean and mean log
		meanX := 0.0
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0.0 {
				s++
				sumX += x
			}
		}
		meanX = sumX / s

		v := 0.0
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0.0 {
				v += (x/meanX - 1) * math.Log(x)
			}
		}
		v /= 2 * s
		out.Set(i, v)
	}
	return out
}
