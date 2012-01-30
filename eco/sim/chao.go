// Chao distance
// Chao's index (Ecol. Lett. 8, 148-159; 2005) tries to take into
// account the number of unseen shared species using Chao's method for
// estimating the number of unseen species.
// Chao, A., Chazdon, R. L., Colwell, R. K. and Shen, T. (2005). A new statistical approach for assessing similarity of species composition with incidence and abundance data. Ecology Letters 8, 148–159. 
// Similarity is 1.00-v

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	"math"
)

// Chao distance matrix
func Chao_D(data *DenseMatrix) *DenseMatrix {
	var v   float64

	rows := data.Rows()
	cols := data.Cols()
	out := Zeros(rows, rows)
	// check whether data are integers; if not, truncate them
	truncData(data)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {

			itot := 0.0
			jtot := 0.0
			ionce := 0.0
			jonce := 0.0
			itwice := 0.0
			jtwice := 0.0
			ishare := 0.0
			jshare := 0.0
			ishar1 := 0.0
			jshar1 := 0.0

			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)

				itot += x
				jtot += y
				if x > 0 && y > 0 {
					ishare += x
					jshare += y
					if math.Abs(y-1) < 0.01 {
						ishar1 += x
						jonce += 1
					} else if math.Abs(y-2) < 0.01 {
						jtwice += 1
					}
					if math.Abs(x-1) < 0.01 {
						jshar1 += y
						ionce += 1
					} else if math.Abs(x-2) < 0.01 {
						itwice += 1
					}
				}

			}

			uu := ishare / itot
			if ishar1 > 0 {
				if jonce < 1 {
					jonce = 1 // Never true if got here?
				}
				if jtwice < 1 {
					jtwice = 1
				}
				uu += (jtot - 1) / jtot * jonce / jtwice / 2.0 * ishar1 / itot
			}
			if uu > 1 {
				uu = 1
			}
			vv := jshare / jtot
			if jshar1 > 0 {
				if ionce < 1 {
					ionce = 1 // Is this never true?
				}
				if itwice < 1 {
					itwice = 1
				}
				vv += (itot - 1) / itot * ionce / itwice / 2.0 * jshar1 / jtot
			}
			if vv > 1 {
				vv = 1
			}
			if uu <= 0.0 || vv <= 0.0 {
				v = 1.0
			} else {
				v = 1.0 - uu*vv/(uu+vv-uu*vv)
			}
			if v < 0.0 {
				v = 0.0
			}
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}


