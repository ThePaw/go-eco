// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package ser

// Sort the pre-(Anti)-Robinson matrix using the Fast ant system. 
// E. D. Taillard 1998. "FANT: Fast ant system.  Technical report IDSIA-46-98, IDSIA, Lugano.
// Use functions in obj_fn_sim.go for Robinson, obj_fn_dis.go for Anti-Robinson matrix.

import (
	"math"
)

// RobFAnt  sorts the pre-Anti-Robinson matrix using the Fast ant system, single trial. 
func RobFAnt(a Matrix64, p IntVector, objFn ObjFn, isLoss bool, r float64, improLagMax int) float64 {
	var inc, c, cost float64
	n := p.Len()
	w := p.Clone()
	trace := NewMatrix64(n, n)
	inc = 1.0
	initTraceF64(inc, trace)
	if isLoss {
		cost = math.Inf(1)
	} else {
		cost = math.Inf(-1)
	}
	lastImpro := 0
	for i := 0; i-lastImpro < improLagMax; i++ {
		// Build a new solution
		genTraceF64(w, trace)
		c = objFn(a, w)
		// Improve solution with a local search
		robLocalSearch(a, w, &c, objFn, isLoss)

		// Best solution improved ?
		if (isLoss && c < cost) || (!isLoss && c > cost) {
			cost = c
			p.CopyFrom(w)
			lastImpro = i
			inc = 1
			initTraceF64(inc, trace)
		} else { // Memory update
			updateTraceF64(w, p, &inc, r, trace)
		}
	}
	return cost
}

// RobFAntK sorts the pre-(Anti)-Robinson matrix using the Fast Ant System, in k trials. 
func RobFAntK(sim Matrix64, objFn ObjFn, isLoss bool, trials, improLagMax int, r float64) (cost float64, best IntVector) {
	if isLoss {
		cost = math.Inf(1)
	} else {
		cost = math.Inf(-1)
	}
	n := sim.Rows()
	p := NewIntVector(n)
	best = p.Clone()
	for i := 0; i < trials; i++ {
		p.Perm()
		c := RobFAnt(sim, p, objFn, isLoss, r, improLagMax)
		if (isLoss && c < cost) || (!isLoss && c > cost) {
			cost = c
			best.CopyFrom(p)
		}
	}
	return
}
