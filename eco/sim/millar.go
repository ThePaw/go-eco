// Millar distance and similarity

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

// Millar distance matrix == Binomial distance
func Millar_D(data *DenseMatrix) *DenseMatrix {
	return Binomial_D(data)
}
