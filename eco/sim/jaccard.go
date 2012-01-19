// Jaccard similarity matrix
// For vectors x and y the "quadratic" terms are J = sum(x*y), A = sum(x^2), B = sum(y^2), and "minimum" terms are J = sum(pmin(x,y)), A = sum(x) and B = sum(y), and "binary" terms are either of these after transforming data into binary form (shared number of species, and number of species for each row). 
// in 'abcd' notation a = J, b = A-J, c = B-J, d = P-A-B+J. 

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)


func Jaccard_S(data *DenseMatrix, which byte) *DenseMatrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	sim := Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = getABCD(data, i, j)
			s:= a / (a + b + c)
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}

// Jaccard distance
// Jaccard = 1 - (1-Dice) / (1 + Dice)  // test it!!

