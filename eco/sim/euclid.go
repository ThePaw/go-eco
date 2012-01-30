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
	rows := data.Rows()
	cols := data.Cols()
	out := Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sum := 0.0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				sum += (x - y) * (x - y)
			}
			v := math.Sqrt(sum)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Mean Euclidean distance matrix
func MeanEuclid_D(data *DenseMatrix) *DenseMatrix {
	out := Euclid_D(data)
	rows := out.Rows()
	for i := 0; i < rows; i++ {
		for j := i + 1; j < out.Cols(); j++ {
			v := out.Get(i, j) / float64(rows)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Mean Censored Euclidean distance matrix
func MeanCensoredEuclid_D(data *DenseMatrix) *DenseMatrix {
	var (
		out *DenseMatrix
	)

	rows := data.Rows()
	cols := data.Cols()
	out = Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
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
			v := math.Sqrt(sum / float64(nonzero))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Squared Boolean Euclidean dissimilarity matrix
func EuclidSqBool_D(data *DenseMatrix) *DenseMatrix {
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	out := Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = getABCD(data, i, j)
			v := (b + c) / (a + b + c + d)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Boolean Euclidean dissimilarity matrix
// Mean Euclidean in Ellis et al. (1993)
func EuclidBool_D(data *DenseMatrix) *DenseMatrix {
	out := EuclidSqBool_D(data)
	rows := data.Rows()
	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			v := math.Sqrt(out.Get(i, j))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
