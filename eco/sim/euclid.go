// Euclidean distance and similarity
// In N dimensions, the Euclidean distance between two points p and q is √(∑i=1N (pi-qi)²) where pi (or qi) is the coordinate of p (or q) in dimension i.
// Similarity is 1.00/(d+1), so that it is in [0, 1]

package sim

import (
	. "go-eco.googlecode.com/hg/eco"
	"math"
)

// Euclidean distance matrix, float data
func Euclid_D(data *Matrix) *Matrix {
	rows := data.R
	cols := data.C
	out := NewMatrix(rows, rows)

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
func MeanEuclid_D(data *Matrix) *Matrix {
	out := Euclid_D(data)
	rows := out.R
	for i := 0; i < rows; i++ {
		for j := i + 1; j < out.C; j++ {
			v := out.Get(i, j) / float64(rows)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Mean Censored Euclidean distance matrix
func MeanCensoredEuclid_D(data *Matrix) *Matrix {
	var (
		out *Matrix
	)

	rows := data.R
	cols := data.C
	out = NewMatrix(rows, rows)

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
func EuclidSqBool_D(data *Matrix) *Matrix {
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = GetABCD(data, i, j)
			v := (b + c) / (a + b + c + d)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Boolean Euclidean dissimilarity matrix
// Mean Euclidean in Ellis et al. (1993)
func EuclidBool_D(data *Matrix) *Matrix {
	out := EuclidSqBool_D(data)
	rows := data.R
	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			v := math.Sqrt(out.Get(i, j))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
