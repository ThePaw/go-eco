// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Hamming distance and similarity
// Hamming distance between two strings  of equal length is the number of positions at which the corresponding symbols are different. Put another way, it measures the minimum number of substitutions required to change one string into the other, or the number of errors that transformed one string into the other.
// For a fixed length n, the Hamming distance is a metric on the vector space of the words of that length.

import (
	"code.google.com/p/go-eco/eco/aux"
)

// hamming_D returns a Hamming distance matrix. 
func hamming_D(data *aux.Matrix) *aux.Matrix {
	rows := data.R
	cols := data.C
	out := aux.NewMatrix(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			count := 0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)

				if x != y {
					count++
				}
			}
			v := float64(count)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// HammingBool_D returns a Hamming distance matrix for boolean data.
func HammingBool_D(data *aux.Matrix) *aux.Matrix {
	aux.WarnIfNotBool(data)
	return hamming_D(data)
}

// HammingCat_D returns a Hamming distance matrix for categorical data.
func HammingCat_D(data *aux.Matrix) *aux.Matrix {
	//	checkIfCat(data)
	return hamming_D(data)
}
