// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Czekanowski dissimilarity

import (
	"code.google.com/p/go-eco/eco/aux"
	"math"
)

// Czekanowski1_D returns a Czekanowski dissimilarity matrix #1 (D8 index in Legendre & Legendre, 1998), for floating-point data. 
func Czekanowski1_D(data *aux.Matrix) *aux.Matrix {
	// Czekanowski (1909)
	// Legendre & Legendre (1998): 282, eq. 7.46 (D8 index)
	rows := data.R
	cols := data.C
	out := aux.NewMatrix(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sum := 0.0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				sum += math.Abs(x - y)
			}
			v := sum / float64(cols)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Czekanowski2_D returns a Czekanowski dissimilarity matrix #2 for floating-point data. 
func Czekanowski2_D(data *aux.Matrix) *aux.Matrix {
	// Reference needed !
	rows := data.R
	cols := data.C
	out := aux.NewMatrix(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 1.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sumXX := 0.0
			sumYY := 0.0
			sumXY := 0.0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				sumXX += x * x
				sumYY += y * y
				sumXY += x * y
			}
			v := 1 - 200.0*sumXX/(sumYY+sumXY)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
