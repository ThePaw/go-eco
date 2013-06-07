// Copyright 2012 The Gt Authors. All rights reserved. See the LICENSE file.
package ser

// Common functions for  the Quadratic Assignment Problem.

func cost(a, b IntMatrix, p IntVector) (c int) {
	c = 0
	for i := 0; i < p.Len(); i++ {
		for j := 0; j < p.Len(); j++ {
			c += a[i][j] * b[p[i]][p[j]]
		}
	}
	return c
}

func delta(a, b IntMatrix, p IntVector, r, s int) (d int) {
	d = (a[r][r]-a[s][s])*(b[p[s]][p[s]]-b[p[r]][p[r]]) +
		(a[r][s]-a[s][r])*(b[p[s]][p[r]]-b[p[r]][p[s]])
	for i := 0; i < p.Len(); i++ {
		if i != r && i != s {
			d += (a[i][r]-a[i][s])*(b[p[i]][p[s]]-b[p[i]][p[r]]) +
				(a[r][i]-a[s][i])*(b[p[s]][p[i]]-b[p[r]][p[i]])
		}
	}
	return d
}

// Cost difference if elements i and j  are swapped in permutation (solution) p, 
// but the value of dist[i][j] is supposed to
// be known before the transposition of elements r and s. 
func delta_part(a, b, dist IntMatrix, p IntVector, i, j, r, s int) int {
	return dist[i][j] + (a[r][i]-a[r][j]+a[s][j]-a[s][i])*
		(b[p[s]][p[i]]-b[p[s]][p[j]]+b[p[r]][p[j]]-b[p[r]][p[i]]) +
		(a[i][r]-a[j][r]+a[j][s]-a[i][s])*
			(b[p[i]][p[s]]-b[p[j]][p[s]]+b[p[j]][p[r]]-b[p[i]][p[r]])
}
