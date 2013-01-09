package ser

//   Anti-Robinson Seriation
//   simulated annealing algorithm 
//   by Brusco, M., Kohn, H.F., and Stahl, S. 

import (
	"fmt"
	"math"
	"math/rand"
)

var verbose bool

//   cool = 0.5
//   tMin = 0.1
//   nRep = 20

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

// ARSA implements Anti-Robinson Seriation of distance matrix using simulated annealing algorithm by Brusco, Kohn and Stahl. 
// Returns (quasi)optimal permutation of matrix rows/cols. 
func ARSA(n int, a [][]float64, cool, tMin float64, nRep int) (perm []int) {
	var (
		count, i1, j1, k, l, loop, m, nLoop, q, set, unsel                                     int
		delta, dummy, eps, rCrit, rule, s1, span, span2, sum, temp, tMax, z, zBest, zMax, zMin float64
	)
	if verbose {
		fmt.Println("Anti-Robinson seriation by simulated annealing")
		fmt.Println("based on arsa.f by Brusco, M., Kohn, H.F.,and Stahl, S. (2007)")
		fmt.Println("cool = ", cool)
		fmt.Println("tMin = ", tMin)
		fmt.Println("nRep= ", nRep)
	}

	r1 := make([]float64, n*n/2)
	r2 := make([]float64, n*n/2)
	d := make2DsliceFloat64(n, n)
	u := make([]int, n)
	s := make([]int, n)
	sb := make([]int, n)
	t := make2DsliceInt(100, n)
	perm = make([]int, n)

	rule = .5
	eps = 1e-8

	count = 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			count++
			d[i][j] = float64(j - i)
			d[j][i] = d[i][j]
			r1[count] = d[i][j]
			r2[count] = a[i][j]
		}
	}

	for i := 0; i < count-1; i++ {
		for j := i + 1; j < count; j++ {
			if r1[j] > r1[i] {
				dummy = r1[j]
				r1[j] = r1[i]
				r1[i] = dummy
			}
			if r2[j] > r2[i] {
				dummy = r2[j]
				r2[j] = r2[i]
				r2[i] = dummy
			}
		}
	}

	sum = 0
	for i := 0; i <= count; i++ {
		sum += r1[i] * r2[i]
	}

	for iii := 0; iii < nRep; iii++ {
		for i := 0; i < n; i++ {
			u[i] = i
			t[iii][i] = 0
		}
		unsel = n - 1
		for i := 0; i < n; i++ {
			if unsel == 0 { // ++pac
				i1 = 0
			}else{
				i1 = rand.Intn(unsel)
			}
			if set > unsel {
				set = unsel
			}
			t[iii][i] = u[set]
			for j := set; j < unsel; j++ {
				u[j] = u[j+1]
			}
				unsel--
		}
	}

	zMin = 9.9e20
	zMax = 0

	for iii := 0; iii < nRep; iii++ {
		for i := 0; i < n; i++ {
			s[i] = t[iii][i]
		}
		z = 0
		for i := 0; i < n-1; i++ {
			k = s[i]
			for j := i + 1; j < n; j++ {
				l = s[j]
				z += d[i][j] * a[k][l]
			}
		}

		zBest = z
		tMax = 0
		for lll := 0; lll < 5000; lll++ {
			i1 = rand.Intn(n - 1)
			if i1 > n {
				i1 = n
			}
			//L199:
			j1 = i1
			for i1 == j1 {
				j1 = rand.Intn(n - 1)
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

			for l1 := 0; l1 < n; l1++ {
				if !(i1 == l1 || j1 == l1) {
					l = s[l1]
					delta += (d[l1][i1] - d[l1][j1]) * (a[l][m] - a[l][k])
				}
			}
			if delta < 0 {
				if math.Abs(delta) > tMax {
					tMax = math.Abs(delta)
				}
			}
		}

		loop = n * 100
		nLoop = int(math.Floor((math.Log(tMin) - math.Log(tMax)) / math.Log(cool)))

		if verbose {
			fmt.Println("Steps needed:  ", nLoop)
		}

		temp = tMax
		for i := 0; i < n; i++ {
			sb[i] = s[i]
		}

		for ijk := 0; ijk < nLoop; ijk++ {
			if verbose {
				fmt.Printf("Temp = ", temp)
			}

			for kkk := 0; kkk < loop; kkk++ {
				s1 = rand.Float64()
				if s1 <= rule {
					// interchange, insertion, or both
					i1 = rand.Intn(n - 1)
					if i1 > n {
						i1 = n
					}
					j1 = i1
					for i1 == j1 {
						j1 = rand.Intn(n - 1)
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

					for l1 := 0; l1 < n; l1++ {
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
						if z > zBest {
							zBest = z
							for i := 0; i < n; i++ {
								sb[i] = s[i]
							}
						}
					} else {
						s1 = rand.Float64()
						rCrit = math.Exp(delta / temp)
						if s1 <= rCrit {
							z += delta
							s[i1] = m
							s[j1] = k
						}
					}
				} else {
					// insertion 
					i1 = rand.Intn(n - 1)
					// object position is i1 (base zero)
					if i1 > n-1 {
						i1 = n - 1
					}
					j1 = i1
					for i1 == j1 {
						j1 = rand.Intn(n - 1)
						if j1 > n-1 {
							j1 = n - 1
						}
					}

					k = s[i1]
					delta = 0
					if j1 > i1 {
						span = float64(j1 - i1)
						for l = i1 + 1; l <= j1; l++ {
							q = s[l]
							for i := j1 + 1; i < n; i++ {
								m = s[i]
								delta += a[m][q]
							}
							for i := 0; i < i1-1; i++ {
								m = s[i]
								delta -= a[m][q]
							}
						}

						for i := 0; i < i1-1; i++ {
							m = s[i]
							delta += span * a[m][k]
						}

						for i := j1 + 1; i < n; i++ {
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
							for i := i1 + 1; i < n; i++ {
								m = s[i]
								delta -= a[m][q]
							}
							for i := 0; i <= j1-1; i++ {
								m = s[i]
								delta += a[m][q]
							}
						}

						for i := 0; i < j1-1; i++ {
							m = s[i]
							delta -= span * a[m][k]
						}

						for i := i1 + 1; i < n; i++ {
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
						if z > zBest {
							zBest = z
							for i := 0; i < n; i++ {
								sb[i] = s[i]
							}
						}
					} else {
						s1 = rand.Float64()
						rCrit = math.Exp(delta / temp)
						if s1 <= rCrit {
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
		if zBest < zMin {
			zMin = zBest
		}
		if zBest > zMax {
			zMax = zBest
			for i := 0; i < n; i++ {
				perm[i] = sb[i]
			}
		}
	}
	return
}
