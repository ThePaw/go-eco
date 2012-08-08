// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Horn-Morisita similarity matrix

import (
	. "code.google.com/p/go-eco/eco"
)

// HornMorisita_S returns a Horn-Morisita similarity matrix for floating-point data. 
func HornMorisita_S(data *Matrix) *Matrix {
	//  From R:vegan
	rows := data.R
	cols := data.C
	out := NewMatrix(rows, rows)

	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			sumXY := 0.0
			sumX := 0.0
			sumY := 0.0
			λx := 0.0
			λy := 0.0

			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				sumXY += x * y
				sumX += x
				sumY += y
				λx += x * x
				λy += y * y
			}

			v := 2 * sumXY / (λx/sumX/sumX + λy/sumY/sumY) / sumX / sumY
			//			2*sim/(sq1/t1/t1 + sq2/t2/t2)/t1/t2
			if v < 0 {
				v = 0.0
			}
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Morisita-Horn similarity matrix, from R:fossil
func MorisitaHorn2_S(data *Matrix) *Matrix {
	rows := data.R
	cols := data.C
	out := NewMatrix(rows, rows)

	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			sumXY := 0.0
			sumX := 0.0
			sumY := 0.0
			sumXX := 0.0
			sumYY := 0.0

			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				sumXY += x * y
				sumX += x
				sumY += y
				sumXX += x * x
				sumYY += y * y
			}
			da := sumXX / sumX * sumX
			db := sumYY / sumY * sumY
			v := 2 * sumXY / ((da + db) * sumX * sumY)
			if v < 0 {
				v = 0.0
			}
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
