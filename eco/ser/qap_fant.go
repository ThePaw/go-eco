// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package ser

// Solve the Quadratic Assignment Problem using the Fast ant system. 
// E. D. Taillard 1998. "FANT: Fast ant system.  Technical report IDSIA-46-98, IDSIA, Lugano.

import (
	"fmt"
)

var Verbose bool

// Local search: Scan the neighbourhood at most twice. 
// Perform improvements as soon as they are found. 
func localSearch(a, b IntMatrix, p IntVector, cost *int) {
	// set of moves, numbered from 0 to index
	n := p.Len()
	move := NewIntVector(n * (n - 1) / 2)
	nMov := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			move[nMov] = n*i + j
			nMov++
		}
	}
	improved := true
	for k := 0; k < 2 && improved; k++ {
		improved = false
		for i := 0; i < nMov-1; i++ {
			move.Swap(i, unif(i+1, nMov-1))
		}
		for i := 0; i < nMov; i++ {
			r := move[i] / n
			s := move[i] % n
			d := delta(a, b, p, r, s)
			if d < 0 {
				*cost += d
				p.Swap(r, s)
				improved = true
			}
		}
	}
	return
}

// QAP_SolveFANT solves the Quadratic Assignment Problem using Fast Ant System, single trial. 
func QAP_SolveFANT(a, b IntMatrix, p IntVector, r, m int) int {
	n := p.Len()
	w := p.Clone()
	trace := NewIntMatrix(n, n)
	inc := 1
	initTrace(inc, trace)
	cc := iInf

	// FANT iterations
	lastImpro := 0
	for i := 0; i-lastImpro < m; i++ {
		// Build a new solution
		genTrace(w, trace)
		c := cost(a, b, w)
		// Improve solution with a local search
		localSearch(a, b, w, &c)
		// Best solution improved ?
		if c < cc {
			cc = c
			p.CopyFrom(w)
			if Verbose {
				//				fmt.Printf("iteration %d: cost=%d\n", i, cc)
				//				p.Print()
				fmt.Println(i - lastImpro)
			}
			lastImpro = i
			inc = 1
			initTrace(inc, trace)
		} else {
			// Memory update
			updateTrace(w, p, &inc, r, trace)
		}
	}
	return cc
}

// QAP_fant solves the Quadratic Assignment Problem using Fast Ant System, in k trials. 
func QAP_fant(dist, flow IntMatrix, trials, improLagMax, r int) (int, IntVector) {
	var cost = iInf
	n := dist.Rows()
	p := NewIntVector(n)
	best := p.Clone()
	for i := 0; i < trials; i++ {
		p.Perm()
		cc := QAP_SolveFANT(dist, flow, p, r, improLagMax)
		if cc < cost {
			if i > 0 {
				fmt.Println("improvement in trial ", i)
			}
			cost = cc
			best.CopyFrom(p)
		}
	}

	return cost, best
}
