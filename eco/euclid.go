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


// Mean Euclidean distance matrix
func MeanEuclid_D(data *DenseMatrix) *DenseMatrix {
	var (
		dis *DenseMatrix
	)

	dis = Euclid_D(data)
	rows := dis.Rows()
	for i := 0; i < rows; i++ {
		for j := i + 1; j < dis.Cols(); j++ {
			d := dis.Get(i, j) / float64(rows)
			dis.Set(i, j, d)
			dis.Set(j, i, d)
		}
	}
	return dis
}

// Mean Euclidean similarity matrix
// If d denotes Mean Euclidean distance, similarity is s=1.00/(d+1), so that it is in [0, 1]
func MeanEuclid_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim, dis *DenseMatrix
	)

	dis = MeanEuclid_D(data)
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

// Mean Censored Euclidean distance matrix
func MeanCensoredEuclid_D(data *DenseMatrix) *DenseMatrix {
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
			nonzero := 0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				sum += (x - y) * (x - y)
				if x != 0.0 || y != 0.0 {
					nonzero++
				}
			}
			d := Sqrt(sum / float64(nonzero))
			dis.Set(i, j, d)
			dis.Set(j, i, d)
		}
	}
	return dis
}

// Mean Censored Euclidean similarity matrix
// If d denotes Mean Censored Euclidean distance, similarity is s=1.00/(d+1), so that it is in [0, 1]
func MeanCensoredEuclid_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim, dis *DenseMatrix
	)

	dis = MeanCensoredEuclid_D(data)
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
// Squared Boolean Euclidean distance matrix
func EuclidSqBool_D(data *DenseMatrix) *DenseMatrix {
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
			d := float64(b + c)
			dis.Set(i, j, d)
			dis.Set(j, i, d)
		}
	}
	return dis
}

// Boolean Euclidean distance matrix
func EuclidBool_D(data *DenseMatrix) *DenseMatrix {
	var (
		dis        *DenseMatrix
	)
	dis = Euclid_D(data)
	rows := data.Rows()
	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			d := Sqrt(dis.Get(i, j))
			dis.Set(i, j, d)
			dis.Set(j, i, d)
		}
	}
	return dis
}


