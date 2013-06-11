// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package ser

// Local improvement of solution.

// SubMatOpt implements submatrix optimization.
// inspired by Brusco et al. 2008: 509.
func SubMatOpt(dis Matrix64, p IntVector, v int, objFn ObjFn, isLoss bool, optMethod OptMethod3) {
	var cost float64
	if !dis.IsSymmetric() {
		panic("distance matrix not symmetric")
	}
	n := p.Len()
	if dis.Rows() != n {
		panic("dimensions not equal")
	}
	if v >= n {
		v = n - 1
	}

	q := NewMatrix64(v, v)
	ident := NewIntVector(v)
	ident.Order()

	improved := true
	for improved {
		improved = false
		for i := 0; i < n-v; i++ {
			psiSub := ident.Clone()
			d := ident.Clone()

			// Step 1a: copy a segment of permutation vector
			for k := 0; k < v; k++ {
				psiSub[k] = p[i+k]
			}

			// Step 1b: populate the submatrix
			for k := 0; k < v; k++ {
				for l := 0; l < v; l++ {
					q[k][l] = dis[p[i+k]][p[i+l]]
				}
			}

			// Step 1c: optimize(q, psiSub)
			oldCost := objFn(q, ident)
			if !isLoss {
				oldCost = -oldCost
			}

			cost = optMethod(q, d, objFn, isLoss) // cost inverted inside if !isLoss 

			// Step 1d  +  1e 
			if cost < oldCost {
				oldCost = cost
				improved = true

				for k := 0; k < v; k++ {
					p[i+k] = psiSub[d[k]]
				}
				if !p.IsPermutation() {
					psiSub.Print()
					p.Print()
					panic("not a permutation")
				}

			}
		}
	}
}

func Segment4Opt(a Matrix64, p IntVector, objFn ObjFn, isLoss bool) {

	perm4 := IntMatrix{
		{0, 1, 3, 2},
		{0, 2, 1, 3},
		{0, 2, 3, 1},
		{0, 3, 1, 2},
		{0, 3, 2, 1},
		{1, 0, 2, 3},
		{1, 0, 3, 2},
		{1, 2, 0, 3},
		{1, 2, 3, 0},
		{1, 3, 0, 2},
		{1, 3, 2, 0},
		{2, 0, 1, 3},
		{2, 0, 3, 1},
		{2, 1, 0, 3},
		{2, 1, 3, 0},
		{2, 3, 0, 1},
		{2, 3, 1, 0},
		{3, 0, 1, 2},
		{3, 0, 2, 1},
		{3, 1, 0, 2},
		{3, 1, 2, 0},
		{3, 2, 0, 1},
	}

	cost := objFn(a, p)

	if !isLoss {
		cost = -cost
	}
	best := cost
	seg := NewIntVector(4)
	n := p.Len() - 3
	for i := 0; i < n; i++ {
		for j := 0; j < 22; j++ {
			w := p.Clone()
			for k := 0; k < 4; k++ {
				seg[k] = w[i+perm4[j][k]]
			}
			for k := 0; k < 4; k++ {
				w[i+k] = seg[k]
			}
			cost = objFn(a, w)
			if !isLoss {
				cost = -cost
			}
			if cost < best {
				//				fmt.Println("=== IMPROVED ===", cost, best, i, j)
				best = cost
				p.CopyFrom(w)
				j = 0
				if !p.IsPermutation() {
					seg.Print()
					p.Print()
					panic("not a permutation")
				}
			}
		}
	}
	//	cost = objFn(a, p)
	//					p.Print()
	//				fmt.Println("=== Final cost:  ===", cost)
}

func fact(n int) int { // factorial
	if n == 0 {
		return 1
	}
	f := 1
	for i := 0; i < n; i++ {
		f *= f
	}
	return f
}

func SegmentOpt(a Matrix64, p IntVector, v int, objFn ObjFn, isLoss bool) {
	cost := objFn(a, p)
	// Factorials up to 8.
	factorial := []int{
		1,
		1,
		2,
		6,
		24,
		120,
		720,
		5040,
		40320,
	}

	if v > 8 {
		v = 8
	}
	if !isLoss {
		cost = -cost
	}
	best := cost
	seg := NewIntVector(v)
	perm := NewIntVector(v)
	n := p.Len() - v + 1
	for i := 0; i < n; i++ {
		perm.Order()

		for j := 0; j < factorial[v]-1; j++ {
			w := p.Clone()
			perm.NextPermLex()
			for k := 0; k < v; k++ {
				seg[k] = w[i+perm[k]]
			}
			for k := 0; k < v; k++ {
				w[i+k] = seg[k]
			}
			cost = objFn(a, w)
			if !isLoss {
				cost = -cost
			}
			if cost < best {
				//				fmt.Println("=== IMPROVED ===", cost, best, i, j)
				best = cost
				p.CopyFrom(w)
				if !p.IsPermutation() {
					seg.Print()
					p.Print()
					panic("not a permutation")
				}
				j = 0
				perm.Order()
			}
		}
	}
	//	cost = objFn(a, p)
	//					p.Print()
	//				fmt.Println("=== Final cost:  ===", cost)
}
