// Copyright 2012 The Gt Authors. All rights reserved. See the LICENSE file.

package ser

// Objective (loss and gain) functions for similarity matrices. 

import (
	"math"
)

// PsiLossSim computes energy ψ(p) of the permuted similarity matrix according to Podani (1994); see  Miklos (2005), Eq. 4.
func PsiLossSim(sim Matrix64, p IntVector) float64 {
	if !sim.IsSymmetric() {
		panic("similarity matrix not symmetric")
	}
	n := p.Len()
	if sim.Rows() != n {
		panic("dimensions not equal")
	}

	loss := 0.0
	rows := p.Len()
	cols := p.Len()
	for i := 0; i < p.Len(); i++ {
		for j := 0; j < p.Len(); j++ {
			x := sim[p[i]][p[j]]
			a := math.Abs(float64(cols*(i+1))/float64(rows) - float64(j+1))
			b := math.Abs(float64(rows*(j+1))/float64(cols) - float64(i+1))
			loss += x*a + b
		}
	}
	return loss
}

// BertinLossSim returns loss of the permuted matrix according to Kostopoulos & Goulermas
func BertinLossSim(dis Matrix64, p IntVector) float64 {
	if !dis.IsSymmetric() {
		panic("distance matrix not symmetric")
	}
	n := p.Len()
	if dis.Rows() != n {
		panic("dimensions not equal")
	}
	sum := 0.0
	for i := 1; i < n; i++ {
		for j := 0; j < n-1; j++ {
			for k := 0; k < i-1; k++ {
				for l := j + 1; l < n; l++ {
					sum += dis[p[k]][p[l]]
				}
			}
			sum *= dis[p[i]][p[j]]
		}
	}
	return sum
}