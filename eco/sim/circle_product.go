// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Circle product similarity
// Marquardt, W.H. 1978 Archaeological seriation. In: Schiffer, M.B.(ed.)
// Advances in Archaeological Method and Theory. Academic Press, N.Y., p.281.
// Kendall, D.G. 1971b Seriation from abundance matrices. In: Hodson, F.R.,
// Kendall, D.G. & Tautu, P: Mathematics in the archaeological and historical
// sciences. Edinburgh University Press, pp. 215-252. 

import (
	"code.google.com/p/go-eco/eco/aux"
	. "math"
)

// CircleProduct_S returns a Circle product similarity matrix. 
func CircleProduct_S(data *aux.Matrix) *aux.Matrix {
	rows := data.R
	cols := data.C
	out := aux.NewMatrix(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 1.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sum := 0.0
			for k := 0; k < cols; k++ {
				sum += Min(data.Get(i, k), data.Get(j, k))
			}
			out.Set(i, j, sum)
			out.Set(j, i, sum)
		}
	}
	return out
}
