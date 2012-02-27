// Jackknife estimators of species richness (incidence-based) 
// Burnham and Overton 1978,1979; Heltshe and Forrester 1983
// These nonparametic estimators of species richness are minimum estimators: their computed values should be viewed as lower bounds of total species numbers, given the information in a sample or sample set.
package rich

import (
	. "go-eco.googlecode.com/hg/eco"
)

// First-order jackknife estimator of species richness for boolean (= incidence, presence-absence) data
// Burnham and Overton 1978,1979; Heltshe and Forrester 1983
func Jack1S(data *Matrix) *Vector {
	rows := data.R
	cols := data.C
	out := NewVector(rows)
	m := float64(rows)

	ToBool(data)

	for i := 0; i < rows; i++ {
		q1 := 0.0
		sObs := 0.0
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0 {
				sObs++
				if x == 1 {
					q1++
				}
			}
		}
		v := sObs + q1*(m-1)/m
		out.Set(i, v)
	}
	return out
}

// Second-order jackknife estimator of species richness for boolean (= incidence, presence-absence) data
// Smith and van Belle 1984
func Jack2S(data *Matrix) *Vector {
	rows := data.R
	cols := data.C
	out := NewVector(rows)
	m := float64(rows)

	ToBool(data)

	for i := 0; i < rows; i++ {
		q1 := 0.0
		q2 := 0.0
		sObs := 0.0
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0 {
				sObs++
				if x == 1 {
					q1++
				} else if x == 2 {
					q2++
				}
			}
		}
		a := q1 * (2*m - 3) / m
		b := q2 * (m - 2) * (m - 2) / (m * (m - 1))

		v := sObs + a - b
		out.Set(i, v)
	}
	return out
}
