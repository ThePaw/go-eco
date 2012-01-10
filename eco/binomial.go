// Binomial distance and similarity
// d[jk] = sum(x[ij]*log(x[ij]/n[i]) + x[ik]*log(x[ik]/n[i]) - n[i]*log(1/2))/n[i] 
// where n[i] = x[ij] + x[ik]
// Similarity is 1.00/(d+1), so that it is in [0, 1]
// Binomial index is derived from Binomial deviance under null hypothesis that the two compared communities are equal. It should be able to handle variable sample sizes. The index does not have a fixed upper limit, but can vary among sites with no shared species. For further discussion, see Anderson & Millar (2004). 
// Anderson, M.J. and Millar, R.B. (2004). Spatial variation and effects of habitat on temperate reef fish assemblages in northeastern New Zealand. Journal of Experimental Marine Biology and Ecology 305, 191–221. 

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	"math"
)

// Binomial distance matrix
func Binomial_D(data *DenseMatrix) *DenseMatrix {
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
			sum := 0.0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				n := x + y
				sum += (x*math.Log(x/n) + y*math.Log(y/n) - n*math.Log(0.5))/n
			}
			d:= sum
			dis.Set(i, j, d)
			dis.Set(j, i, d)
		}
	}
	return dis
}

// Binomial similarity matrix
func Binomial_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim, dis *DenseMatrix
	)

	dis = Binomial_D(data)
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

func BinomialBool_D(data *DenseMatrix) *DenseMatrix {
	var (
		dis        *DenseMatrix
		a, b, c, d int64
	)

	rows := data.Rows()
	cols := data.Cols()
	dis = Zeros(rows, rows)
	a = 0
	b = 0
	c = 0
	d = 0

	checkIfBool(data)

	for i := 0; i < rows; i++ {
		dis.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)

				switch {
				case x != 0 && y != 0:
					a++
				case x != 0 && y == 0:
					b++
				case x == 0 && y != 0:
					c++
				case x == 0 && y == 0:
					d++
				}

			}
			d := math.Log(2.0) * float64(b+c)
			dis.Set(i, j, d)
			dis.Set(j, i, d)
		}
	}
	return dis
}


