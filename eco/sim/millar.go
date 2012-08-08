// Millar distance and similarity

package sim

import (
	. "code.google.com/p/go-eco/eco"
)

// Millar distance matrix == Binomial distance
func Millar_D(data *Matrix) *Matrix {
	return Binomial_D(data)
}
