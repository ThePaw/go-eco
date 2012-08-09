// Menhinick diversity and equitability

package div

import (
	"code.google.com/p/go-eco/eco/aux"
	"math"
)

// Menhinick diversity index
// Menhinick 1967
func Menhinick_D(data *aux.Matrix) *Vector {
	rows := data.R
	cols := data.C
	out := NewVector(rows)

	for i := 0; i < rows; i++ {
		n := 0.0 // total number of all individuals in the sample
		s := 0.0 // number of species
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0.0 {
				s++
				n += x
			}
		}
		v := s / math.Sqrt(n)
		out.Set(i, v)
	}
	return out
}
