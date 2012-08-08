// Abundance - based Coverage Estimator

package rich

import (
	//	"go-fn.googlecode.com/hg/fn"
	. "code.google.com/p/go-eco/eco"
	"math"
)

// Computes the extrapolated species richness of a population using the Abundance - based Coverage Estimator
// Returns a vector of values representing a minimum number of species present in each assemblage if the entire population was censused.
// Colwell K et al.  (2012)
func ACE(data *Matrix) *Vector {
	rows := data.R
	cols := data.C
	out := NewVector(rows)

	WarnIfNotCounts(data)

	for i := 0; i < rows; i++ {
		nr := 0.0
		sr := 0.0
		f1 := 0.0
		sa := 0.0

		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0 {
				if x <= 10 {
					nr += x
					sr++
					if x == 1 {
						f1++
					}
				} else {
					sa++
				}
			}
		}
		ca := 1 - f1/nr
		sumf := 0.0
		for j := 0; j < 10; j++ {
			length := 0.0
			x := data.Get(i, j)
			for k := 0; k < cols; k++ {
				if x == float64(j+1) {
					length++
				}
			}
			sumf += float64(j) * length
		}
		g2a := math.Max((sr/ca)*(sumf/(nr*(nr-1)))-1, 0)
		ace := sa + sr/ca + (f1/ca)*g2a
		out.Set(i, ace)
	}
	return out
}

// Computes the extrapolated species richness of a population using the Incidence - based Coverage Estimator
// Returns a vector of values representing a minimum number of species present in each assemblage if the entire population was censused.
// Colwell K et al.  (2012)
func ICE(data *Matrix) *Vector {
	rows := data.R
	cols := data.C
	out := NewVector(rows)

	WarnIfNotBool(data)

	for i := 0; i < rows; i++ {
		nr := 0.0
		sr := 0.0
		f1 := 0.0
		sa := 0.0

		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0 {
				if x <= 10 {
					nr += x
					sr++
					if x == 1 {
						f1++
					}
				} else {
					sa++
				}
			}
		}
		ca := 1 - f1/nr
		sumf := 0.0
		for j := 0; j < 10; j++ {
			length := 0.0
			x := data.Get(i, j)
			for k := 0; k < cols; k++ {
				if x == float64(j+1) {
					length++
				}
			}
			sumf += float64(j) * length
		}
		g2a := math.Max((sr/ca)*(sumf/(nr*(nr-1)))-1, 0)
		ice := sa + sr/ca + (f1/ca)*g2a
		out.Set(i, ice)
	}
	return out
}
