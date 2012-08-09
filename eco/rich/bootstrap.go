// Bootstrap estimator of species richness (incidence-based)
// Smith and van Belle 1984
// These nonparametic estimators of species richness are minimum estimators: their computed values should be viewed as lower bounds of total species numbers, given the information in a sample or sample set.
package rich

import (
	"code.google.com/p/go-eco/eco/aux"
)

// Bootstrap estimator of species richness (incidence-based)
// Smith and van Belle 1984
func BootS(data *aux.Matrix) *aux.Vector {
	rows := data.R
	cols := data.C
	p := aux.NewVector(cols)
	out := aux.NewVector(rows)
	m := float64(rows)

	aux.ToBool(data)

	// proportions
	for i := 0; i < rows; i++ {
		sObs := 0.0
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0 {
				sObs++
			}
		}
		p.Set(i, sObs/m)
	}

	// estimator
	for i := 0; i < rows; i++ {
		q1 := 0.0
		q2 := 0.0
		sObs := 0.0
		sum := 0.0
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0 {
				sObs++
				add := 1 - p.Get(j)
				for k := 0; k < rows-1; k++ { // power term
					add *= add
				}
				sum += add
				if x == 1 {
					q1++
				} else if x == 2 {
					q2++
				}
			}
		}
		v := sObs + sum
		out.Set(i, v)
	}
	return out
}
