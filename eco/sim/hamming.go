// Hamming distance and similarity
// Hamming distance between two strings  of equal length is the number of positions at which the corresponding symbols are different. Put another way, it measures the minimum number of substitutions required to change one string into the other, or the number of errors that transformed one string into the other.
// For a fixed length n, the Hamming distance is a metric on the vector space of the words of that length.

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

func hamming_D(data *DenseMatrix) *DenseMatrix {
	rows := data.Rows()
	cols := data.Cols()
	dis := Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		dis.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			count := 0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)

				if x != y {
					count++
				}
			}
			d := float64(count)
			dis.Set(i, j, d)
			dis.Set(j, i, d)
		}
	}
	return dis
}


// Hamming distance matrix, for boolean data
func HammingBool_D(data *DenseMatrix) *DenseMatrix {
	checkIfBool(data)
	return hamming_D(data)
}

// Hamming distance matrix, for categorical data
func HammingCat_D(data *DenseMatrix) *DenseMatrix {
//	checkIfCat(data)
	return hamming_D(data)
}
