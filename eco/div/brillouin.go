// Brillouin diversity matrix
package div

import (
	. "go-eco.googlecode.com/hg/eco"
	"go-fn.googlecode.com/hg/fn"
	"math"
)

// Brillouin diversity
// Peet RK 1974 The Measurement of Species Diversity. Annual Review of Ecology and Systematics, 5: 293.
// If heterogeneity  is equated with uncertainty, Shannon-Weaver is a biased indicator validnonly for an infinite sample. 
// The correct formulation for the finite sample is given by the Brillouin formula.
func Brillouin(data *Matrix) *Vector {
	var tot int64
	rows := data.R
	cols := data.C
	div := NewVector(rows)

	for i := 0; i < rows; i++ {
		tot = 0 // total number of all individuals in the sample
		sumLnF := 0.0
		for j := 0; j < cols; j++ {
			x := int64(math.Floor(data.Get(i, j))) // must be int64
			if x > 0 {
				tot += x
			}
			sumLnF += fn.LnFact(x)
		}

		h := math.Exp((fn.LnFact(tot) - sumLnF) - math.Log(float64(tot)))
		div.Set(i, h)
	}
	return div
}
