// Euclidean distance and similarity
// In N dimensions, the Euclidean distance between two points p and q is √(∑i=1N (pi-qi)²) where pi (or qi) is the coordinate of p (or q) in dimension i.
// Similarity is 1.00/(d+1), so that it is in [0, 1]

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	"math"
)

// Euclidean distance matrix, float data
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
			d := math.Sqrt(sum)
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
			d := math.Sqrt(sum / float64(nonzero))
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

// Squared Boolean Euclidean similatity matrix
func EuclidSqBool_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim           *DenseMatrix
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	sim = Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = getABCD(data, i, j)
			s:= (b + c) / (a+b+c+d)
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}

// Boolean Euclidean similatity matrix
// Mean Euclidean in Ellis et al. (1993)
func EuclidBool_S(data *DenseMatrix) *DenseMatrix {
	sim := EuclidSqBool_S(data)
	rows := data.Rows()
	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			s := math.Sqrt(sim.Get(i, j))
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}

