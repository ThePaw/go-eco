// Kulczynski similarity and distance matrix
// Kulczynski (1928)
// Oosting (1956), Southwood (1978)

package sim

import (
	. "go-eco.googlecode.com/hg/eco"
	. "math"
)

// Kulczynski similarity matrix #1
// Legendre & Legendre (1998): 257, eq. 7.15 (S12 index)
func Kulczynski1Bool_S(data *Matrix) *Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = GetABCD(data, i, j)
			v := a / (b + c)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Kulczynski similarity matrix #2
func Kulczynski2Bool_S(data *Matrix) *Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = GetABCD(data, i, j)
			v := ((a / 2) * ((2 * a) + b + c)) / ((a + b) * (a + c))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Kulczynski distance matrix
func Kulczynski_D(data *Matrix) *Matrix {
	rows := data.R
	cols := data.C
	out := NewMatrix(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sumMin := 0.0
			sumX := 0.0
			sumY := 0.0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				sumMin += Min(x, y)
				sumX += x
				sumY += x
			}
			v := 1 - 0.5*(sumMin/sumX+sumMin/sumY)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Kulczynski similarity matrix
// Legendre & Legendre (1998): 265, eq. 7.25 (S18 index)
// for count or interval data

func Kulczynski_S(data *Matrix) *Matrix {
	rows := data.R
	cols := data.C
	out := NewMatrix(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sumMin := 0.0
			sumX := 0.0
			sumY := 0.0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				sumMin += Min(x, y)
				sumX += x
				sumY += x
			}
			v := 0.5 * (sumMin/sumX + sumMin/sumY)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
