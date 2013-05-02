// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package ser

// Sort the pre-(Anti-)Robinson matrix using the Fast ant system. Common functions.
// E. D. Taillard 1998. "FANT: Fast ant system.  Technical report IDSIA-46-98, IDSIA, Lugano.

type ObjFn func(sim Matrix64, p IntVector) float64

func robDelta(a Matrix64, p IntVector, r, s int, objFn ObjFn) float64 {
	pNew := p.Clone()
	pNew.Swap(r, s)
	d0 := objFn(a, p)
	d1 := objFn(a, pNew)
	return d1 - d0
}

// Local search: Scan the neighbourhood at most twice. 
// Perform improvements as soon as they are found. 
func robLocalSearch(a Matrix64, p IntVector, cost *float64, objFn ObjFn, isLoss bool) {
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
			d := robDelta(a, p, r, s, objFn)
			if (isLoss && d < 0) || (!isLoss && d > 0) {
				*cost += d
				p.Swap(r, s)
				improved = true
			}
		}
	}
	return
}
