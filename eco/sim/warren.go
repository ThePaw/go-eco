// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Warren 
// 

import (
	"code.google.com/p/go-eco/eco/aux"
	"math"
)

// Warren_S returns a Hellinger distance matrix for floating-point data. 
// Warren:2870, Eq. 3 emended in 10.1111/j.1558-5646.2010.01204.x
func Warren_S(data *aux.Matrix) *aux.Matrix {
	// Rao, C.R. (1995) Use of Hellinger distance in graphical displays. 
	// In E.-M. Tiit, T. Kollo, & H. Niemi (Ed.): Multivariate statistics
	// and matrices in statistics. Leiden (Netherland): Brill Academic
	// Publisher. pp. 143–161.
	rows := data.R
	cols := data.C
	out := aux.NewMatrix(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sum := 0.0
			for k := 0; k < cols; k++ {
				x := math.Sqrt(data.Get(i, k))
				y := math.Sqrt(data.Get(j, k))
				sum += (x - y) * (x - y)
			}
			v := 1-0.5*sum
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
