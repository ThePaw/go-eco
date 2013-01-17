// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package ser

// MCMC reordering of data matrix rows and columns to minimize energy function
// see Miklos (2005)

import (
	"math"
	"math/rand"
)

const (
	T = 1.0
	a = 0.95
	b = 1.0
)

// Boltzmann does MCMC reordering of data matrix rows and columns to minimize energy function
// see Miklos (2005)
func Boltzmann(rows, cols int, matrix [][]int, rowPerm, colPerm []int, energyFn string, iter int, permute bool, seed int64) float64 {
	var (
		rowPermNew, colPermNew   []int
		rowPermBest, colPermBest []int
		i, j, k, s, o, c         int
		eX, enOld, enNew, enBest float64
		eTo, eT                  float64 //ETo = original, typed temperature, the temperature of the variable eT
		permuteCols              bool
	)

	// initialize
	eTo = T
	c = 0
	permuteCols = true
	rand.Seed(seed)

	// select energy function
	en := Psi
	switch energyFn {
	case "Mirror":
		en = Mirror
	case "Psi":
		en = Psi
	case "Psi2":
		en = Psi2
		permuteCols = false
	case "Psi3":
		en = Psi3
		permuteCols = false
	case "Psi4":
		en = Psi4
		permuteCols = false
	}

	// allocate slices
	rowPermNew = make([]int, rows)
	colPermNew = make([]int, cols)
	rowPermBest = make([]int, rows)
	colPermBest = make([]int, cols)

	// initial permutation
	if permute {
		PermVec(rowPerm)
		PermVec(colPerm)
	}

	// normalization of temperature
	eT = 0.0
	for j = 0; j < cols; j++ {
		for i = 0; i < rows-1; i++ {
			for k = i + 1; k < rows; k++ {
				eT += math.Abs(float64(matrix[i][j]-matrix[k][j])) / float64(rows*cols)
			}
		}
	}

	eT = 0.0
	for i = 0; i < rows; i++ {
		for j = 0; j < cols-1; j++ {
			for k = j + 1; k < cols; k++ {
				eT += math.Abs(float64(matrix[i][j]-matrix[i][k])) / float64(rows*cols)
			}
		}
	}

	eT *= eTo
	enOld = en(matrix, rows, cols, rowPerm, colPerm)
	enBest = enOld

	for s = 0; s < rows; s++ {
		rowPermBest[s] = rowPerm[s]
	}
	for o = 0; o < cols; o++ {
		colPermBest[o] = colPerm[o]
	}

	// MCMC
	for c = 0; c < iter; c++ {
		proposeRow(matrix, rowPerm, rowPermNew, rows, cols)
		enNew = en(matrix, rows, cols, rowPermNew, colPerm)
		eX = a + b*rand.Float64()

		if math.Exp(-(enNew-enOld)/eT) > eX { // accept
			enOld = enNew
			if enOld < enBest {
				enBest = enOld
				for s = 0; s < rows; s++ {
					rowPermBest[s] = rowPermNew[s]
				}

			}

			for s = 0; s < rows; s++ {
				rowPerm[s] = rowPermNew[s]
			}
		}

		if permuteCols {
			proposeCol(matrix, colPerm, colPermNew, rows, cols)
			enNew = en(matrix, rows, cols, colPerm, colPermNew)
			eX = a + b*rand.Float64()

			if math.Exp(-(enNew-enOld)/eT) > eX { // accept
				enOld = enNew
				if enOld < enBest {
					enBest = enOld
					for o = 0; o < cols; o++ {
						colPermBest[o] = colPermNew[o]
					}

				}

				for o = 0; o < cols; o++ {
					colPerm[o] = colPermNew[o]
				}
			}
		}
	}

	// copy the best permutations back
	for s = 0; s < rows; s++ {
		rowPerm[s] = rowPermBest[s]
	}
	for o = 0; o < cols; o++ {
		colPerm[o] = colPermBest[o]
	}
	return enBest
}

// Mirror computes energy of the permuted matrix according to Miklos (2005), Eq. 2 (mirror)
func Mirror(matrix [][]int, rows, cols int, rowPerm, colPerm []int) float64 {
	var (
		eE, sum1, sum2 float64
		s, o           int
	)

	sum1 = 0
	sum2 = 0
	for s = 0; s < rows-1; s++ {
		for o = 0; o < cols; o++ {
			sum1 += math.Abs(float64(matrix[rowPerm[s]][o] - matrix[rowPerm[s+1]][o]))
		}
	}

	for o = 0; o < cols; o++ {
		sum1 += math.Abs(float64(matrix[rowPerm[0]][o]-matrix[rowPerm[1]][o])) + math.Abs(float64(matrix[rowPerm[rows-1]][o]-matrix[rowPerm[rows-2]][o]))
	}

	for s = 0; s < rows; s++ {
		for o = 0; o < cols-1; o++ {
			sum2 += math.Abs(float64(matrix[s][colPerm[o]] - matrix[s][colPerm[o+1]]))
		}
	}

	for s = 0; s < rows; s++ {
		sum2 += math.Abs(float64(matrix[s][colPerm[0]]-matrix[s][colPerm[1]])) + math.Abs(float64(matrix[s][colPerm[cols-1]]-matrix[s][colPerm[cols-2]]))
	}

	eE = sum1 + sum2
	return eE
}

// proposeRow()
func proposeRow(matrix [][]int, rowPerm, rowPermNew []int, rows, cols int) float64 {
	var (
		d, e, iI, iJ int
		eDiff        float64
	)

	d = rand.Intn(rows)
	e = rand.Intn(rows - 1)
	if e >= d {
		e++
	} else {
		//    swap
		d, e = e, d
	}

	eDiff = 0.0
	for iJ = 0; iJ < cols; iJ++ {
		if d > 1 {
			eDiff += math.Abs(float64(matrix[rowPerm[e]][iJ]-matrix[rowPerm[d-1]][iJ])) - math.Abs(float64(matrix[rowPerm[d]][iJ]-matrix[rowPerm[d-1]][iJ]))
		} else if d == 0 {
			eDiff += math.Abs(float64(matrix[rowPerm[e]][iJ]-matrix[rowPerm[e-1]][iJ])) - math.Abs(float64(matrix[rowPerm[d]][iJ]-matrix[rowPerm[d+1]][iJ]))
		} else {
			eDiff += 2.0 * (math.Abs(float64(matrix[rowPerm[e]][iJ]-matrix[rowPerm[d-1]][iJ])) - math.Abs(float64(matrix[rowPerm[d]][iJ]-matrix[rowPerm[d-1]][iJ])))
		}

		if e < rows-2 {
			eDiff += math.Abs(float64(matrix[rowPerm[d]][iJ]-matrix[rowPerm[e+1]][iJ])) - math.Abs(float64(matrix[rowPerm[e]][iJ]-matrix[rowPerm[e+1]][iJ]))
		} else if e == rows-1 {
			eDiff += math.Abs(float64(matrix[rowPerm[d]][iJ]-matrix[rowPerm[d+1]][iJ])) - math.Abs(float64(matrix[rowPerm[e]][iJ]-matrix[rowPerm[e-1]][iJ]))
		} else {
			eDiff += 2.0 * (math.Abs(float64(matrix[rowPerm[d]][iJ]-matrix[rowPerm[e+1]][iJ])) - math.Abs(float64(matrix[rowPerm[e]][iJ]-matrix[rowPerm[e+1]][iJ])))
		}
	}

	for iI = 0; iI < d; iI++ {
		rowPermNew[iI] = rowPerm[iI]
	}
	for iI = d; iI <= e; iI++ {
		rowPermNew[iI] = rowPerm[e+d-iI]
	}
	for iI = e + 1; iI < rows; iI++ {
		rowPermNew[iI] = rowPerm[iI]
	}
	return eDiff
}

//  proposeCol()
func proposeCol(matrix [][]int, colPerm, colPermNew []int, rows, cols int) float64 {
	var (
		d, e, iI, iJ int
		eDiff        float64
	)
	d = rand.Intn(cols)
	e = rand.Intn(cols - 1)
	if e >= d {
		e++
	} else {
		//    swap
		d, e = e, d
	}
	eDiff = 0.0
	for iI = 0; iI < rows; iI++ {
		if d > 1 {
			eDiff += math.Abs(float64(matrix[iI][colPerm[e]]-matrix[iI][colPerm[d-1]])) -
				math.Abs(float64(matrix[iI][colPerm[d]]-matrix[iI][colPerm[d-1]]))
		} else if d == 0 {
			eDiff += math.Abs(float64(matrix[iI][colPerm[e]]-matrix[iI][colPerm[e-1]])) -
				math.Abs(float64(matrix[iI][colPerm[d]]-matrix[iI][colPerm[d+1]]))
		} else {
			eDiff += 2.0 * (math.Abs(float64(matrix[iI][colPerm[e]]-matrix[iI][colPerm[d-1]])) -
				math.Abs(float64(matrix[iI][colPerm[d]]-matrix[iI][colPerm[d-1]])))
		}

		if e < cols-2 {
			eDiff += math.Abs(float64(matrix[iI][colPerm[d]]-matrix[iI][colPerm[e+1]])) -
				math.Abs(float64(matrix[iI][colPerm[e]]-matrix[iI][colPerm[e+1]]))
		} else if e == cols-1 {
			eDiff += math.Abs(float64(matrix[iI][colPerm[d]]-matrix[iI][colPerm[d+1]])) -
				math.Abs(float64(matrix[iI][colPerm[e]]-matrix[iI][colPerm[e-1]]))
		} else {
			eDiff += 2.0 * (math.Abs(float64(matrix[iI][colPerm[d]]-matrix[iI][colPerm[e+1]])) -
				math.Abs(float64(matrix[iI][colPerm[e]]-matrix[iI][colPerm[e+1]])))
		}
	}

	for iJ = 0; iJ < d; iJ++ {
		colPermNew[iJ] = colPerm[iJ]
	}
	for iJ = d; iJ <= e; iJ++ {
		colPermNew[iJ] = colPerm[d+e-iJ]
	}
	for iJ = e + 1; iJ < cols; iJ++ {
		colPermNew[iJ] = colPerm[iJ]
	}
	return eDiff
}

// generate a random permutation p
func PermVec(p []int) {
	n := len(p)
	var i int
	for i = 0; i < n; i++ {
		p[i] = i
	}

	for i = 0; i < n; i++ {
		SwapVec(p, i, i+rand.Intn(n-i))
	}
}

func SwapVec(p []int, i, j int) {
	x := p[i]
	p[i] = p[j]
	p[j] = x
}

// Psi computes energy of the permuted matrix according to Podani (1994); see  Miklos (2005), Eq. 4
func Psi(matrix [][]int, rows, cols int, rowPerm, colPerm []int) float64 {
	var (
		eE   float64
		s, o int
	)
	// s =i
	//o = j
	// rows = m
	//cols = n

	eE = 0
	for s = 0; s < rows; s++ {
		for o = 0; o < cols; o++ {
			x := float64(matrix[rowPerm[s]][rowPerm[o]])
			a := math.Abs(float64(cols*(s+1))/float64(rows) - float64(o+1))
			b := math.Abs(float64(rows*(o+1))/float64(cols) - float64(s+1))
			eE += x * (a + b)
		}
	}
	return eE
}

// Psi2 computes energy of the permuted matrix according to Cejchan (unpublished) to seriate Robinson matrix
func Psi2(matrix [][]int, rows, cols int, rowPerm, colPerm []int) float64 {
	var (
		eE, mod      float64
		s, o, posMod int
	)

	eE = 0
	for o = 0; o < cols; o++ {

		// for every species (column) find its modal value
		mod = 0
		for s = 0; s < rows; s++ {
			x := float64(matrix[rowPerm[s]][rowPerm[o]])
			if mod < x {
				mod = x
				posMod = s
			}
		}
		// and use it to calc contribution to energy
		for s = 0; s < rows; s++ {
			x := float64(matrix[rowPerm[s]][rowPerm[o]])
			d := math.Abs(float64(s - posMod))
			eE += x * d
		}
	}
	return eE
}

// Psi3 computes energy of the permuted matrix according to Cejchan (unpublished) to seriate Robinson matrix
func Psi3(matrix [][]int, rows, cols int, rowPerm, colPerm []int) float64 {
	var (
		eE, mod, exp float64
		s, o, posMod int
	)

	exp = 2

	eE = 0
	for o = 0; o < cols; o++ {

		// for every species (column) find its modal value
		mod = 0
		for s = 0; s < rows; s++ {
			x := float64(matrix[rowPerm[s]][rowPerm[o]])
			if mod < x {
				mod = x
				posMod = s
			}
		}
		// and use it to calc contribution to energy
		for s = 0; s < rows; s++ {
			x := float64(matrix[rowPerm[s]][rowPerm[o]])
			d := math.Abs(math.Pow(float64(s-posMod), exp)) // like in Minkowski metric
			eE += x * d
		}
	}
	eE = math.Pow(eE, 1/exp)

	return eE
}

// Psi4 computes energy of the permuted matrix according to Cejchan (unpublished) to seriate Anti-Robinson matrix
func Psi4(matrix [][]int, rows, cols int, rowPerm, colPerm []int) float64 {
	var (
		eE, mod      float64
		s, o, posMod int
	)

	eE = 0
	for o = 0; o < cols; o++ {

		// for every species (column) find its modal value
		mod = 0
		for s = 0; s < rows; s++ {
			x := 1 / (float64(matrix[rowPerm[s]][rowPerm[o]]) + 1) // turn distance to similarity
			if mod < x {
				mod = x
				posMod = s
			}
		}
		// and use it to calc contribution to energy
		for s = 0; s < rows; s++ {
			x := 1 / (float64(matrix[rowPerm[s]][rowPerm[o]]) + 1) // turn distance to similarity
			d := math.Abs(float64(s - posMod))
			eE += x * d
		}
	}
	return eE
}
