// Wilson - Shmida similarity

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

// For boolean data, calculates a, b, c, d values of the contingency table. 
// To be use for calculation of similarity indices. 
func getABCD(data *DenseMatrix, row1, row2 int) (a, b, c, d float64) {
	cols := data.Cols()

	checkIfBool(data)

	a = 0
	b = 0
	c = 0
	d = 0

			for k := 0; k < cols; k++ {
				x := data.Get(row1, k)
				y := data.Get(row2, k)

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
	return a, b, c, d
}

