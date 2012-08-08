// Test of triangular inequality for distance matrix

package sim

import (
	. "code.google.com/p/go-eco/eco"
	"math"
)

// Test of triangular inequality for distance matrix
// Returns 0 if holds, or positive integer = number of violations
func TriIneq(dist *Matrix) int {
	n := dist.R
	ineq := 0
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; j < n; j++ {
				x := dist.Get(j, i)
				y := dist.Get(k, i)
				z := dist.Get(k, j)

				lng := x
				lng = math.Max(lng, y)
				lng = math.Max(lng, z)
				if lng > x+y+z-lng {
					ineq++
				}
			}
		}
	}
	return ineq
}
