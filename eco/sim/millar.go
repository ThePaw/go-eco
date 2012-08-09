// Millar distance and similarity

package sim

import (
	"code.google.com/p/go-eco/eco/aux"
)

// Millar distance matrix == Binomial distance
func Millar_D(data *aux.Matrix) *aux.Matrix {
	return Binomial_D(data)
}
