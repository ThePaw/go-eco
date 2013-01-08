// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package div

// Simpson diversity. 
// Simpson, E H. 1949. Measurement of diversity. Nature 163:688

import "code.google.com/p/go-eco/eco/aux"

// Simpson diversity index. 
// The Simpson index was introduced in 1949 by Edward H. Simpson to measure the degree of concentration when individuals are classified into types.[6] 
// The same index was rediscovered by Orris C. Herfindahl in 1950.[7] The square root of the index had already been introduced in 1945 by the economist Albert O. Hirschman.[8] 
// As a result, the same measure is usually known as the Simpson index in ecology, and as the Herfindahl index or the Herfindahl-Hirschman index (HHI) in economics.
// The measure equals the probability that two entities taken at random from the dataset of interest represent the same type.[6] It equals:
// λ = Σ(p[i]^2)
// This also equals the weighted arithmetic mean of the proportional abundances pi of the types of interest, with the proportional abundances themselves being used as the weights.[1]  
// Proportional abundances are by definition constrained to values between zero and unity, but their weighted arithmetic mean, and hence λ, can never be smaller than 1/S, which is reached when all types are equally abundant.
// By comparing the equation used to calculate λ with the equations used to calculate true diversity, it can be seen that 1/λ equals 2D, i.e. true diversity as calculated with q = 2. 
// The original Simpson's index hence equals the corresponding basic sum.[2]
// The interpretation of λ as the probability that two entities taken at random from the dataset of interest represent the same type assumes that the first entity is replaced to the dataset before taking the second entity. 
// If the dataset is very large, sampling without replacement gives approximately the same result, but in small datasets the difference can be substantial. 
// If the dataset is small, and sampling without replacement is assumed, the probability of obtaining the same type with both random draws is:
// λ = Σ(n[i]*(n[i]-1)) / (N*(N-1))
// where n[i] is the number of entities belonging to the ith type and N is the total number of entities in the dataset.[6]
// Since mean proportional abundance of the types increases with decreasing number of types and increasing abundance of the most abundant type, 
// λ obtains small values in datasets of high diversity and large values in datasets of low diversity. This is counterintuitive behavior for a diversity index, 
// so often such transformations of λ that increase with increasing diversity have been used instead. The most popular of such indices have been the inverse Simpson index (1/λ) and the Gini-Simpson index (1 - λ).[1][2] 
// Both of these have also been called the Simpson index in the ecological literature, so care is needed to avoid accidentally comparing the different indices as if they were the same.

// SimpsonLambdaDiv returns vector of Simpson λ diversities. 
// Simpson 1949. 
func SimpsonLambdaDiv(data *aux.Matrix, small bool) *aux.Vector {
	rows := data.R
	cols := data.C
	out := aux.NewVector(rows)

	for i := 0; i < rows; i++ {
		λ := 0.0
		tot := 0.0
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			tot += x
			if x > 0 {
				if small {
					λ += x * (x - 1)
				} else {
					λ += x * x
				}
			}
		}

		if small {
			λ /= tot * (tot - 1)
		} else {
			λ /= tot * tot
		}
		out.Set(i, λ)
	}
	return out
}

// SimpsonDiv returns vector of Simpson diversities. 
// Simpson 1949. 
func SimpsonDiv(data *aux.Matrix, which byte, small bool) *aux.Vector {
	var d float64
	rows := data.R
	div := SimpsonLambdaDiv(data, small)

	for i := 0; i < rows; i++ {
		λ := div.Get(i)
		switch {
		case which == 'c': // Gini-Simpson, complement
			d = 1 - λ
		case which == 'i': // Inverse Simpson, inverse
			d = 1 / λ
		default: // Inverse Simpson
			d = 1 / λ
		}
		div.Set(i, d)
	}
	return div
}

/*
// SimpsonEq returns vector of Simpson equitabilities. 
func SimpsonEq(data *aux.Matrix, which byte, small bool) *aux.Vector {
	rows := data.R
	div := SimpsonDiv(data, which, small)
	equ := aux.NewVector(rows)
	for i := 0; i < rows; i++ {
		d := div.Get(i)
		e := (1 / d) / S
		equ.Set(i, d)
	}
	return equ
}
*/
