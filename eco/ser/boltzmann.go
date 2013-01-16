package ser

import (
	"math"
	"math/rand"
)

const (
	TOTALSAMPLE = 1500
	ITERATION   = 100
	T           = 1.0
	SEED        = 335
)

// Boltzmann does MCMC reordering of data matrix rows and columns to minimize energy function
// see Miklos (2005)
func Boltzmann(rows, cols int, iMatrix [][]int, iRow, iCol []int, energyFn int, permute bool) float64 {
	var (
		iRowNew, iColNew                     []int
		iRowBest, iColBest                   []int
		i, j, k, s, o, c, m                     int
		totSamp, iter                int
		eX, enOld, enNew, enBest float64
		rowDiff, colDiff                   float64
		eTo, eT                                 float64 //ETo = original, typed temperature, the temperature of the variable eT
	)

	// which energy function
	en := Psi
	switch energyFn {
	case 0:
		en = Psi
	case 1:
		en = Mirror
	}

	// initialize
	eTo = T
	totSamp = TOTALSAMPLE
	iter = ITERATION
	c = 0
	m = 0

	// allocate slices
	iRowNew = make([]int, rows)
	iColNew = make([]int, cols)
	iRowBest = make([]int, rows)
	iColBest = make([]int, cols)

	// initial permutation
	if permute {
		rand.Seed(SEED)
		PermVec(iRow)
		PermVec(iCol)
	}

	// normalization of temperature
	eT = 0.0
	for j = 0; j < cols; j++ {
		for i = 0; i < rows-1; i++ {
			for k = i + 1; k < rows; k++ {
				eT += math.Abs(float64(iMatrix[i][j]-iMatrix[k][j])) / float64(rows*cols)
			}
		}
	}

	eT = 0.0
	for i = 0; i < rows; i++ {
		for j = 0; j < cols-1; j++ {
			for k = j + 1; k < cols; k++ {
				eT += math.Abs(float64(iMatrix[i][j]-iMatrix[i][k])) / float64(rows*cols)
			}
		}
	}

	eT *= eTo
	//  printf("eT: %e\n",eT)

	enOld = en(iMatrix, rows, cols, iRow, iCol)
	enBest = enOld

	for s = 0; s < rows; s++ {
		iRowBest[s] = iRow[s]
	}
	for o = 0; o < cols; o++ {
		iColBest[o] = iCol[o]
	}

	// MCMC
	for m = 0; m < totSamp; m++ {
		for c = 0; c < iter; c++ {
			rowDiff = proposeRow(iMatrix, iRow, iRowNew, rows, cols)
			enNew = enOld + rowDiff
			eX = rand.Float64()
			if math.Exp(-(enNew-enOld)/eT) > eX {
				// accept
				enOld = enNew
				if enOld < enBest {
					enBest = enOld
					for s = 0; s < rows; s++ {
						iRowBest[s] = iRowNew[s]
					}

				}

				for s = 0; s < rows; s++ {
					iRow[s] = iRowNew[s]
				}
			}

			colDiff = proposeCol(iMatrix, iCol, iColNew, rows, cols)
			enNew = enOld + colDiff
			eX = rand.Float64()

			if math.Exp(-(enNew-enOld)/eT) > eX {
				// accept
				enOld = enNew
				if enOld < enBest {
					enBest = enOld
					for o = 0; o < cols; o++ {
						iColBest[o] = iColNew[o]
					}

				}

				for o = 0; o < cols; o++ {
					iCol[o] = iColNew[o]
				}
			}
		}
	}

	// copy the best permutations back
	for s = 0; s < rows; s++ {
		iRow[s] = iRowBest[s]
	}
	for o = 0; o < cols; o++ {
		iCol[o] = iColBest[o]
	}
	return enBest
}

// Mirror computes energy of the permuted matrix according to Miklos (2005), Eq. 2 (mirror)
func Mirror(iMatrix [][]int, rows, cols int, iRow, iCol []int) float64 {
	var (
		eE, sum1, sum2 float64
		s, o           int
	)

	sum1 = 0
	sum2 = 0
	for s = 0; s < rows-1; s++ {
		for o = 0; o < cols; o++ {
			sum1 += math.Abs(float64(iMatrix[iRow[s]][o] - iMatrix[iRow[s+1]][o]))
		}
	}

	for o = 0; o < cols; o++ {
		sum1 += math.Abs(float64(iMatrix[iRow[0]][o]-iMatrix[iRow[1]][o])) + math.Abs(float64(iMatrix[iRow[rows-1]][o]-iMatrix[iRow[rows-2]][o]))
	}

	for s = 0; s < rows; s++ {
		for o = 0; o < cols-1; o++ {
			sum2 += math.Abs(float64(iMatrix[s][iCol[o]] - iMatrix[s][iCol[o+1]]))
		}
	}

	for s = 0; s < rows; s++ {
		sum2 += math.Abs(float64(iMatrix[s][iCol[0]]-iMatrix[s][iCol[1]])) + math.Abs(float64(iMatrix[s][iCol[cols-1]]-iMatrix[s][iCol[cols-2]]))
	}

	eE = sum1 + sum2
	return eE
}

// proposeRow()
func proposeRow(iMatrix [][]int, iRow, iRowNew []int, rows, cols int) float64 {
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
			eDiff += math.Abs(float64(iMatrix[iRow[e]][iJ]-iMatrix[iRow[d-1]][iJ])) - math.Abs(float64(iMatrix[iRow[d]][iJ]-iMatrix[iRow[d-1]][iJ]))
		} else if d == 0 {
			eDiff += math.Abs(float64(iMatrix[iRow[e]][iJ]-iMatrix[iRow[e-1]][iJ])) - math.Abs(float64(iMatrix[iRow[d]][iJ]-iMatrix[iRow[d+1]][iJ]))
		} else {
			eDiff += 2.0 * (math.Abs(float64(iMatrix[iRow[e]][iJ]-iMatrix[iRow[d-1]][iJ])) - math.Abs(float64(iMatrix[iRow[d]][iJ]-iMatrix[iRow[d-1]][iJ])))
		}

		if e < rows-2 {
			eDiff += math.Abs(float64(iMatrix[iRow[d]][iJ]-iMatrix[iRow[e+1]][iJ])) - math.Abs(float64(iMatrix[iRow[e]][iJ]-iMatrix[iRow[e+1]][iJ]))
		} else if e == rows-1 {
			eDiff += math.Abs(float64(iMatrix[iRow[d]][iJ]-iMatrix[iRow[d+1]][iJ])) - math.Abs(float64(iMatrix[iRow[e]][iJ]-iMatrix[iRow[e-1]][iJ]))
		} else {
			eDiff += 2.0 * (math.Abs(float64(iMatrix[iRow[d]][iJ]-iMatrix[iRow[e+1]][iJ])) - math.Abs(float64(iMatrix[iRow[e]][iJ]-iMatrix[iRow[e+1]][iJ])))
		}
	}

	for iI = 0; iI < d; iI++ {
		iRowNew[iI] = iRow[iI]
	}
	for iI = d; iI <= e; iI++ {
		iRowNew[iI] = iRow[e+d-iI]
	}
	for iI = e + 1; iI < rows; iI++ {
		iRowNew[iI] = iRow[iI]
	}

	return eDiff
}

//  proposeCol()
func proposeCol(iMatrix [][]int, iCol, iColNew []int, rows, cols int) float64 {
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
			eDiff += math.Abs(float64(iMatrix[iI][iCol[e]]-iMatrix[iI][iCol[d-1]])) -
				math.Abs(float64(iMatrix[iI][iCol[d]]-iMatrix[iI][iCol[d-1]]))
		} else if d == 0 {
			eDiff += math.Abs(float64(iMatrix[iI][iCol[e]]-iMatrix[iI][iCol[e-1]])) -
				math.Abs(float64(iMatrix[iI][iCol[d]]-iMatrix[iI][iCol[d+1]]))
		} else {
			eDiff += 2.0 * (math.Abs(float64(iMatrix[iI][iCol[e]]-iMatrix[iI][iCol[d-1]])) -
				math.Abs(float64(iMatrix[iI][iCol[d]]-iMatrix[iI][iCol[d-1]])))
		}

		if e < cols-2 {
			eDiff += math.Abs(float64(iMatrix[iI][iCol[d]]-iMatrix[iI][iCol[e+1]])) -
				math.Abs(float64(iMatrix[iI][iCol[e]]-iMatrix[iI][iCol[e+1]]))
		} else if e == cols-1 {
			eDiff += math.Abs(float64(iMatrix[iI][iCol[d]]-iMatrix[iI][iCol[d+1]])) -
				math.Abs(float64(iMatrix[iI][iCol[e]]-iMatrix[iI][iCol[e-1]]))
		} else {
			eDiff += 2.0 * (math.Abs(float64(iMatrix[iI][iCol[d]]-iMatrix[iI][iCol[e+1]])) -
				math.Abs(float64(iMatrix[iI][iCol[e]]-iMatrix[iI][iCol[e+1]])))
		}
	}

	for iJ = 0; iJ < d; iJ++ {
		iColNew[iJ] = iCol[iJ]
	}
	for iJ = d; iJ <= e; iJ++ {
		iColNew[iJ] = iCol[d+e-iJ]
	}
	for iJ = e + 1; iJ < cols; iJ++ {
		iColNew[iJ] = iCol[iJ]
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
func Psi(iMatrix [][]int, rows, cols int, iRow, iCol []int) float64 {
	var (
		eE   float64
		s, o int
	)

	eE = 0
	for s = 0; s < rows-1; s++ {
		for o = 0; o < cols; o++ {
			x := float64(iMatrix[iRow[s]][iRow[o]])
			c := float64(cols*s) / float64(rows)
			eE += x * c
		}
	}
	return eE
}
