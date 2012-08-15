// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package rich

import "code.google.com/p/go-eco/eco/aux"

// Simple species richness.
// Richness R simply quantifies how many different types the dataset of interest contains. For example, species richness (usually notated S) of a dataset is the number of different species 
// in the corresponding species list. Richness is a simple measure, so it has been a popular diversity index in ecology, where abundance data are often not available for the datasets of interest. 
// Because richness does not take the abundances of the types into account, it is not the same thing as diversity, which does take abundances into account. 
// However, if true diversity is calculated with q = 0, the effective number of types (0D) equals the actual number of types (R).[2][4]


// SObs returns a vector of numbers of species actually observed in every row of the data matrix.
func SObs(data *aux.Matrix) *aux.Vector {
	rows := data.R
	cols := data.C
	out := aux.NewVector(rows)

	for i := 0; i < rows; i++ {
		r := 0.0
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0 {
				r++
			}
		}
		out.Set(i, r)
	}
	return out
}
