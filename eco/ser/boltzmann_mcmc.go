// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package ser

// MCMC reordering of similarity matrix rows and columns to sample from the Boltzmann distribution.
// see Miklos (2005)

import (
	"math"
	"math/rand"
)

// proposeRowPerm()
func proposeRowPerm(rowPerm, rowPermNew IntVector) {
	rows := rowPerm.Len()
	d := rand.Intn(rows)
	e := rand.Intn(rows - 1)
	if e >= d {
		e++
	}
	//    swap
	rowPermNew.CopyFrom(rowPerm)
	rowPermNew.Swap(d, e)
}

// BoltzmannMCMC does MCMC sampling from Boltzmann distribution of reordered similarity matrices  (close to Robinson form).
// Inspired by Miklos (2005).
// WARNING: when data is a similarity (correlation, gain) matrix, then columns MUST NOT be permuted separately!! Implemented. Publish!!
func BoltzmannMCMC(sim Matrix64, energyFn string, temp float64, burnIn, totalSamples, iter int) (permBest, rhoH IntVector, h, pOH IntMatrix, enBest float64) {

	rows, _ := sim.Dims()
	if !sim.IsSymmetric() {
		panic("not a symmetric similarity matrix")
	}

	en := Psis

	switch energyFn {
	case "Psi":
		en = Psis
	default:
		panic("bad energyFn")
	}

	// allocate slices
	rowPerm := NewIntVector(rows)
	rowPermNew := NewIntVector(rows)
	permBest = NewIntVector(rows)
	rhoH = NewIntVector(20)        // rho histogram
	pOH = NewIntMatrix(rows, rows) // pair-order histogram
	h = NewIntMatrix(rows, rows)   // ranks histogram

	// initial permutation
	rowPerm.Perm()

	enOld := en(sim, rowPerm)

	// MCMC
	for m := 0; m < totalSamples+burnIn+1; m++ {
		for c := 0; c < iter; c++ {
			proposeRowPerm(rowPerm, rowPermNew)
			enNew := en(sim, rowPermNew)

			// save if best so far
			if enNew < enOld {
				enBest = enNew
				permBest = rowPerm
			}
			// Metropolis ratio (Miklos 2005:3403, Eq. 10)
			deltaE := enNew - enOld
			metropolisRatio := math.Exp(-deltaE / temp)

			//threshold
			eX := rand.Float64()

			if metropolisRatio > eX { // accept
				enOld = enNew
				rowPerm.CopyFrom(rowPermNew)
			}

		}
		if m > burnIn {
			// save the current permutation
			// reverse, if needed
			rho := reverseIfNeeded2(rowPerm)
			// rank correlation
			rr := math.Abs(rho)

			// add pair-orders to pair-order histogram
			addToPairOrderHistogram(rowPerm, pOH)

			// add ranks to histogram
			addToRankHistogram(rowPerm, h)

			// add rho to histogram
			addToRhoHistogram(rr, rhoH)

		}
	}
	return
}
