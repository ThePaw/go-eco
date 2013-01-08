// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Similarity matrices from Legendre & Legendre (1998), wrappers for named functions. 

import (
	"code.google.com/p/go-eco/eco/aux"
)

// S1 returns an S1 similarity matrix for boolean data. 
// Legendre & Legendre (1998): 255, eq. 7.1 (S1 index). 
func S1(data *aux.Matrix) *aux.Matrix {
	return SimpleMatchingBool_S(data)
}

// S2 returns an S2 similarity matrix for boolean data. 
// Legendre & Legendre (1998): 255, eq. 7.2 (S2 index). 
func S2(data *aux.Matrix) *aux.Matrix {
	return RogersTanimotoBool_S(data)
}

// S3 returns an S3 similarity matrix for boolean data. 
// Legendre & Legendre (1998): 255, eq. 7.3 (S3 index). 
func S3(data *aux.Matrix) *aux.Matrix {
	return SokalSneath1Bool_S(data)
}

// S4 returns an S4 similarity matrix for boolean data. 
// Legendre & Legendre 1998: 255, eq. 7.4 (S4 index). 
func S4(data *aux.Matrix) *aux.Matrix {
	return SokalSneath2Bool_S(data)
}

// S5 returns an S5 similarity matrix for boolean data. 
// Legendre & Legendre 1998: 255, eq. 7.5 (S5 index). 
func S5(data *aux.Matrix) *aux.Matrix {
	return SokalSneath3Bool_S(data)
}

// S6 returns an S6 similarity matrix for boolean data. 
// Legendre & Legendre 1998: 255, eq. 7.6 (S6 index). 
func S6(data *aux.Matrix) *aux.Matrix {
	return SokalSneath4Bool_S(data)
}

// S7 returns an S7 similarity matrix for boolean data. 
// Legendre & Legendre 1998: 256, eq. 7.10 (S7 index). 
func S7(data *aux.Matrix) *aux.Matrix {
	return JaccardBool_S(data)
}

// S8 returns an S8 similarity matrix for boolean data. 
// Legendre & Legendre (1998): 256, eq. 7.11 (S8 index). 
func S8(data *aux.Matrix) *aux.Matrix {
	return SorensenBool_S(data)
}

// S9 returns an S9 similarity matrix for boolean data. 
// Legendre & Legendre (1998): 257, eq. 7.12 (S9 index). 
func S9(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = aux.GetABCD(data, i, j)
			v := 3 * a / (3*a + b + c)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// S10 returns an S10 similarity matrix for boolean data. 
// Legendre & Legendre 1998: 255, eq. 7.13 (S10 index). 
func S10(data *aux.Matrix) *aux.Matrix {
	return SokalSneath5Bool_S(data)
}

// S11 returns an S11 similarity matrix for boolean data. 
// Legendre & Legendre (1998): 257, eq. 7.14 (S11 index). 
func S11(data *aux.Matrix) *aux.Matrix {
	return RusselRaoBool_S(data)
}

// S12 returns an S12 similarity matrix for boolean data. 
// Legendre & Legendre (1998): 257, eq. 7.15 (S12 index). 
func S12(data *aux.Matrix) *aux.Matrix {
	return Kulczynski1Bool_S(data)
}

// S13 returns an S13 similarity matrix for boolean data. 
// Legendre & Legendre 1998: 255, eq. 7.16 (S13 index). 
func S13(data *aux.Matrix) *aux.Matrix {
	return SokalSneath9Bool_S(data)
}

// S14 returns an S14 similarity matrix for boolean data. 
// Legendre & Legendre (1998): 258, eq. 7.17 (S14 index). 
func S14(data *aux.Matrix) *aux.Matrix {
	return OchiaiBool_S(data)
}

// S15 returns an S15 similarity matrix for boolean data. 
// Legendre & Legendre (1998): 259, eq. 7.20 (S15 index). 
func S15(data *aux.Matrix) *aux.Matrix {
	return GowerBool_S(data)
}

// S16 returns an S16 similarity matrix for boolean data. 
// Legendre & Legendre (1998): 261, eq. 7.21 (S16 index). 
// S16 matrix == Estabrook & Rogers
// TO BE IMPLEMENTED

// S26 returns an S26 similarity matrix for  boolean data. 
// Legendre & Legendre (1998): 258, eq. 7.18 (S26 index). 
func S26(data *aux.Matrix) *aux.Matrix {
	return FaithBool_S(data)
}
