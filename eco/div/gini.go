// Gini inequality index

package div

import (
	. "go-eco.googlecode.com/hg/eco"
	"sort"
)


// Gini inequality index
// F A Cowell: Measurement of Inequality, 2000, in A B Atkinson & F Bourguignon (Eds): Handbook of Income Distribution. Amsterdam.
// F A Cowell: Measuring Inequality, 1995 Prentice Hall/Harvester Wheatshef.
// Marshall & Olkin: Inequalities: Theory of Majorization and Its Applications, New York 1979 (Academic Press).
// In some cases, Gini coefficient can be computed without direct reference to the Lorenz curve. For example, 
// for a population uniform on the values xi, i = 1 to s, indexed in non-decreasing order ( xi ≤ xi+1).
func Gini_D(data *Matrix) *Vector {
	rows := data.R
	cols := data.C
	out := NewVector(rows)

	for i := 0; i < rows; i++ {
		// unload data row to slice
		arr := make([]float64, cols)
		arr = data.A[i*cols : i*cols+cols]
		sort.Float64s(arr)


		// calculate number of species and sums
		s := 0.0    // number of species
		sumX := 0.0
		sumXJ := 0.0
		for j := 0; j < cols; j++ {
			x := arr[j]
			if x > 0.0 {
				s++
				sumX += x
				sumXJ += x*float64(j+1)
			}
		}
		v:= 2*sumXJ/(s*sumX) - (s+1)/s
		out.Set(i, v)
	}
	return out
}


