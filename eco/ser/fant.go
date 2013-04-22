// Copyright 2012 The Gt Authors. All rights reserved. See the LICENSE file.

package ser

// Fast Ant System. Float64 version
// E. D. Taillard 1998. "FANT: Fast ant system.  Technical report IDSIA-46-98, IDSIA, Lugano.

import (
	"math/rand"
)

// (Re-) initialization of the trace. 
func initTraceF64(inc float64, trace Matrix64) {
	n := trace.Rows()
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			trace[i][j] = inc
		}
	}
}

// Trace update. 
func updateTraceF64(p, best_p IntVector, inc *float64, r float64, trace Matrix64) {
	var i int
	n := p.Len()
	for i = 0; i < n && p[i] == best_p[i]; i++ { // skip
	}
	if i == n {
		(*inc)++
		initTraceF64(*inc, trace)
	} else {
		for i = 0; i < n; i++ {
			trace[i][p[i]] += *inc
			trace[i][best_p[i]] += r
		}
	}
}

// Generate a solution with probability of setting p[i] == j 
// proportionnal to trace[i][j]. 
func genTraceF64(p IntVector, trace Matrix64) {
	var target, sum float64
	n := p.Len()
	nexti := NewIntVector(n)
	nextj := NewIntVector(n)
	sum_trace := NewVector64(n)

	nexti.Perm()
	nextj.Perm()

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum_trace[i] += trace[i][j]
		}
	}

	for i := 0; i < n; i++ {
		//		target = unif(0, sum_trace[nexti[i]]-1)
		target = rand.Float64()*sum_trace[nexti[i]] - 1
		j := i
		sum = trace[nexti[i]][nextj[j]]
		for sum < target {
			j++
			sum += trace[nexti[i]][nextj[j]]
		}
		p[nexti[i]] = nextj[j]
		for k := i; k < n; k++ {
			sum_trace[nexti[k]] -= trace[nexti[k]][nextj[j]]
		}
		nextj.Swap(j, i)
	}
}
