// Copyright 2012 The Gt Authors. All rights reserved. See the LICENSE file.

package ser

// Some handy functions. 

import (
	"math"
	"math/rand"
)

const Inf int64 = math.MaxInt64

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b int64) int64 {
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

// make2DsliceInt makes [][]int
func make2DsliceInt(rows, cols int) [][]int {
	arr2 := make([][]int, rows)
	for i := 0; i < rows; i++ {
		arr2[i] = make([]int, cols)
	}
	return arr2
}

// make2DsliceFloat64 makes [][]float64
func make2DsliceFloat64(rows, cols int) [][]float64 {
	arr2 := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		arr2[i] = make([]float64, cols)
	}
	return arr2
}

// Product computes the product of two matrices. 
func Product(a, b [][]int, rows, cols int) [][]int {
	c := make2DsliceInt(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			c[i][j] = 0
			for k := 0; k < cols; k++ {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return c
}

// CircleProduct computes circular product of two matrices. 
func CircleProduct(a, b [][]int, rows, cols int) [][]int {
	c := make2DsliceInt(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			c[i][j] = 0
			for k := 0; k < cols; k++ {
				c[i][j] += imin(a[i][k], b[k][j])
			}
		}
	}
	return c
}

// Transpose returns transposed matrix. 
func Transpose(a [][]int, rows, cols int) [][]int {
	c := make2DsliceInt(cols, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			c[i][j] = a[j][i]
		}
	}
	return c
}

// Laplacian returns the Laplacian of a symmetric matrix (adjacency matrix). 
// See Atkins et al. (1998). 
func Laplacian(adj [][]int, rows int) [][]int {
	lap := make2DsliceInt(rows, rows)
	deg := make2DsliceInt(rows, rows)
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