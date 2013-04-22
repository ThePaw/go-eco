// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package div

// Brillouin diversity. 
// Peet RK 1974 The Measurement of Species Diversity. Annual Review of Ecology and Systematics, 5: 293.

import (
	"code.google.com/p/go-eco/eco/aux"
	"code.google.com/p/go-fn/fn"
	"math"
)

// BrillouinDiv returns a vector of Brillouin diversities. 
// If heterogeneity  is equated with uncertainty, Shannon-Weaver is a biased indicator valid only for an infinite sample. 
// The correct formulation for the finite sample is given by the Brillouin formula. 
// Peet (1974). 
func BrillouinDiv(data *aux.Matrix) *aux.Vector {
	rows := data.R
	cols := data.C
	div := aux.NewVector(rows)

	for i := 0; i < rows; i++ {
		tot := 0.0 // total number of all individuals in the sample
		sumLnF := 0.0
		for j := 0; j < cols; j++ {
			x := math.Floor(data.Get(i, j))
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
