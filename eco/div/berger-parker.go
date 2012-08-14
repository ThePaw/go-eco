// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package div

// Berger - Parker diversity. 
import (
	"code.google.com/p/go-eco/eco/aux"
	"math"
)

// BergerParkerDiv returns vector of Berger-Parker diversities. 
// The Berger-Parker index equals the maximum p[i] value in the dataset, i.e. the proportional abundance of the most abundant type. 
// This corresponds to the weighted generalized mean of the p[i] values when q approaches infinity, and hence equals the inverse of true diversity of order infinity, 1/∞D. 
func BergerParkerDiv(data *aux.Matrix) *aux.Vector {
	rows := data.R
	cols := data.C
	div := aux.NewVector(rows)

	for i := 0; i < rows; i++ {
		tot := 0.0 // total number of all individuals in the sample
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0 {
				tot += x
			}
		}

		pmax := 0.0
		for j := 0; j < cols; j++ {
			p := data.Get(i, j) / tot
			pmax = math.Max(p, pmax)
		}
		div.Set(i, pmax)
	}
	return div
}
