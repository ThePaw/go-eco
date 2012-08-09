// Entropy inequality index

package div

import (
	"code.google.com/p/go-eco/eco/aux"
	"math"
)

// Entropy inequality index TT
// F A Cowell: Measurement of Inequality, 2000, in A B Atkinson & F Bourguignon (Eds): Handbook of Income Distribution. Amsterdam.
// F A Cowell: Measuring Inequality, 1995 Prentice Hall/Harvester Wheatshef.
// Marshall & Olkin: Inequalities: Theory of Majorization and Its Applications, New York 1979 (Academic Press).
// Algorithm inspired by R:ineq
func Entropy_D(data *aux.Matrix, p float64) *Vector {
	var out *Vector
	rows := data.R
	cols := data.C

	if p == 0 {
		out = Theil_D(data, 1)
	} else if p == 1 {
		out = Theil_D(data, 0)
	} else {
		out = NewVector(rows)
		for i := 0; i < rows; i++ {
			// calculate mean and count number of species
			meanX := 0.0
			s := 0.0 // number of species
			for j := 0; j < cols; j++ {
				x := data.Get(i, j)
				if x > 0.0 {
					s++
					meanX += x
				}
			}
			meanX /= s

			// calculate mean term
			mean2 := 0.0
			for j := 0; j < cols; j++ {
				x := data.Get(i, j)
				if x > 0.0 {
					mean2 += math.Pow(x/meanX, p) - 1
				}
			}
			mean2 /= s
			v := mean2 / (p * (p - 1))
			out.Set(i, v)
		}
	}
	return out
}
