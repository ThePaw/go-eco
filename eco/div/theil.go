// Theil inequality index
// F A Cowell: Measurement of Inequality, 2000, in A B Atkinson & F Bourguignon (Eds): Handbook of Income Distribution. Amsterdam.
// F A Cowell: Measuring Inequality, 1995 Prentice Hall/Harvester Wheatshef.
// Marshall & Olkin: Inequalities: Theory of Majorization and Its Applications, New York 1979 (Academic Press).
// Algorithm inspired by R:ineq

package div

import (
	. "go-eco.googlecode.com/hg/eco"
	"math"
)

// Theil inequality index
func Theil_D(data *Matrix,  m int64) *Vector {
	rows := data.R
	cols := data.C
	out := NewVector(rows)

	for i := 0; i < rows; i++ {
		s := 0.0    // number of species
		sumX := 0.0 // total number of all individuals in the sample

		// calculate mean and mean log
		meanX := 0.0
		meanLnX := 0.0
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0.0 {
				s++
				sumX += x
				meanLnX += math.Log(x)
			}
		}
		meanX = sumX / s
		meanLnX /= s

		v := 0.0
		if m == 0 {

			for j := 0; j < cols; j++ {
				x := data.Get(i, j)
				y := x / meanX
				if x > 0.0 {
					y = x * math.Log(y)
					v += y
				}
			}
			v /= sumX
		} else {
			v = math.Exp(meanLnX) / meanX
			v = -math.Log(v)
		}
		out.Set(i, v)
	}
	return out
}
