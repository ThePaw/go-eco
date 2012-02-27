// Cosine complement distance

package sim

import (
	. "go-eco.googlecode.com/hg/eco"
	"math"
)

// Cosine complement distance matrix, for boolean data
func CosineBool_D(data *Matrix) *Matrix {
	var (
		aa, bb, jj float64
		out        *Matrix
	)

	rows := data.R
	out = NewMatrix(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			aa, bb, jj, _ = GetABJPquad(data, i, j)
			// 1-J/sqrt(A*B)
			v := 1.0 - jj/math.Sqrt(aa*bb)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Cosine distance matrix
// Algorithm taken from: Carbonell, J.G.& al. 1997 Translingual Information
// Retrieval: A comparative evaluation. IJCAI'97. See also Salton, G. 1989
// Automatic text processing: The transformation, Analysis, and retrieval of
// information by computer. Addison-Wesley, Reading, Pennsylvania.
// Jongman, et. al., 1995, page 178)--"More emphasis is given to qualitative
// aspects by not considering a site as point but as a vector.Understandably,
// the direction of this vector tells us something about the relative
// abundances of species. The similarity of two sites can be expressed as some 
// function of the angle between the vector of these sites. Quite common is
// the use of the cosine (or Ochiai coefficient):
// cos=OS=sigma(k)Y(ki)Y(kj)/sqrt{[sigma(k)(Y(ki)^2)][sigma(k)(Y(kj))^2)]}"
// <-- this is obviously for disance between data->cols, not data->rows (++pac). 
func Cosine_D(data *Matrix) *Matrix {
	rows := data.R
	cols := data.C
	out := NewMatrix(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sum1 := 0.0
			sum2 := 0.0
			sum3 := 0.0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				sum1 += x * y
				sum2 += x * x
				sum3 += y * y
			}
			v := sum1 / (math.Sqrt(sum2 * sum3))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
