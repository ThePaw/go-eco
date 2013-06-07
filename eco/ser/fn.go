// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package ser

// Some handy functions. 

import (
	"math"
	"math/rand"
)

const iInf int = math.MaxInt32
const inf float64 = math.MaxFloat64

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func imin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func imax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func iSign(a int) int {
	if a == 0 {
		return 0
	}
	if a < 0 {
		return -1
	}
	return 1
}

// See Hubert et al. (2001).
func sign(a float64) int {
	if a == 0 {
		return 0
	}
	if a < 0 {
		return -1
	}
	return 1
}

func iRound(a float64) int {
	return int(math.Floor(a + 0.5))
}

// returns absolute value of an integer
func iAbs(n int) int {
	switch {
	case n < 0:
		return -n
	case n == 0:
		return 0 // return correctly abs(-0)
	}
	return n
}

func cube(x float64) float64 {
	return x * x * x
}

// Uniform random number. 
func unif(low, high int) int {
	return low + rand.Intn(high-low+1)
}

// Uniform random number. 
func unif64(low, high int64) int64 {
	return low + rand.Int63n(high-low+1)
}

// Laplacian returns the Laplacian of a symmetric matrix (adjacency matrix). 
// See Atkins et al. (1998). 
func Laplacian(adj IntMatrix) IntMatrix {
	if !adj.IsSymmetric() {
		panic("matrix not symmetric")
	}
	rows := adj.Rows()
	lap := NewIntMatrix(rows, rows)
	deg := NewIntMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		sum := 0
		for j := 0; j < rows; j++ {
			sum += adj[i][j]
		}
		deg[i][i] = sum
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			lap[i][j] = deg[i][j] - adj[i][j]
		}
	}
	return lap
}

/*
// NormLaplacian returns the normalized Laplacian of a symmetric matrix (adjacency matrix). 
// See  Weisstein, Eric W. "Laplacian Matrix." From MathWorld--A Wolfram Web Resource. http://mathworld.wolfram.com/LaplacianMatrix.html 
func NormLaplacian(adj [][]int, rows int) [][]int {
	lap := make2DsliceInt(rows)
	deg := make2DsliceInt(rows)


To be implemented
	return lap
}
*/

// Adj2Sim constructs similarity matrix from adjacency matrix via circle product.
// See Kendall, 1971; Atkins et al. 1998: 309.
func Adj2Sim(a Matrix64, permuteWhat string) Matrix64 {
	var s Matrix64
	t := a.Transpose()
	switch permuteWhat {
	case "rows":
		s = a.CircleProduct(t)
	case "cols":
		s = t.CircleProduct(a)
	}
	return s
}

// reverses the permutation if it is decreasing
func reverseIfNeeded(perm IntVector) {
	smp := len(perm)
	half := smp / 2

	// sum labels up to half-length
	sum := 0
	for i := 0; i < half; i++ {
		sum += perm[i]
	}

	if sum > half*smp/2 {
		// reverse the series
		for i := 0; i < half; i++ {
			tmp := perm[i]
			perm[i] = perm[smp-i-1]
			perm[smp-i-1] = tmp
		}
	}
	return
}

// reverses the permutation if it is decreasing, using Spearman rho rank correlation coefficient
func reverseIfNeeded2(perm IntVector) float64 {
	smp := len(perm)
	truth := make(IntVector, smp)
	ranks1 := make(IntVector, smp)
	ranks2 := make(IntVector, smp)

	for i := 0; i < smp; i++ {
		truth[i] = i
	}

	// calculate ranks
	for i := 0; i < smp; i++ {
		// count scores lower than this
		count := 0
		for k := 0; k < smp; k++ {
			if truth[k] <= truth[i] {
				count++
			}
		}
		ranks1[i] = count
	}

	for i := 0; i < smp; i++ {
		// count scores lower than this
		count := 0
		for k := 0; k < smp; k++ {
			if perm[k] <= perm[i] {
				count++
			}
		}
		ranks2[i] = count
	}

	// Spearman Rho
	sumd2 := 0
	for k := 0; k < smp; k++ {
		sumd2 += (ranks1[k] - ranks2[k]) * (ranks1[k] - ranks2[k])
	}
	rho := 1.0 - 6.0*float64(sumd2)/float64(smp*smp*smp-smp)
	if rho < 0 {
		// reverse the permutation
		half := smp / 2
		for i := 0; i < half; i++ {
			tmp := perm[i]
			perm[i] = perm[smp-i-1]
			perm[smp-i-1] = tmp
		}
	}
	return rho
}

func addToRankHistogram(perm IntVector, h IntMatrix) {
	smp := len(perm)
	for i := 0; i < smp; i++ {
		for j := 0; j < smp; j++ {
			if perm[j] == i { // find rank of sample i
				h[i][j]++
				break
			}
		}
	}
	return
}

func addToRhoHistogram(rho float64, h IntVector) {
	var i int
	bins := h.Len()
	for i = 0; i < bins; i++ {
		if rho <= float64(i+1)/float64(bins) {
			h[i]++
			break
		}
	}
	return
}

// addToPairOrderHistogram updates Pair-Order matrix
func addToPairOrderHistogram(perm IntVector, h IntMatrix) {
	smp := len(perm)
	for i := 0; i < smp; i++ {
		for j := i + 1; j < smp; j++ {
			h[perm[i]][perm[j]]++
		}
	}
	return
}

/*
// sim2dist converts similarity matrix to  distance matrix (ad hoc !!!)
func sim2dist(mat Matrix64, lambda float64) {
	//distance( x , y ) = -log{λ ⋅ similarity( x , y )}
	for i, row := range mat {
		for j, val := range row {
			if val == 0 {
			mat[i][j] = inf

}else{
			mat[i][j] = -math.Log(lambda * val)
}
		}
	}
}
*/

func isOdd(x int) bool {
	if x%2 == 0 {
		return false
	}
	return true
}

// compareFloat64 returns true iff a equals b.
//
// Two floating point numbers are assumed to be equal if
// absolute error |a-b| < 1e-10
// relative error < 1e-10
//
// compareFloat64(a,b) == compareFloat64(b,a)
// code by Lennart Oymanns
// https://github.com/lmcoy/crossx/blob/master/math/linalg/mathutil.go
func compareFloat64(a, b float64) bool {
	const (
		epsilon = 1.0e-10
	)
	var relError float64
	if a == b {
		return true
	}
	// check absolute error
	if math.Abs(a-b) < epsilon {
		return true
	}
	// check relative error
	// do this since an error of 1 is very large for a value near 0
	// but it is small enough to assume two values to be equal
	// if the values are really large.
	// use two cases two ensure that compareFloat64(a,b) == compareFloat64(b,a) in any case.
	if math.Abs(b) > math.Abs(a) {
		relError = math.Abs((a - b) / b)
	} else {
		relError = math.Abs((a - b) / a)
	}
	if relError <= epsilon {
		return true
	}
	return false
}
