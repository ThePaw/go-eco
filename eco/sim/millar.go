// Millar distance and similarity

package sim

import (
	. "go-eco.googlecode.com/hg/eco"
)

// Millar distance matrix == Binomial distance
func Millar_D(data *Matrix) *Matrix {
	return Binomial_D(data)
}
