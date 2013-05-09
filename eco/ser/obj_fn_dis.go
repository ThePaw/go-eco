// Copyright 2012 The Gt Authors. All rights reserved. See the LICENSE file.

package ser

// Objective (loss and gain) functions for distance (dissimilarity) matrices. 

import (
//"fmt"
	"math"
)

// g1Gain returns gain of the permuted matrix according to Hubert, Arabie & Meulman 2001, Chapter 4; see Brusco 2002: 50, Eq. 6. (WRUG)
func g1Gain(dis Matrix64, p IntVector) float64 {
	if !dis.IsSymmetric() {
		panic("distance matrix not symmetric")
	}
	n := p.Len()
	if dis.Rows() != n {
		panic("dimensions not equal")
	}

	c := 0
	for k := 0; k < n-2; k++ {
		for l := k + 1; l < n-1; l++ {
			for m := l + 1; m < n; m++ {
				x := dis[p[k]][p[m]]
				y := dis[p[k]][p[l]]
				c += sign(x - y)
			}
		}
	}
	return float64(c)
}

// g2Gain returns gain of the permuted matrix according to Hubert, Arabie & Meulman 2001, Chapter 4; see Brusco 2002: 50, Eq. 7. (WRCUG)
func g2Gain(dis Matrix64, p IntVector) float64 {
	if !dis.IsSymmetric() {
		panic("distance matrix not symmetric")
	}
	n := p.Len()
	if dis.Rows() != n {
		panic("dimensions not equal")
	}

	c := 0
	for k := 0; k < n-2; k++ {
		for l := k + 1; l < n-1; l++ {
			for m := l + 1; m < n; m++ {
				x := dis[p[k]][p[m]]
				y := dis[p[k]][p[l]]
				c += sign(x - y)
				y = dis[p[l]][p[m]]
				c += sign(x - y)
			}
		}
	}
	return float64(c)
}

// g3Gain returns gain of the permuted matrix according to Hubert, Arabie & Meulman 2001, Chapter 4; see Brusco 2002: 50, Eq. 8. (WRWG)
func g3Gain(dis Matrix64, p IntVector) float64 {
	if !dis.IsSymmetric() {
		panic("distance matrix not symmetric")
	}
	n := p.Len()
	if dis.Rows() != n {
		panic("dimensions not equal")
	}

	c := 0.0
	for k := 0; k < n-2; k++ {
		for l := k + 1; l < n-1; l++ {
			for m := l + 1; m < n; m++ {
				x := dis[p[k]][p[m]]
				y := dis[p[k]][p[l]]
				c += x - y
			}
		}
	}
	return c
}

// g4Gain returns gain of the permuted matrix according to Hubert, Arabie & Meulman 2001, Chapter 4; see Brusco 2002: 50, Eq. 9. (WRCWG)
func g4Gain(dis Matrix64, p IntVector) float64 {
	if !dis.IsSymmetric() {
		panic("distance matrix not symmetric")
	}
	n := p.Len()
	if dis.Rows() != n {
		panic("dimensions not equal")
	}

	c := 0.0
	for k := 0; k < n-2; k++ {
		for l := k + 1; l < n-1; l++ {
			for m := l + 1; m < n; m++ {
				x := dis[p[k]][p[m]]
				y := dis[p[k]][p[l]]
				z := dis[p[l]][p[m]]
				c += 2*x - y - z
			}
		}
	}
	return c
}

// hGain returns gain of the permuted matrix according to Szczotka 1972; see Brusco et al. 2008: 507, Eq. 7.
func hGain(dis Matrix64, p IntVector) float64 {
	if !dis.IsSymmetric() {
		panic("distance matrix not symmetric")
	}
	n := p.Len()
	if dis.Rows() != n {
		panic("dimensions not equal")
	}

	c := 0.0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			d := math.Abs(float64(i - j))
			x := dis[p[i]][p[j]]
			c += d * x
		}
	}
	return c
}

// hNormGain returns gain of the permuted matrix according to Szczotka 1972; see Brusco et al. 2008: 507-508, Eq. 7.
// TO BE IMPLEMENTED

func optimize(dis Matrix64, p IntVector) {
	// TO BE IMPLEMENTED
}

// strengLoss returns a count of Anti-Robinson events (Streng and Schoenfelder 1978; Chen 2002:21).
func strengLoss(dis Matrix64, p IntVector, which int) float64 {
	//which indicates the weighing scheme
	// 1 ... no weighting (i)
	// 2 ... abs. deviations (s)
	// 3 ... weighted abs. deviations (w)

	if !dis.IsSymmetric() {
		panic("distance matrix not symmetric")
	}
	n := p.Len()
	if dis.Rows() != n {
		panic("dimensions not equal")
	}

	sum := 0.0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			d_ij := dis[p[i]][p[j]]
			for k := 0; k < n; k++ {
				d_ik := dis[p[i]][p[k]]
				if j < k && k < i {
					if d_ij < d_ik {

						switch which {
						case 1:
							sum++
						case 2:
							sum += math.Abs(d_ij - d_ik)
						case 3:
							sum += math.Abs(float64(p[j]-p[k])) * math.Abs(d_ij-d_ik)
						}

					}
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			d_ij := dis[p[i]][p[j]]
			for k := 0; k < n; k++ {
				d_ik := dis[p[i]][p[k]]
				if i < j && j < k {

					if d_ij > d_ik {

						switch which {
						case 1:
							sum++
						case 2:
							sum += math.Abs(d_ij - d_ik)
						case 3:
							sum += math.Abs(float64(p[j]-p[k])) * math.Abs(d_ij-d_ik)
						}

					}
				}
			}
		}
	}
	return sum
}

// StrengLoss1 returns a count of Anti-Robinson events, no weighting (Streng and Schoenfelder 1978; Chen 2002:21).
func StrengLoss1(dis Matrix64, p IntVector) float64 {
	return strengLoss(dis, p, 1)
}

// StrengLoss2 returns a count of Anti-Robinson events, weighted by abs. deviations (Streng and Schoenfelder 1978; Chen 2002:21).
func StrengLoss2(dis Matrix64, p IntVector) float64 {
	return strengLoss(dis, p, 2)
}

// StrengLoss3 returns a count of Anti-Robinson events, weighted by weighted abs. deviations (Streng and Schoenfelder 1978; Chen 2002:21).
func StrengLoss3(dis Matrix64, p IntVector) float64 {
	return strengLoss(dis, p, 3)
}

// InertiaGain returns the Inertia criterion (Caraux and Pinloche 2005).
func InertiaGain(dis Matrix64, p IntVector) float64 {
	if !dis.IsSymmetric() {
		panic("distance matrix not symmetric")
	}
	n := p.Len()
	if dis.Rows() != n {
		panic("dimensions not equal")
	}

	sum := 0.0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum += dis[p[i]][p[j]] * math.Abs(float64((i-j)*(i-j)))
		}
	}
	return sum
}

// LeastSquaresLoss returns the Least Squares criterion (Caraux and Pinloche 2005).
func LeastSquaresLoss(dis Matrix64, p IntVector) float64 {
	if !dis.IsSymmetric() {
		panic("distance matrix not symmetric")
	}
	n := p.Len()
	if dis.Rows() != n {
		panic("dimensions not equal")
	}

	sum := 0.0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			incr := dis[p[i]][p[j]] - math.Abs(float64(i-j))
			incr *= incr
			sum += incr
		}
	}
	return sum
}

// Brusco2008 implements submatrix optimization.
// see Brusco et al. 2008: 509.
func Brusco2008(dis Matrix64, p IntVector, v int) {
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
	psiSub := NewIntVector(v)
	psiSub.Perm()
	δ := NewIntVector(v)
	δ.Perm()

	improved := true
	for improved {
		improved = false
		for i := 0; i < n-v+1; i++ {

			// Step 1a
			for k := i; k < n-v+1; k++ {
				psiSub[k-i] = p[k]
			}

			// Step 1b
			for k := i; k < n-v+1; k++ {
				for l := i; l < n-v+1; l++ {
					q[p[k-i]][p[l-i]] = dis[p[k]][p[l]]
				}
			}

			// Step 1c
			optimize(q, δ)

			// Step 1d
			for k := i; k < n-v+1; k++ {
				p[k] = psiSub[δ[k-i]]
			}

			// Step 1e
			for k := 0; k < v; k++ {
				if δ[k] != k {
					improved = true
					break
				}
			}
		}
	}
}

// MooreStressDisLoss returns the Moore Stress criterion (Niermann 2005:42, Eq. 1, 2) for a distance matrix.
func MooreStressDisLoss(dis Matrix64, p IntVector) float64 {
	return MooreStressLoss(dis, p, p)
}

// VonNeumannStressDisLoss returns the Moore Stress criterion (Niermann 2005:42) for a distance matrix.
func VonNeumannStressDisLoss(dis Matrix64, p IntVector) float64 {
	return VonNeumannStressLoss(dis, p, p)
}

// GARLoss returns the GAR(w) (Wu 2010: 773) generalized anti-Robinson loss function for a distance matrix.
func GARLoss(dis Matrix64, p IntVector, w int) float64 {
	if !dis.IsSymmetric() {
		panic("distance matrix not symmetric")
	}
	n := p.Len()
	if dis.Rows() != n {
		panic("dimensions not equal")
	}

	sum := 0.0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			dij := dis[i][j]
			for k := 0; k < n; k++ {
				dik := dis[i][k]
				if (i-w) <= j && j < k && k < i && dij < dik {
					sum++
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			dij := dis[i][j]
			for k := 0; k < n; k++ {
				dik := dis[i][k]
				if i < j && j < k && k <= i+w && dij > dik {
					sum++
				}
			}
		}
	}
	return sum
}

// RGARLoss returns the relative generalized anti-Robinson loss function for a distance matrix RGAR(w)  (Wu 2010: 773) .
func RGARLoss(dis Matrix64, p IntVector, w int) float64 {
	if !dis.IsSymmetric() {
		panic("distance matrix not symmetric")
	}
	n := p.Len()
	if dis.Rows() != n {
		panic("dimensions not equal")
	}

	gar := GARLoss(dis, p, w)
	sum := 0.0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if (i-w) <= j && j < k && k < i {
					sum++
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if i < j && j < k && k <= i+w {
					sum++
				}
			}
		}
	}
	return gar / sum
}
