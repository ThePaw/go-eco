// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package ser

// Objective (loss and gain) functions for distance (dissimilarity) matrices. 

import (
	"math"
)

func f(x, y float64) float64 {
	if x < y {
		return 1
	}
	if x > y {
		return -1
	}
	return 0
}

func g(x, y float64) float64 {
	if x > y {
		return 1
	}
	return 0
}

// G1Gain returns gain of the permuted matrix according to Hubert, Arabie & Meulman 2001, Chapter 4; see Brusco 2002: 50, Eq. 6. (WRUG)
func G1Gain(dis Matrix64, p IntVector) float64 {
	if !dis.IsSymmetric() {
		panic("distance matrix not symmetric")
	}
	n := p.Len()
	if dis.Rows() != n {
		panic("bad permutation vector length")
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

// G2Gain returns gain of the permuted matrix according to Hubert, Arabie & Meulman 2001, Chapter 4; see Brusco 2002: 50, Eq. 7. (WRCUG)
func G2Gain(dis Matrix64, p IntVector) float64 {
	if !dis.IsSymmetric() {
		panic("distance matrix not symmetric")
	}
	n := p.Len()
	if dis.Rows() != n {
		panic("bad permutation vector length")
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

// G3Gain returns gain of the permuted matrix according to Hubert, Arabie & Meulman 2001, Chapter 4; see Brusco 2002: 50, Eq. 8. (WRWG)
func G3Gain(dis Matrix64, p IntVector) float64 {
	if !dis.IsSymmetric() {
		panic("distance matrix not symmetric")
	}
	n := p.Len()
	if dis.Rows() != n {
		panic("bad permutation vector length")
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

// G4Gain returns gain of the permuted matrix according to Hubert, Arabie & Meulman 2001, Chapter 4; see Brusco 2002: 50, Eq. 9. (WRCWG)  (? approx. -StrengLoss2)
func G4Gain(dis Matrix64, p IntVector) float64 {
	if !dis.IsSymmetric() {
		panic("distance matrix not symmetric")
	}
	n := p.Len()
	if dis.Rows() != n {
		panic("bad permutation vector length")
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

// HGain returns gain of the permuted matrix according to Szczotka 1972; see Brusco et al. 2008: 507, Eq. 7.
func HGain(dis Matrix64, p IntVector) float64 {
	if !dis.IsSymmetric() {
		panic("distance matrix not symmetric")
	}
	n := p.Len()
	if dis.Rows() != n {
		panic("bad permutation vector length")
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

// HNormGain returns gain of the permuted matrix according to Szczotka 1972; see Brusco et al. 2008: 507-508, Eq. 7.
// TO BE IMPLEMENTED

// InertiaGain returns the Inertia criterion (Caraux and Pinloche 2005).
func InertiaGain(dis Matrix64, p IntVector) float64 {
	if !dis.IsSymmetric() {
		panic("distance matrix not symmetric")
	}
	n := p.Len()
	if dis.Rows() != n {
		panic("bad permutation vector length")
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
		panic("bad permutation vector length")
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

// MooreStressDisLoss returns the Moore Stress criterion (Niermann 2005:42, Eq. 1, 2) for a distance matrix.
func MooreStressDisLoss(dis Matrix64, p IntVector) float64 {
	return MooreStressLoss(dis, p, p)
}

// VonNeumannStressDisLoss returns the Moore Stress criterion (Niermann 2005:42) for a distance matrix.
func VonNeumannStressDisLoss(dis Matrix64, p IntVector) float64 {
	return VonNeumannStressLoss(dis, p, p)
}


// RGARLoss returns the relative generalized anti-Robinson loss function for a distance matrix RGAR(w)  (Wu 2010: 773) .
func RGARLoss(dis Matrix64, p IntVector, w int) float64 {
	if !dis.IsSymmetric() {
		panic("distance matrix not symmetric")
	}
	n := p.Len()
	if dis.Rows() != n {
		panic("bad permutation vector length")
	}

	gar := GARLoss(dis, p, w)
	/*
		c:=0.0
		for j := 0; j < n; j++ {
			for k := j + 1; k < n; k++ {
				for i := k + 1; i < n; i++ {
					if i-w <= j && j < k && k < i {
						c ++
					}
				}
			}
		}
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				for k := j + 1; k < n; k++ {
					if i < j && j < k && k <= i+w {
						c ++
					}
				}
			}
		}
	*/
	return gar / (float64(n*w*(w-1)) - 2*float64(w)*float64(w*w-1)/3)

}

func GARLoss5(dis Matrix64, p IntVector) float64 {
	return GARLoss(dis, p, 5)
}

func GARLoss10(dis Matrix64, p IntVector) float64 {
	return GARLoss(dis, p, 10)
}

func GARLoss12(dis Matrix64, p IntVector) float64 {
	return GARLoss(dis, p, 12)
}

func GARLoss15(dis Matrix64, p IntVector) float64 {
	return GARLoss(dis, p, 15)
}

func GARLoss25(dis Matrix64, p IntVector) float64 {
	return GARLoss(dis, p, 25)
}

func GARLoss37(dis Matrix64, p IntVector) float64 {
	return GARLoss(dis, p, 37)
}

func GARLoss50(dis Matrix64, p IntVector) float64 {
	return GARLoss(dis, p, 50)
}

func GARLoss75(dis Matrix64, p IntVector) float64 {
	return GARLoss(dis, p, 75)
}

func GARLoss112(dis Matrix64, p IntVector) float64 {
	return GARLoss(dis, p, 112)
}

func GARLoss125(dis Matrix64, p IntVector) float64 {
	return GARLoss(dis, p, 125)
}

func GARLoss187(dis Matrix64, p IntVector) float64 {
	return GARLoss(dis, p, 187)
}



func RGARLoss5(dis Matrix64, p IntVector) float64 {
	return RGARLoss(dis, p, 5)
}

func RGARLoss10(dis Matrix64, p IntVector) float64 {
	return RGARLoss(dis, p, 10)
}

func RGARLoss12(dis Matrix64, p IntVector) float64 {
	return RGARLoss(dis, p, 12)
}

func RGARLoss15(dis Matrix64, p IntVector) float64 {
	return RGARLoss(dis, p, 15)
}

func RGARLoss25(dis Matrix64, p IntVector) float64 {
	return RGARLoss(dis, p, 25)
}

func RGARLoss37(dis Matrix64, p IntVector) float64 {
	return RGARLoss(dis, p, 37)
}

func RGARLoss50(dis Matrix64, p IntVector) float64 {
	return RGARLoss(dis, p, 50)
}


func RGARLoss75(dis Matrix64, p IntVector) float64 {
	return RGARLoss(dis, p, 75)
}

func RGARLoss112(dis Matrix64, p IntVector) float64 {
	return RGARLoss(dis, p, 112)
}

func RGARLoss125(dis Matrix64, p IntVector) float64 {
	return RGARLoss(dis, p, 125)
}

func RGARLoss187(dis Matrix64, p IntVector) float64 {
	return RGARLoss(dis, p, 187)
}


// HamiltonLoss returns the length of the shortest Hamiltonian path (openTSP).
func HamiltonLoss(dis Matrix64, p IntVector) float64 {
	if !dis.IsSymmetric() {
		panic("distance matrix not symmetric")
	}
	n := p.Len()
	if dis.Rows() != n {
		panic("bad permutation vector length")
	}

	sum := 0.0
	for i := 0; i < n-1; i++ {
		sum += dis[p[i]][p[i+1]]
	}
	return sum
}

// parabolaFit returns coefficients of the polynomial c1 + c2*x + c3*x*x fitted to the data vector, where abscissa is 0, 1, 2, ... , n-1
func parabolaFit(v Vector64) (c1, c2, c3 float64, err bool) {
	eps := 1e-6
	n := float64(v.Len())

	sumX := 0.0
	sumY := 0.0
	sumXY := 0.0
	sumX2 := 0.0
	sumX2Y := 0.0
	sumX3 := 0.0
	sumX4 := 0.0
	for i, y := range v {
		x := float64(i)
		sumX += x
		sumY += y
		sumXY += x * y
		sumX2 += x * x
		sumX2Y += x * x * y
		sumX3 += x * x * x
		sumX4 += x * x * x * x
	}

	// The normal equations
	// sumY = c1*n + c2*sumX + c3*sumX2
	// sumXY = c1*sumX + c2*sumX2  + c3*sumX3
	// sumX2Y = c1*sumX2 + c2*sumX3  + c3*sumX4

	a := n
	b := sumX
	c := sumX2
	d := sumX
	e := sumX2
	f := sumX3
	g := sumX2
	h := sumX3
	k := sumX4

	b0 := sumY
	b1 := sumXY
	b2 := sumX2Y

	// determinant of a 3x3 matrix can be computed by applying the rule of Sarrus as follows:
	det := a*(e*k-f*h) - b*(k*d-f*g) + c*(d*h-e*g)
	if det > -eps && det < eps { // determinant is zero, matrix not invertible
		c1 = 0.0
		c2 = 0.0
		c3 = 0.0
		err = true
		return
	}

	// inverse 3x3 matrix
	// http://en.wikipedia.org/wiki/Matrix_inverse
	i00 := (e*k - f*h) / det
	i10 := (f*g - d*k) / det
	i20 := (d*h - e*g) / det
	i01 := (c*h - b*k) / det
	i11 := (a*k - c*g) / det
	i21 := (b*g - a*h) / det
	i02 := (b*f - c*e) / det
	i12 := (c*d - a*f) / det
	i22 := (a*e - b*d) / det

	c1 = i00*b0 + i01*b1 + i02*b2
	c2 = i10*b0 + i11*b1 + i12*b2
	c3 = i20*b0 + i21*b1 + i22*b2
	err = false
	return
}

// ParabolaLoss returns the sum of squared residues of fitted parabola.
func ParabolaLoss(sim Matrix64, p IntVector) float64 {
	if !sim.IsSymmetric() {
		panic("similarity matrix not symmetric")
	}
	n := p.Len()
	if sim.Rows() != n {
		panic("bad permutation vector length")
	}

	loss := 0.0
	rows := p.Len()
	cols := p.Len()

	v := NewVector64(cols)
	for i := 0; i < rows; i++ {
		// unload the permuted row to a vector
		for j := 0; j < cols; j++ {
			v[j] = sim[p[i]][p[j]]
		}

		// find position of the maximum
		mx := -inf
		pos := 0
		for j := 0; j < cols; j++ {
			if v[j] > mx {
				mx = v[j]
				pos = j
			}
		}

		if pos < 3 {
			// fit parabola to upper part
			m := cols - pos
			//			m:= cols-pos-1 // do not include pos
			w := NewVector64(m)
			for j := 0; j < m; j++ {

				// fmt.Println(m, j, pos)
				w[j] = v[j+pos]
				//w[j] = v[j+pos+1] // do not include pos
			}
			c1, c2, c3, err := parabolaFit(w)
			if !err {
				for j := 0; j < m; j++ {
					k := float64(j)
					x := w[j]
					y := c1 + c2*k + c3*k*k
					z := x - y
					//fmt.Println(x, yyy, z)
					z *= z
					loss += z
				}
			}
		} else if pos > n-4 {
			// fit parabola to lower part
			m := pos + 1
			//	m:= pos // do not include pos
			w := NewVector64(m)
			for j := 0; j < m; j++ {
				w[j] = v[j]
				//			w[j] = v[j+1] // do not include pos
			}
			c1, c2, c3, err := parabolaFit(w)
			if !err {
				for j := 0; j < m; j++ {
					k := float64(j)
					x := w[j]
					y := c1 + c2*k + c3*k*k
					z := x - y
					//fmt.Println(x, yyy, z)
					z *= z
					loss += z
				}
			}
		} else {

			// fit parabola to both: lower part
			m := pos + 1
			//m := pos // do not include pos

			w := NewVector64(m)
			for j := 0; j < m; j++ {
				w[j] = v[j]
				//w[j] = v[j+1] // do not include pos
			}

			//w.Print()
			c1, c2, c3, err := parabolaFit(w)
			if !err {
				for j := 0; j < m; j++ {
					k := float64(j)
					x := w[j]
					y := c1 + c2*k + c3*k*k
					z := x - y
					z *= z
					loss += z
				}
			}

			// fit parabola to upper part
			//fmt.Println("---")
			m = cols - pos
			//m = cols - pos - 1 // do not include pos
			w = NewVector64(m)
			for j := 0; j < m; j++ {

				// fmt.Println(m, j, pos)
				w[j] = v[j+pos]
				//w[j] = v[j+pos+1] // do not include pos
			}
			// w.Print()
			c1, c2, c3, err = parabolaFit(w)
			if !err {
				for j := 0; j < m; j++ {
					k := float64(j)
					x := w[j]
					y := c1 + c2*k + c3*k*k
					z := x - y
					//fmt.Println(x, yyy, z)
					z *= z
					loss += z
				}
			}

		}

	}
	return loss
}

/*
// AREventsViolationLoss returns gain of the permuted matrix according to Kostopoulos & Goulermas
func AREventsViolationLoss(dis Matrix64, p IntVector) float64 {
	if !dis.IsSymmetric() {
		panic("distance matrix not symmetric")
	}
	n := p.Len()
	if dis.Rows() != n {
		panic("bad permutation vector length")
	}

	c := 0.0
	for i := 0; i < n; i++ {
		for j := i + 2; j < n; j++ {
			for k := i + 1; k < j; k++ {
				x := dis[p[i]][p[k]]
				y := dis[p[i]][p[j]]
				c += g(x, y)
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := i + 2; j < n; j++ {
			for k := i + 1; k < j; k++ {
				x := dis[p[k]][p[j]]
				y := dis[p[i]][p[j]]
				c += g(x, y)
			}
		}
	}
	return c
}
*/

// AREventsViolationLoss returns loss of the permuted matrix according to Kostopoulos & Goulermas
func AREventsViolationLoss(dis Matrix64, p IntVector) float64 {
	if !dis.IsSymmetric() {
		panic("distance matrix not symmetric")
	}
	n := p.Len()
	if dis.Rows() != n {
		panic("bad permutation vector length")
	}

	c := 0.0
	for i := 0; i < n-2; i++ {
		for j := i + 2; j < n; j++ {
			for k := i + 1; k < j; k++ {
				x := dis[p[i]][p[k]]
				y := dis[p[i]][p[j]]
				c += g(x, y)
			}
		}
	}
	for i := 0; i < n-2; i++ {
		for j := i + 2; j < n; j++ {
			for k := i + 1; k < j; k++ {
				x := dis[p[k]][p[j]]
				y := dis[p[i]][p[j]]
				c += g(x, y)
			}
		}
	}
	return c
}

// WeightedAREventsViolationLoss returns gain of the permuted matrix according to Kostopoulos & Goulermas
func WeightedAREventsViolationLoss(dis Matrix64, p IntVector) float64 {
	if !dis.IsSymmetric() {
		panic("distance matrix not symmetric")
	}
	n := p.Len()
	if dis.Rows() != n {
		panic("bad permutation vector length")
	}

	c := 0.0
	for i := 0; i < n-2; i++ {
		for j := i + 2; j < n; j++ {
			for k := i + 1; k < j; k++ {
				x := dis[p[i]][p[k]]
				y := dis[p[i]][p[j]]
				d := math.Abs(x - y)
				c += d * g(x, y)
			}
		}
	}
	for i := 0; i < n-2; i++ {
		for j := i + 2; j < n; j++ {
			for k := i + 1; k < j; k++ {
				x := dis[p[k]][p[j]]
				y := dis[p[i]][p[j]]
				d := math.Abs(x - y)
				c += d * g(x, y)
			}
		}
	}
	return c
}

// DoublyWeightedAREventsViolationLoss returns loss of the permuted matrix according to Kostopoulos & Goulermas
func DoublyWeightedAREventsViolationLoss(dis Matrix64, p IntVector) float64 {
	if !dis.IsSymmetric() {
		panic("distance matrix not symmetric")
	}
	n := p.Len()
	if dis.Rows() != n {
		panic("bad permutation vector length")
	}

	c := 0.0
	for i := 0; i < n-2; i++ {
		for j := i + 2; j < n; j++ {
			ij := math.Abs(float64(i - j))
			for k := i + 1; k < j; k++ {
				x := dis[p[i]][p[k]]
				y := dis[p[i]][p[j]]
				d := math.Abs(x - y)
				c += ij * d * g(x, y)
			}
		}
	}
	for i := 0; i < n-2; i++ {
		for j := i + 2; j < n; j++ {
			ij := math.Abs(float64(i - j))
			for k := i + 1; k < j; k++ {
				x := dis[p[k]][p[j]]
				y := dis[p[i]][p[j]]
				d := math.Abs(x - y)
				c += ij * d * g(x, y)
			}
		}
	}
	return c
}

// MEffGainDis returns the measure of Effectiveness (McCormick 1972).
func MEffGainDis(a Matrix64, p IntVector) float64 {
	var x0, x1, x2, x3, x4 float64
	rows := a.Rows()

	if !(p.Len() == rows) {
		panic("bad permutation vector length")
	}
	gain := 0.0
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			x0 = a[p[i]][p[j]]
			if j-1 < 0 {
				x1 = 0
			} else {
				x1 = a[p[i]][p[j-1]]
			}
			if j+1 > rows-1 {
				x2 = 0
			} else {
				x2 = a[p[i]][p[j+1]]
			}
			if i-1 < 0 {
				x3 = 0
			} else {
				x3 = a[p[i-1]][p[j]]
			}

			if i+1 > rows-1 {
				x4 = 0
			} else {

				x4 = a[p[i+1]][p[j]]
			}
			gain += x0 * (x1 + x2 + x3 + x4)
		}
	}
	return gain / 2
}

/*
// GARLoss returns gain of the permuted matrix according to Kostopoulos & Goulermas
func GARLoss(dis Matrix64, p IntVector, w int) float64 {
	if !dis.IsSymmetric() {
		panic("distance matrix not symmetric")
	}
	n := p.Len()
	if dis.Rows() != n {
		panic("bad permutation vector length")
	}

	c := 0.0
	for j := 0; j < n; j++ {
		for k := j + 1; k < n; k++ {
			for i := k + 1; i < n; i++ {
				if i-w <= j && j < k && k < i {
					x := dis[p[i]][p[k]]
					y := dis[p[i]][p[j]]
					c += g(x, y)
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if i < j && j < k && k <= i+w {
					x := dis[p[k]][p[j]]
					y := dis[p[i]][p[j]]
					c += g(x, y)
				}
			}
		}
	}
	return c
}

*/
// GARLoss returns the generalized anti-Robinson loss function for a distance matrix GAR(w) (Wu 2010: 773) .
func GARLoss(dis Matrix64, p IntVector, w int) float64 {
	if !dis.IsSymmetric() {
		panic("distance matrix not symmetric")
	}
	n := p.Len()
	if dis.Rows() != n {
		panic("bad permutation vector length")
	}

	sum := 0.0
	for j := 0; j < n; j++ {
		for k := 0; k < j; k++ {
			for i := j - w; i < k; i++ {
				if i >= 0 {
					dik := dis[p[i]][p[k]]
					dij := dis[p[i]][p[j]]
					sum += g(dik, dij)
				}
			}
		}
	}
	for j := 0; j < n; j++ {
		for k := 0; k < j; k++ {
			for i := j - w; i < k; i++ {
				if i >= 0 {
					dkj := dis[p[k]][p[j]]
					dij := dis[p[i]][p[j]]
					sum += g(dkj, dij)
				}
			}
		}
	}
	return sum
}
