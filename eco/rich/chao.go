// Abundance - based Coverage Estimator
// Robert K. Colwell, Anne Chao, Nicholas J. Gotelli, Shang-Yi Lin, Chang Xuan Mao, Robin L. Chazdon,  and John T. Longino 2012: Models and estimators linking individual-based and sample-based rarefaction, extrapolation and comparison of assemblages J Plant Ecol (2012) 5(1): 3-21 doi:10.1093/jpe/rtr044.
// These nonparametic estimators of species richness are minimum estimators: their computed values should be viewed as lower bounds of total species numbers, given the information in a sample or sample set.
package rich

import (
	. "go-eco.googlecode.com/hg/eco"
)

// Computes the Chao species estimator for abundance data
// Chao 1984, 1987
func Chao(data *Matrix) *Vector {
	var v float64
	rows := data.R
	cols := data.C
	out := NewVector(rows)

	WarnIfNotCounts(data)

	for i := 0; i < rows; i++ {
		s0 := 0.0
		s1 := 0.0
		s2 := 0.0
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0 {
				s0++
				if x == 1 {
					s1++
				} else if x == 2 {
					s2++
				}
			}

		}
		if (s1-s2)*(s1-s2) == (s1+s2)*(s1+s2) {
			v = s0+s1*(s1-1)/((s2+1)*2)
		} else {
			v = s0+s1*s1/(s2*2)
		}
		out.Set(i, v)
	}
	return out
}


// Computes the Chao species estimator for boolean (presence-absence) data
// Chao 1984, 1987
func ChaoBool(data *Matrix) *Vector {
	var v float64
	rows := data.R
	cols := data.C
	out := NewVector(rows)

	WarnIfNotBool(data)

	for i := 0; i < rows; i++ {
		s0 := 0.0
		s1 := 0.0
		s2 := 0.0
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0 {
				s0++
				if x == 1 {
					s1++
				} else if x == 2 {
					s2++
				}
			}

		}
		if (s1-s2)*(s1-s2) ==(s1+s2)*(s1+s2) {
			v = s0+s1*(s1-1)/((s2+1)*2)
		} else {
			v = s0+s1*s1/(s2*2)
		}
		out.Set(i, v)
	}
	return out
}


