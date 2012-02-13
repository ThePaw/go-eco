// Mountford dissimilarity and similarity

package sim

import (
	. "go-eco.googlecode.com/hg/eco"
	. "math"
)

func mount_fun(theta, j, a, b float64) float64 {
	return Exp(theta*a) + Exp(theta*b) - Exp(theta*(a+b-j)) - 1
}

func mount_der(theta, j, a, b float64) float64 {
	return a*Exp(theta*a) + b*Exp(theta*b) - (a+b-j)*Exp(theta*(a+b-j))
}

// Mountford distance matrix
// Mountford index is defined as M = 1/α where α is the parameter of Fisher's logseries 
// assuming that the compared communities are samples from the same community. 
// The index M is found as the positive root of equation exp(a*M) + exp(b*M) = 1 + exp((a+b-j)*M), 
// where j is the number of species occurring in both communities, and a and b are the number of species 
// in each separate community (so the index uses presence–absence information). 
// Mountford index is usually misrepresented in the literature: indeed Mountford (1962) suggested 
// an approximation to be used as starting value in iterations, but the proper index is defined as the root of the equation above. 
// The function solves M with the Newton method. Please note that if either a or b are equal to j, 
// one of the communities could be a subset of other, and the dissimilarity is 0 meaning that non-identical objects may be regarded 
// as similar and the index is non-metric. The Mountford index is in the range 0 ... log(2), but the dissimilarities are divided by log(2) 
// so that the results will be in the conventional range 0 ... 1. 
func Mountford_D(data *Matrix) *Matrix {
	const (
		maxit = 20
		ε     = 1e-12
		tol   = 1e-5
	)

	rows := data.R
	cols := data.C
	out := NewMatrix(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sim := 0
			t1 := 0
			t2 := 0
			count := 0
			v := 0.0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				if x > 0.0 && y > 0.0 {
					sim++
				}
				if x > 0.0 {
					t1++
				}
				if y > 0.0 {
					t2++
				}
				count++
			}
			if count == 0 {
				panic("NaN")
			}
			if t1 == 0 || t2 == 0 {
				v = NaN()
			} else if sim == 0 {
				v = 0
			} else if sim == t1 || sim == t2 {
				v = Log(2.0)
			} else {
				jjj := float64(sim)
				aaa := float64(t1)
				bbb := float64(t2)
				v = 2 * jjj / (2*aaa*bbb - (aaa+bbb)*jjj)
				for k := 0; k < maxit; k++ {
					oldist := v
					v -= mount_fun(v, jjj, aaa, bbb) / mount_der(v, jjj, aaa, bbb)
					if Abs(oldist-v)/oldist < tol || Abs(oldist-v) < ε {
						break
					}
				}
			}

			v = 1 - v/Log(2.0)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Mountford similarity matrix, for boolean data
func MountfordBool_S(data *Matrix) *Matrix {
	var (
		a, b, c, v float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = GetABCD(data, i, j)
			if (a*(b+c) + (2 * b * c)) != 0 {
				v = 2 * a / (a*(b+c) + (2 * b * c))
			} else {
				v = 2 * a / (a*(b+c) + (2 * b * c))

			}
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
