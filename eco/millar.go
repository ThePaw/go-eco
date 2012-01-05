// Millar distance and similarity

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	. "math"
)

// Millar distance matrix
// "Millar dissimilarity" is unpublished.  Jari Oksanen found this in the lecture
// notes of Marti Anderson over the internet, and she attributes this
// idea to her colleague Russell Millar.  The index is basically
// binomial deviance under H0 that species are equally common in the
// two compared communities.  This could be easily generalized over
// to, say, Poisson case.
func Millar_D(data *DenseMatrix) *DenseMatrix {
	var (
		dis *DenseMatrix
	)

	rows := data.Rows()
	cols := data.Cols()
	dis = Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		dis.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			count := 0
			d := 0.0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				nk := x + y
				if nk == 0 {
					continue
				}
				lognk := Log(nk)
				t1 := 0.0
				t2 := 0.0

				if x > 0 {
					t1 = x * (Log(x) - lognk)
				}
				if y > 0 {
					t2 = y * (Log(y) - lognk)
				}
				d += (t1 + t2 + nk*Log(2)) / nk
				count++
			}

			if count == 0 {
				panic("error")
			}
			if d < 0 {
				d = 0.0
			}
			dis.Set(i, j, d)
			dis.Set(j, i, d)
		}
	}
	return dis
}

// Millar similarity matrix
// If d denotes Millar distance, similarity is s=1.00/(d+1), so that it is in [0, 1]
func Millar_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim, dis *DenseMatrix
	)

	dis = Millar_D(data)
	rows := data.Rows()
	sim = Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		sim.Set(i, i, 1.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			s := 1.00 / (dis.Get(i, j) + 1.0)
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}
