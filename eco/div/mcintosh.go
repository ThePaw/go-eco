// McIntosh diversity and equitability

package div

import (
	. "code.google.com/p/go-eco/eco"
	"math"
)

// McIntosh D  diversity index
// McIntosh 1967
func McIntosh_D(data *Matrix) *Vector {
	rows := data.R
	cols := data.C
	out := NewVector(rows)

	for i := 0; i < rows; i++ {
		n := 0.0 // total number of all individuals in the sample
		u := 0.0 // sum of squares
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			n += x
			u += x * x
		}
		u = math.Sqrt(u)
		v := (n - u) / (n - math.Sqrt(n))
		out.Set(i, v)
	}
	return out
}

// McIntosh E  equitability index
func McIntosh_E(data *Matrix) *Vector {
	rows := data.R
	cols := data.C
	out := NewVector(rows)

	for i := 0; i < rows; i++ {
		s := 0.0 // number of species
		n := 0.0 // total number of all individuals in the sample
		u := 0.0 // sum of squares
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0.0 {
				s++
				n += x
				u += x * x
			}

		}
		u = math.Sqrt(u)
		v := (n - u) / (n - (n / math.Sqrt(s)))
		out.Set(i, v)
	}
	return out
}
