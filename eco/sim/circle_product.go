// Circle product similarity
// Marquardt, W.H. 1978 Archaeological seriation. In: Schiffer, M.B.(ed.)
// Advances in Archaeological Method and Theory. Academic Press, N.Y., p.281.
// Kendall, D.G. 1971b Seriation from abundance matrices. In: Hodson, F.R.,
// Kendall, D.G. & Tautu, P: Mathematics in the archaeological and historical
// sciences. Edinburgh University Press, pp. 215-252. 

package sim

import (
	. "go-eco.googlecode.com/hg/eco"
	. "math"
)

func circleProduct_S(data *Matrix) *Matrix {
	rows := data.R
	cols := data.C
	out := NewMatrix(rows, rows)

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
