// Kolm inequality index

package div

import (
	"code.google.com/p/go-eco/eco/aux"
	"math"
)

// Kolm inequality index TT
// F A Cowell: Measurement of Inequality, 2000, in A B Atkinson & F Bourguignon (Eds): Handbook of Income Distribution. Amsterdam.
// F A Cowell: Measuring Inequality, 1995 Prentice Hall/Harvester Wheatshef.
// Marshall & Olkin: Inequalities: Theory of Majorization and Its Applications, New York 1979 (Academic Press).
// Algorithm inspired by R:ineq
func Kolm_D(data *aux.Matrix, m float64) *Vector {
	rows := data.R
	cols := data.C
	out := NewVector(rows)

	for i := 0; i < rows; i++ {
		s := 0.0 // number of species

		// calculate mean and mean log
		meanX := 0.0
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0.0 {
				s++
				meanX += x
			}
		}
		meanX /= s

		// calculate mean and mean log
		mean2 := 0.0
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0.0 {
				mean2 += math.Exp(m * (meanX - x))
			}
		}
		mean2 /= s

		v := math.Log(mean2) / m
		out.Set(i, v)
	}
	return out
}
