// Euclidean distance and similarity
// In N dimensions, the Euclidean distance between two points p and q is √(∑i=1N (pi-qi)²) where pi (or qi) is the coordinate of p (or q) in dimension i.
// Similarity is 1.00/(d+1), so that it is in [0, 1]

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	. "math"
)

// Euclidean distance matrix
func Euclid_D(data *DenseMatrix) *DenseMatrix {
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
				sum += (x - y) * (x - y)
			}
			d := Sqrt(sum)
			dis.Set(i, j, d)
			dis.Set(j, i, d)
		}
	}
	return dis
}

// Euclidean similarity matrix
func Euclid_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim, dis *DenseMatrix
	)

	dis = Euclid_D(data)
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
