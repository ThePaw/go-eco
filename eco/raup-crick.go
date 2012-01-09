// Raup - Crick distance and similarity

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	. "gostat.googlecode.com/hg/stat"
	"fmt"
	"os"
)

// Raup - Crick distance matrix for presence-absence data
// Raup–Crick distance is a probabilistic index based on presence/absence data. It is defined as 1 - prob(j), 
// or based on the probability of observing at least j species in shared in compared communities. 
// Legendre & Legendre (1998) suggest using simulations to assess the probability, but the current function uses analytic result 
// from hypergeometric distribution instead. This probability (and the index) is dependent on the number of species missing in both sites, 
// and adding all-zero species to the data or removing missing species from the data will influence the index. 
// The probability (and the index) may be almost zero or almost one for a wide range of parameter values. 
// The index is nonmetric: two communities with no shared species may have a dissimilarity slightly below one, 
// and two identical communities may have dissimilarity slightly above zero. 
// Compared to other metrics for p/a data, Raup-Crick seems to be very robust for small samples.

func RaupCrick_D(data *DenseMatrix) *DenseMatrix {
	var (
		d float64
		aaa, bbb, jjj, count, t1, t2, sim int64
		dis                    *DenseMatrix
	)

	rows := data.Rows()
	cols := data.Cols()
	dis = Zeros(rows, rows)
	checkData(data)

	for i := 0; i < rows; i++ {
		dis.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sim = 0
			t1 = 0
			t2 = 0
			count = 0
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
				count++
			}
			if count == 0 {
				panic("error")
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
			
			/*
			phyper(k, m, size-m,   n)
			       ==
			Hypergeometric_CDF_At(size, m, n, k)
			*/
//fmt.Println("hyper: ", count, aaa, bbb, jjj)
			d = 1.0 - Hypergeometric_CDF_At(count, aaa, bbb, jjj)
			dis.Set(i, j, d)
			dis.Set(j, i, d)
		}
	}
	return dis
}

// Raup - Crick similarity matrix
// If d denotes Raup - Crick distance, similarity is s = 1.00 - d, so that it is in [0, 1]
func RaupCrick_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim, dis *DenseMatrix
	)

	dis = RaupCrick_D(data)
	rows := data.Rows()
	sim = Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		sim.Set(i, i, 1.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			s := 1.0 - dis.Get(i, j)
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}

func checkData(data *DenseMatrix) {
	rows := data.Rows()
	cols := data.Cols()
	warning:=false
L:
	for i := 0; i < cols; i++ {
		colSum:=0
		for j := 0; j < rows; j++ {
			if data.Get(i, j) > 0.0 {
				colSum++
			}
		}
		if colSum == 0 {
			warning=true
			break L
		}
	}
	if warning {
		fmt.Fprint(os.Stderr, "warning: data have empty species which influence the results\n")
	}
	return
}
