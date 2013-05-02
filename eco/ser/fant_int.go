// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package ser

// Fast Ant System. Integer version.
// E. D. Taillard 1998. "FANT: Fast ant system.  Technical report IDSIA-46-98, IDSIA, Lugano.

// (Re-) initialization of the trace. 
func initTrace(inc int, trace IntMatrix) {
	n := trace.Rows()
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			trace[i][j] = inc
		}
	}
}

// Trace update. 
func updateTrace(p, best_p IntVector, inc *int, r int, trace IntMatrix) {
	var i int
	n := p.Len()
	for i = 0; i < n && p[i] == best_p[i]; i++ { // skip
	}
	if i == n {
		(*inc)++
		initTrace(*inc, trace)
	} else {
		for i = 0; i < n; i++ {
			trace[i][p[i]] += *inc
			trace[i][best_p[i]] += r
		}
	}
}

// Generate a solution with probability of setting p[i] == j 
// proportionnal to trace[i][j]. 
func genTrace(p IntVector, trace IntMatrix) {
	var target, sum int
	n := p.Len()
	nexti := NewIntVector(n)
	nextj := NewIntVector(n)
	sum_trace := NewIntVector(n)

	nexti.Perm()
	nextj.Perm()

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum_trace[i] += trace[i][j]
		}
	}

	for i := 0; i < n; i++ {
		target = unif(0, sum_trace[nexti[i]]-1)
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
