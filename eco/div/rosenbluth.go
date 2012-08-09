// Rosenbluth index of concentration

package div

import (
	"code.google.com/p/go-eco/eco/aux"
	"sort"
)

// Rosenbluth index of concentration
// F A Cowell: Measurement of Inequality, 2000, in A B Atkinson & F Bourguignon (Eds): Handbook of Income Distribution. Amsterdam.
// F A Cowell: Measuring Inequality, 1995 Prentice Hall/Harvester Wheatshef.
// M Hall & N Tidemann: Measures of Concentration, 1967, JASA 62, 162-168.
func Rosenbluth_D(data *aux.Matrix) *Vector {
	rows := data.R
	cols := data.C
	out := NewVector(rows)

	for i := 0; i < rows; i++ {
		// unload data row to slice
		arr := make([]float64, cols)
		arr = data.A[i*cols : i*cols+cols]
		sort.Float64s(arr)

		s := 0.0    // number of species
		sumX := 0.0 // total number of all individuals in the sample

		for j := 0; j < cols; j++ {
			x := arr[j]
			if x > 0.0 {
				s++
				sumX += x
			}
		}

		v := 0.0
		for j := 0; j < cols; j++ {
			x := arr[j]
			if x > 0.0 {
				y := x * float64(cols-j)
				y /= sumX
				v += 2 * y
			}
		}
		v = 1 / (v - 1)
		out.Set(i, v)
	}
	return out
}
