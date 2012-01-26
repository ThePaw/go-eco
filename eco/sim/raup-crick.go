// Raup - Crick distance and similarity
// Raup & Crick (1979)

package eco

import (
	"math"
	"gostat.googlecode.com/hg/stat"
	"go-fn.googlecode.com/hg/fn"
	"gomatrix.googlecode.com/hg/matrix"
)

// Raup - Crick distance matrix for presence-absence data
// Raup - Crick distance is a probabilistic index based on presence/absence data. It is defined as 1 - prob(j), 
// or based on the probability of observing at least j species in shared in compared communities. 
// Legendre & Legendre (1998) suggest using simulations to assess the probability, but the current function uses analytic result 
// from hypergeometric distribution instead. This probability (and the index) is dependent on the number of species missing in both sites, 
// and adding all-zero species to the data or removing missing species from the data will influence the index. 
// The probability (and the index) may be almost zero or almost one for a wide range of parameter values. 
// The index is nonmetric: two communities with no shared species may have a dissimilarity slightly below one, 
// and two identical communities may have dissimilarity slightly above zero. 
// Compared to other metrics for p/a data, Raup-Crick seems to be very robust for small samples.
// Algorithm from R:vegan
// phyper(k, m, size-m, n) == Hypergeometric_CDF_At(size, m, n, k)

func RaupCrick_D(data *matrix.DenseMatrix) *matrix.DenseMatrix {
	var (
		d                                 float64
		aaa, bbb, jjj, t1, t2, sim int64
		dis                               *matrix.DenseMatrix
	)

	rows := data.Rows()
	cols := data.Cols()
	dis = matrix.Zeros(rows, rows)
	warnIfNotBool(data)
	warnIfDblZeros(data)

	for i := 0; i < rows; i++ {
		dis.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sim = 0
			t1 = 0
			t2 = 0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)

				if x > 0.0 && y > 0.0 {
					sim++
				}
				if x > 0.0 {
					t1++
				}
				if y > 0.0 {
					t2++
				}
			}

			jjj = sim - 1
			if t1 < t2 {
				aaa = t1
				bbb = t2

			} else {
				aaa = t2
				bbb = t1
			}
			//	d = 1 - phyper(jjj, aaa, float64(count) - aaa, bbb, 1, 0);

			//fmt.Println("hyper: ", cols, aaa, bbb, jjj)
			d = 1.0 - stat.Hypergeometric_CDF_At(int64(cols), aaa, bbb, jjj)
			dis.Set(i, j, d)
			dis.Set(j, i, d)
		}
	}
	return dis
}



// Raup - Crick similarity matrix #1
// Raup & Crick (1979): 1217, eq. 4
// This is the naive version of their similarity index;
// for final version, use the algorithm described on page 1219
func RaupCrick1_S(data *matrix.DenseMatrix) *matrix.DenseMatrix {
	var a, b, n int64

	rows := data.Rows()
	cols := data.Cols()
	sim := matrix.Zeros(rows, rows)
	warnIfNotBool(data)
	warnIfDblZeros(data)


	n = int64(cols)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a = 0
			b = 0
			common := 0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				if x > 0.0 && y > 0.0 {
					common++
				}

				if x > 0.0 {
					a++
				}
				if y > 0.0 {
					b++
				}
			}

			p:= 0.0
			for k := 0; k < common; k++ {
				p += probK(a, b, n, int64(k))
			}

			sim.Set(i, j, p)
			sim.Set(j, i, p)
		}
	}
	return sim
}

func probK(a, b, n, k int64) float64{
	logNum:= fn.LnFact(a)+fn.LnFact(b)+fn.LnFact(n-a)+fn.LnFact(n-b)
	logDen:= fn.LnFact(n)+fn.LnFact(k)+fn.LnFact(a-k)+fn.LnFact(k-b)+fn.LnFact(b-k)
	return math.Exp(logNum-logDen)
}

