// Coenocline modeller

/* To do:
Poisson / clustered optima placing 
implement HOF model
rewrite Beta model to generate 'u', not 'lo'
enable Pareto and Yule models when ready
*/


package main

import (
	"flag"
	"sort"
	"fmt"
	. "go-eco.googlecode.com/hg/eco"
	. "gostat.googlecode.com/hg/stat/prob"
	"math"
	"math/rand"
)

// generate sampling points along the gradient
func generate_points(k int, randomSpacing bool) (arr []float64) {
	arr = make([]float64, k)
	if ! randomSpacing { // regular spacing, default
		for i := 0; i < k; i++ {
			arr[i] = float64(i) / float64(k-1)
		}

	} else { // random spacing
		for i := 0; i < k; i++ {
			arr[i] = rand.Float64()

		}
		sort.Float64s(arr) // sort in increasing order
	}
// now scale it to allow overlap over the gradient's ends
		for i := 0; i < k; i++ {
			arr[i] = arr[i]*0.6 + 0.2
		}
	return
}

// `roof' (triangular) response function of a taxon on a gradient
func roof(x, a, e, u, r float64) (y float64) {
	/*
		x	 point on the gradient
		a	 amplitude, maximum abundance
		e	 excentricity = left/range
		u	 mean, position of max. abundance on the gradient
		r	 range of nonzero values of abundance
	*/

	if x < u {
		y = x*(a/(e*r)) + a - u*(a/(e*r))
	} else if x < (u + r - r*e) {
		y = x*(-a)/(r-e*r) + a - u*(-a)/(r-e*r)
	} else {
		y = 0
	}
	return
}

// Gaussian response function 
func gauss(x, u, t, c float64) (y float64) {
	// Z = 1/math.Sqrt(2*math.Pi)*math.Exp(-x*x/2)
	// y = c*math.Exp(-0.5(x-u)^2/t^2)
	// u	 optimum
	// t	 tolerance
	// c	 maximum
	// x	 point at which the function is evaluated
	//return c * math.Exp(-0.5*(x-u)*(x-u)/(t*t))
	tails:=2*2.326348	// 1% tails of Z
	maxZ := 0.3989423	// value of Z at 0 (=mean=mode)
	x -= u
	x *= tails/t
	y=(c/maxZ)/math.Sqrt(2*math.Pi)*math.Exp(-x*x/2)
	return
}

// Beta response function 
func beta(k, a, b, alpha, gamma, x float64) (y float64) {
	// Return zero if x is not in (a,b)
	if x <= a || x >= b {
		y = 0
	} else {
		// Otherwise evaluate the beta-function at x
		t2 := math.Pow(x-a, alpha)
		t3 := math.Pow(b-x, gamma)
		y = k * t2 * t3
	}
	return
}

// Solve k from the maximum height of the response function
// thanks to Jari Oksanen, I think
func ksol(a, b, alpha, gamma, height float64) (k float64) {
	t1 := b - a
	t4 := t1 / (alpha + gamma)
	t6 := math.Pow(alpha*t4, alpha)
	t11 := math.Pow(gamma*t4, gamma)
	return height / t6 / t11
}

// For Beta-Binomial sampling model: Estimates the parameters a,b of beta distribution from expected proportion (pi), binomial denominator (m), and shape parameter (tau2). Solution (hopefully correct) of Exercise 4.17 of McCullagh & Nelder 1989, helped by Moore, Appl Stat 36, 8-14; 1987.
// thanks to Jari Oksanen, I think
func betapara(pi, m, tau2 float64) (a, b float64) {
	t1 := tau2 * m
	t2 := t1 - m - tau2 + 1
	t3 := 1 / (1 + t1 - tau2)
	t4 := t2 * t3
	a = -t4 * pi
	b = t4 * (pi - 1)
	return
}

func Coenocline(model, nSpec, nSamp, abuModel int, aa, ba float64, tModel int, at, bt float64, randomSpacing bool, abumax, tmax, alphamax, gammamax float64) (out *Matrix) {
	var lo float64
	out = NewMatrix(nSamp, nSpec)
	points := generate_points(nSamp, randomSpacing)
	rngA := rndFn(abuModel, aa, ba)
	rngT := rndFn(tModel, at, bt)
	if model < 0 || model > 2 {
		model = 0
	}
	switch {
	case model == 0: // Gaussian
		for j := 0; j < nSpec; j++ {
			u := rand.Float64()          // optimum (point on the gradient)
			t := tmax * rngT()   // tolerance (range of acceptable gradient values)
			for t < 2/float64(nSamp) { // if tolerance is too small, generate new value
				t = tmax * rand.Float64()
			}

			c := abumax * rngA()	// species' max abundance
			for i := 0; i < nSamp; i++ {
				x := points[i]
				y := gauss(x, u, t, c)
				out.Set(i, j, y)
			}
		}
	case model == 1: // Beta		
		for j := 0; j < nSpec; j++ {
			t := tmax * rngT()   // tolerance (range of acceptable gradient values)
			for t < 2/float64(nSamp) { // if tolerance is too small, generate new value
				t = tmax * rand.Float64()
			}
			lo := rand.Float64()   // lower end
			hi := lo + t   // upper end

			c := abumax * rngA()	// species' max abundance
			alpha := alphamax * rand.Float64() // shape parameters alpha, gamma
			gamma := gammamax * rand.Float64()
			k := ksol(lo, hi, alpha, gamma, c)
			for i := 0; i < nSamp; i++ {
				x := points[i]
				y := beta(k, lo, hi, alpha, gamma, x)
				out.Set(i, j, y)
			}
		}
	case model == 2: // Triangular
		for j := 0; j < nSpec; j++ {
			t := tmax * rngT()   // tolerance (range of acceptable gradient values)
			for t < 2/float64(nSamp) { // if tolerance is too small, generate new value
				t = tmax * rand.Float64()
			}

			c := abumax * rngA()	// species' max abundance
			e := rand.Float64()          // excentricity
			lo = rand.Float64()   // lower end

			u := lo + e*t                // optimum
			for i := 0; i < nSamp; i++ {
				x := points[i]
				y := roof(x, c, e, u, t)
				if y < 0 {
					y = 0
				}
				out.Set(i, j, y)
			}
		}
	}
//	fmt.Println("MODEL: ", model)

	return
}

// Random variable generator
func rndFn(which int, a, b float64) func() (x float64) {
	return func() (x float64) {
		switch {
		case which == 0:	// flat
			x = rand.Float64()
		case which == 1:	// Gaussian
			x = NextNormal(a, b)
		case which == 2:	// Beta
			x = NextBeta(a, b)
/* to be implemented (in stat/prob):
		case which == 2:	// Pareto
			x = NextPareto(a, b)
		case which == 3:	// Yule
			x = NextYule(a, b)
*/
		}
		return
	}
}


// Main
func main() {
	help := flag.Bool("h", false, "Coenocline modeller\nUsage: coenocline -m [0|1|2] [-rsxya]")
	model := flag.Int("m", 0, "Response function: 0 = Gaussian, 1 = Beta, 2 = Triangular")
	randomSpacing := flag.Bool("r", false, "true if spacing of samples along the gradient is random")
	nSpec := flag.Int("x", 20, "number of species")
	nSamp := flag.Int("y", 30, "number of samples")
	abuModel := flag.Int("abuModel", 0, "model of distribution of abundances")
	aa := flag.Float64("aa", 0, "scale param of distribution of abundances")
	ba := flag.Float64("ba", 0, "shape param of distribution of abundances")
	tolModel := flag.Int("tolModel", 0, "model of distribution of tolerances")
	at := flag.Float64("at", 0, "scale param of distribution of tolerances")
	bt := flag.Float64("bt", 0, "shape param of distribution of tolerances")
	//	seed := flag.Int64("s", 0, "random number seed")
	abumax := flag.Float64("a", 100.0, "maximum abundance")
	tmax := flag.Float64("tmax", 0.3, "maximum t")
	alphamax := flag.Float64("alpha", 5.5, "maximum alpha")
	gammamax := flag.Float64("gamma", 5.5, "maximum gamma")

	flag.Parse()

	if *help {
		flag.PrintDefaults()
	} else {
		mtx := Coenocline(*model, *nSpec, *nSamp, *abuModel, *aa, *ba, *tolModel, *at, *bt, *randomSpacing, *abumax, *tmax, *alphamax, *gammamax)
		for i := 0; i < *nSamp; i++ {
			for j := 0; j < *nSpec; j++ {
				fmt.Print(mtx.Get(i, j), ",")
			}
			fmt.Println()
		}
	}
}



