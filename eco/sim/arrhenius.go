// Arrhenius dissimilarity: the value of z in the species-area model
// S = c*A^z when combining two sites of equal areas, where S is the
// number of species, A is the area, and c and z are model parameters.
// The A below is not the area (which cancels out), but number of
// species in one of the sites, as defined in designdist().

package sim

import (
	. "go-eco.googlecode.com/hg/eco"
	"math"
)

// Arrhenius distance matrix, for boolean data
func ArrheniusBool_D(data *Matrix) *Matrix {
	var (
		aa, bb, jj float64
		out        *Matrix
	)

	rows := data.R
	out = NewMatrix(rows, rows)
	WarnIfNotBool(data)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			aa, bb, jj, _ = GetABJPbool(data, i, j)
			// (log(A+B-J)-log(A+B)+log(2))/log(2)
			v := (math.Log(aa+bb-jj) - math.Log(aa+bb) + math.Log(2)) / math.Log(2)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
