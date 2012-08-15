// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package div

// Coefficient of variation. 

import (
	"code.google.com/p/go-eco/eco/aux"
	"math"
)

// VarCoeffIneq returns vector of Coefficient of variation for population. 
// F A Cowell: Measurement of Inequality, 2000, in A B Atkinson & F Bourguignon (Eds): Handbook of Income Distribution. Amsterdam.
// F A Cowell: Measuring Inequality, 1995 Prentice Hall/Harvester Wheatshef.
// Marshall & Olkin: Inequalities: Theory of Majorization and Its Applications, New York 1979 (Academic Press).
func VarCoeffIneq(data *aux.Matrix) *aux.Vector {
	// Algorithm inspired by R:ineq. 
	rows := data.R
	cols := data.C
	out := aux.NewVector(rows)

	for i := 0; i < rows; i++ {
		// calculate mean and variance
		meanX := 0.0
		s := 0.0 // number of species
		m2 := 0.0
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0.0 {
				s++
				delta := x - meanX
				meanX += delta / s
				if s > 1 {
					m2 += delta * (x - meanX)
				}
			}
		}
		varX := m2 / s
		v := math.Sqrt((s-1)*varX/s) / meanX
		out.Set(i, v)
	}
	return out
}

// VarCoeffSmpIneq returns vector of Coefficient of variation for sample. 
// F A Cowell: Measurement of Inequality, 2000, in A B Atkinson & F Bourguignon (Eds): Handbook of Income Distribution. Amsterdam.
// F A Cowell: Measuring Inequality, 1995 Prentice Hall/Harvester Wheatshef.
// Marshall & Olkin: Inequalities: Theory of Majorization and Its Applications, New York 1979 (Academic Press).
// Algorithm inspired by R:ineq
func VarCoeffSmpIneq(data *aux.Matrix) *aux.Vector {
	rows := data.R
	cols := data.C
	out := aux.NewVector(rows)

	for i := 0; i < rows; i++ {
		// calculate mean and variance
		meanX := 0.0
		s := 0.0 // number of species
		m2 := 0.0
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0.0 {
				s++
				delta := x - meanX
				meanX += delta / s
				if s > 1 {
					m2 += delta * (x - meanX)
				}
			}
		}
		varX := m2 / (s - 1)
		v := math.Sqrt((s-1)*varX/s) / meanX
		out.Set(i, v)
	}
	return out
}

// VarCoeffSqIneq returns vector of Squared coefficient of variation for population. 
func VarCoeffSqIneq(data *aux.Matrix) *aux.Vector {
	rows := data.R
	out := aux.NewVector(rows)
	vc := VarCoeffIneq(data)

	for i := 0; i < rows; i++ {
		v := vc.Get(i)
		v *= v
		out.Set(i, v)
	}
	return out
}

// VarCoeffSqSmpIneq returns vector of Squared coefficient of variation for sample. 
func VarCoeffSqSmpIneq(data *aux.Matrix) *aux.Vector {
	rows := data.R
	out := aux.NewVector(rows)
	vc := VarCoeffSmpIneq(data)

	for i := 0; i < rows; i++ {
		v := vc.Get(i)
		v *= v
		out.Set(i, v)
	}
	return out
}
