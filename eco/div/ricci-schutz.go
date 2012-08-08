// Ricci-Schutz inequality index (also called Pietra’s measure)

package div

import (
	. "code.google.com/p/go-eco/eco"
	"math"
)

// Ricci-Schutz inequality index
// F A Cowell: Measurement of Inequality, 2000, in A B Atkinson & F Bourguignon (Eds): Handbook of Income Distribution. Amsterdam.
// F A Cowell: Measuring Inequality, 1995 Prentice Hall/Harvester Wheatshef.
// Marshall & Olkin: Inequalities: Theory of Majorization and Its Applications, New York 1979 (Academic Press).
// Algorithm inspired by R:ineq
func RicciSchutz_D(data *Matrix) *Vector {
	rows := data.R
	cols := data.C
	out := NewVector(rows)

	for i := 0; i < rows; i++ {
		s := 0.0 // number of species

		// calculate mean
		meanX := 0.0
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0.0 {
				s++
				meanX += x
			}
		}
		meanX /= s

		// calculate mean difference
		meanD := 0.0
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0.0 {
				meanD += math.Abs(x - meanX)
			}
		}
		meanD /= s
		v := meanD / (2 * meanX)
		out.Set(i, v)
	}
	return out
}
