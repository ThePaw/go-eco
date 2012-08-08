// Atkinson inequality coefficient matrix
package div

import (
	. "code.google.com/p/go-eco/eco"
	"math"
)

func Atkinson(data *Matrix, epsilon float64) *Vector {
	rows := data.R
	cols := data.C
	div := NewVector(rows)

	//check params
	if epsilon <= 0 || epsilon > 1 {
		panic("epsilon")
	}

	for i := 0; i < rows; i++ {
		z := 0.0
		mu := 0.0
		logx := 0.0
		y := 0.0
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x < 0 {
				panic("x < 0")
			}
			if x > 0 {
				z = x
				mu += x
				logx += math.Log(x)
			}
		}

		if z == 0 {
			panic("z = 0")
		}
		mu /= float64(cols)

		if epsilon == 1 {
			// R:			 
			// A <- 1 - (exp(mean(log(x)))/mean(x))
			// mean(log(x))=logx/cols
			y = 1 - (math.Exp(logx/float64(cols)) / mu)
		} else {
			// R:			 
			// x <- (x/mean(x))^(1-parameter)
			meanx := 0.0
			for j := 0; j < cols; j++ {
				x := data.Get(i, j)
				meanx += math.Pow(x/mu, 1-epsilon)
			}
			meanx /= float64(cols)
			// A <- 1 - mean(x)^(1/(1-parameter))
			y = 1 - math.Pow(meanx, (1/(1-epsilon)))
		}
		div.Set(i, y)
	}
	return div
}
