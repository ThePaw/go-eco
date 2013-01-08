package ser

//   Anti-Robinson Seriation
//   simulated annealing algorithm 
//   by Brusco, M., Kohn, H.F., and Stahl, S. 

// To do:
// slices should start from 0, not 1

import (
	"fmt"
	"math"
	"math/rand"
)

//   cool = 0.5
//   tmin = 0.1
//   nreps = 20

// create2dsliceInt makes [][]int
func create2dsliceInt(dimensionX, dimensionY int) [][]int {
	_2d := make([][]int, dimensionX)
	for i := 0; i < dimensionX; i++ {
		_2d[i] = make([]int, dimensionY)
	}
	return _2d
}

// create2dsliceFloat64 makes [][]float64
func create2dsliceFloat64(dimensionX, dimensionY int) [][]float64 {
	_2d := make([][]float64, dimensionX)
	for i := 0; i < dimensionX; i++ {
		_2d[i] = make([]float64, dimensionY)
	}
	return _2d
}

// ARSA implements Anti-Robinson Seriation of distance matrix using simulated annealing algorithm by Brusco, Kohn and Stahl. 
func ARSA(n int, a [][]float64, cool, tmin float64, nreps int, verbose bool) (iperm []int) {
	var (
		k, l, m, q, i1, j1, ict, ijk, kkk, iset, iloop, unsel, nloop                           int
		s1, eps, span, asum, temp, rule, rdum, tmax, zmin, zmax, span2, delta, rcrit, z, zbest float64
	)
	if verbose {
		fmt.Println("Anti-Robinson seriation by simulated annealing")
		fmt.Println("based on arsa.f by Brusco, M., Kohn, H.F.,and Stahl, S. (2007)")
		fmt.Println("COOL = ", cool)
		fmt.Println("TMIN = ", tmin)
		fmt.Println("NREPS= ", nreps)
	}

	w := n + 1
	r1 := make([]float64, w*w/2)
	r2 := make([]float64, w*w/2)
	d := create2dsliceFloat64(w, w)
	u := make([]int, w)
	s := make([]int, w)
	sb := make([]int, w)
	t := create2dsliceInt(100, w)
	iperm = make([]int, w)

	rule = .5
	ict = 0
	for i := 1; i < n-1; i++ {
		for j := i + 1; j <= n; j++ {
			ict++
			d[i][j] = float64(j - i)
			d[j][i] = d[i][j]
			r1[ict] = d[i][j]
			r2[ict] = a[i][j]
		}
	}

	for i := 1; i <= ict-1; i++ {
		for j := i + 1; j <= ict; j++ {
			if r1[j] > r1[i] {
				rdum = r1[j]
				r1[j] = r1[i]
				r1[i] = rdum
			}
			if r2[j] > r2[i] {
				rdum = r2[j]
				r2[j] = r2[i]
				r2[i] = rdum
			}
		}
	}

	asum = 0
	for i := 1; i <= ict; i++ {
		asum += r1[i] * r2[i]
	}
	eps = 1e-8

	for iii := 1; iii <= nreps; iii++ {
		for i := 1; i <= n; i++ {
			u[i] = i
			t[iii][i] = 0
		}
		unsel = n
		for i := 1; i <= n; i++ {
			//s1 = rand.Float64()
			i1 = rand.Intn(unsel) + 1
			if iset > unsel {
				iset = unsel
			}
			t[iii][i] = u[iset]
			// 	    out of bounds error reported by Rohan Shah (9/13/12) 
			for j := iset; j <= unsel-1; j++ {
				u[j] = u[j+1]
			}
			unsel--
			// L1: 
		}
		// L999: 
	}

	zmin = 9.9e20
	zmax = 0

	for iii := 1; iii <= nreps; iii++ {
		for i := 1; i <= n; i++ {
			s[i] = t[iii][i]
		}
		z = 0
		for i := 1; i <= n-1; i++ {
			k = s[i]
			for j := i + 1; j <= n; j++ {
				l = s[j]
				z += d[i][j] * a[k][l]
			}
		}

		zbest = z
		tmax = 0
		for lll := 1; lll <= 5000; lll++ {
			s1 = rand.Float64()
			i1 = rand.Intn(n) + 1
			if i1 > n {
				i1 = n
			}
			//L199:
			j1 = i1
			for i1 == j1 {
				s1 = rand.Float64()
				j1 = rand.Intn(n) + 1
				if j1 > n {
					j1 = n
				}
			}

			if i1 > j1 {
				// swap
				i1, j1 = j1, i1
			}

			k = s[i1]
			m = s[j1]
			delta = 0.

			for l1 := 1; l1 <= n; l1++ {
				if !(i1 == l1 || j1 == l1) {
					l = s[l1]
					delta += (d[l1][i1] - d[l1][j1]) * (a[l][m] - a[l][k])
				}
			}
			if delta < 0 {
				if math.Abs(delta) > tmax {
					tmax = math.Abs(delta)
				}
			}
		}

		iloop = n * 100
		nloop = int(math.Floor((math.Log(tmin) - math.Log(tmax)) / math.Log(cool)))

		if verbose {
			fmt.Println("Steps needed:  ", nloop)
		}

		temp = tmax
		for i := 1; i <= n; i++ {
			sb[i] = s[i]
		}

		for ijk = 1; ijk <= nloop; ijk++ {
			if verbose {
				fmt.Printf("Temp = ", temp)
			}

			for kkk = 1; kkk <= iloop; kkk++ {
				s1 = rand.Float64()
				if s1 <= rule {
					// interchange, insertion, or both
					i1 = rand.Intn(n) + 1
					if i1 > n {
						i1 = n
					}
					j1 = i1
					for i1 == j1 {
						j1 = rand.Intn(n) + 1
						if j1 > n {
							j1 = n
						}
					}

					if i1 > j1 {
						//swap
						i1, j1 = j1, i1
					}

					k = s[i1]
					m = s[j1]
					delta = 0.

					for l1 := 1; l1 <= n; l1++ {
						if !(i1 == l1 || j1 == l1) {
							l = s[l1]
							delta += (d[l1][i1] - d[l1][j1]) * (a[l][m] - a[l][k])
							//L250:
						}
					}

					if delta > -eps {
						z += delta
						s[i1] = m
						s[j1] = k
						if z > zbest {
							zbest = z
							for i := 1; i <= n; i++ {
								sb[i] = s[i]
							}
						}
					} else {
						s1 = rand.Float64()
						rcrit = math.Exp(delta / temp)
						if s1 <= rcrit {
							z += delta
							s[i1] = m
							s[j1] = k
						}
					}
				} else {
					// insertion 
					i1 = rand.Intn(n) + 1
					// object position is I1 
					if i1 > n {
						i1 = n
					}
					j1 = i1
					for i1 == j1 {
						j1 = rand.Intn(n) + 1
						if j1 > n {
							j1 = n
						}
					}

					k = s[i1]
					delta = 0
					if j1 > i1 {
						span = float64(j1 - i1)
						for l = i1 + 1; l <= j1; l++ {
							q = s[l]
							for i := j1 + 1; i <= n; i++ {
								m = s[i]
								delta += a[m][q]
							}
							for i := 1; i <= i1-1; i++ {
								m = s[i]
								delta -= a[m][q]
							}
						}

						for i := 1; i <= i1-1; i++ {
							m = s[i]
							delta += span * a[m][k]
						}

						for i := j1 + 1; i <= n; i++ {
							m = s[i]
							delta -= span * a[k][m]
						}

						span2 = span + 1
						for i := i1 + 1; i <= j1; i++ {
							span2 += -2
							m = s[i]
							delta += span2 * a[k][m]
						}
					} else {
						span = float64(i1 - j1)
						for l = j1; l <= i1-1; l++ {
							q = s[l]
							for i := i1 + 1; i <= n; i++ {
								m = s[i]
								delta -= a[m][q]
							}
							for i := 1; i <= j1-1; i++ {
								m = s[i]
								delta += a[m][q]
							}
						}

						for i := 1; i <= j1-1; i++ {
							m = s[i]
							delta -= span * a[m][k]
						}

						for i := i1 + 1; i <= n; i++ {
							m = s[i]
							delta += span * a[k][m]
						}

						span2 = span + 1
						for i := j1; i <= i1-1; i++ {
							span2 -= 2
							m = s[i]
							delta -= span2 * a[k][m]
						}
					}
					if delta > -eps {
						z += delta
						if j1 > i1 {
							for l = i1; l <= j1-1; l++ {
								s[l] = s[l+1]
							}
							s[j1] = k
						} else {
							for l = i1; l >= j1+1; l-- {
								s[l] = s[l-1]
							}
							s[j1] = k
						}
						if z > zbest {
							zbest = z
							for i := 1; i <= n; i++ {
								sb[i] = s[i]
							}
						}
					} else {
						s1 = rand.Float64()
						rcrit = math.Exp(delta / temp)
						if s1 <= rcrit {
							z += delta
							if j1 > i1 {
								for l = i1; l <= j1-1; l++ {
									s[l] = s[l+1]
								}
								s[j1] = k
							} else {
								for l = i1; l >= j1+1; l-- {
									s[l] = s[l-1]
								}
								s[j1] = k
							}
						}
					}

				}
			}
			temp *= cool
		}
		if zbest < zmin {
			zmin = zbest
		}
		if zbest > zmax {
			zmax = zbest
			for i := 1; i <= n; i++ {
				iperm[i] = sb[i]
			}
		}
	}
	return
}
