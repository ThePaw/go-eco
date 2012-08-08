// Margalef's D  diversity index

package div

import (
	. "code.google.com/p/go-eco/eco"
	"math"
)

// Margalef's D  diversity index
func Margalef_D(data *Matrix) *Vector {
	rows := data.R
	cols := data.C
	out := NewVector(rows)

	for i := 0; i < rows; i++ {
		s := 0.0 // number of species
		n := 0.0 // total number of all individuals in the sample
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			n += x
			if x > 0.0 {
				s++
			}
		}
		v := (s - 1) / math.Log(n)
		out.Set(i, v)
	}
	return out
}
