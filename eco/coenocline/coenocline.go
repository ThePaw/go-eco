// Coenocline modeller

/* To do:
implement Niche_apportionment_models http://en.wikipedia.org/wiki/Niche_apportionment_models for optima placing 
rewrite Beta model to generate 'u', not 'lo'
enable Pareto Zipf Planck Digamma Trigamma model when ready
write manpage
write publication ;-)
*/


package main

import (
	"sort"
	. "gostat.googlecode.com/hg/stat/prob"
	"math"
	"math/rand"
)

// generate sampling points along the gradient
func generate_points(k int, spacing byte) (arr []float64) {
	const offset=0.2
	arr = make([]float64, k)
	switch spacing { 
	default: // regular spacing
		for i := 0; i < k; i++ {
			arr[i] = float64(i) / float64(k-1)
		}

	case 1: // uniform random spacing
		for i := 0; i < k; i++ {
			arr[i] = rand.Float64()

		}
		sort.Float64s(arr) // sort in increasing order
	case 2: // exponential random spacing (spacing at points of a Poisson process)
		λ := 1/k
		arr[0] = 0
		for i := 1; i < k; i++ {
			arr[i] = arr[i-1] + NextExp(float64(λ))
		}
		// scale to [0..1]
		for i := 1; i < k; i++ {
			arr[i] /= arr[k-1]
		}
	}
	// now shrink it to be in some distance from the gradient's ends [0..1]
	for i := 0; i < k; i++ {
		arr[i] = arr[i]*(1-2*offset) + offset	// sampling starts at 'offset' and ends at '1-offset'
	}
	return
}
/*
// Triangular response function of a taxon on a gradient. 
func triang(x, max, exc, opt, tol float64) (y float64) {
		// x	 point on the gradient
		// max	 amplitude, maximum abundance
		// exc	 excentricity = left/range   <--- not very nice definition
		// opt	 mean, position of max. abundance on the gradient
		// tol	 range of nonzero values of abundance

	if x < opt {
		y = x*(max/(exc*tol)) + max - opt*(max/(exc*tol))
	} else if x < (opt + tol - tol*exc) {
		y = x*(-max)/(tol-exc*tol) + max - opt*(-max)/(tol-exc*tol)
	} else {
		y = 0
	}
	return
}
*/

// Triangular response function of a taxon on a gradient. 
func triangSRF(opt, tol, max, exc, x float64) (y float64) {
		// x	 point on the gradient
		// max	 amplitude, maximum abundance
		// exc	 excentricity = left-right; is zero if symmetric, -1 or +1 if extremely asymmetric
		// opt	 modus, position of max. abundance on the gradient
		// tol	 tolerance, range of nonzero values of abundance

		// exc = left-right; tol = left-right; thus:
	right := (tol-exc)/2	// segment above optimum
	left := exc+right	// segment below optimum
	lo := opt - left		// lower tolerance bound
	hi := opt + right	// lower tolerance bound

	if x <= lo || x >= hi {
		y = 0

	} else if x <= opt {
		a := max/left
		y = a*(x-lo)
	} else {	// x > opt
		a := -max/right
		y = a*(x-opt)
	}
	return
}

// Gaussian response function 
func gaussSRF(opt, tol, max, x float64) (y float64) {
	// x	 point at which the function is evaluated
	// opt	 optimum
	// tol	 tolerance (fraction of gradient length)
	// max	 maximum abundance = modus = mean
	spanZ:=2*2.326348	// span between 2% (arbitrarily chosen) tails of Z distribution
	maxZ := 0.3989423	// value of Z at 0 (=mean=mode)
	x -= opt
	x *= spanZ/tol
	y=(max/maxZ)/math.Sqrt(2*math.Pi)*math.Exp(-x*x/2)
	return
}

// Beta response function 
// Austin, M.P., 1976. On non-linear species responses models in ordination. Vegetatio 33, 33-41. DOI: 10.1007/BF00055297
// Austin, M.P., Gaywood, M.J., 1994. Current problems of environmental gradients and species response curves in relation to continuum theory. J. Veg. Sci. 5, 473-482. DOI: 10.2307/3235973
// This is NOT the Beta PDF !
// thanks to Jari Oksanen, betasimu.c

// func beta(max, lo, hi, α, γ, x float64) (y float64) {
func betaSRF(opt, tol, max, α, γ, x float64) (y float64) {
	// opt is where first derivative is zero
	// solve lo, hi
/////gnuplot> f(x)=k*(x-l)**a * (h-x)**g

	// Return zero if x is not in (lo,hi)
	if x <= lo || x >= hi {
		y = 0
	} else {
		// Otherwise evaluate the beta-function at x
		k := kSolve(tol, α, γ, max)
		t2 := math.Pow(x-lo, α)
		t3 := math.Pow(hi-x, γ)
		y = k * t2 * t3
	}
	return
}


// Solve k from the maximum height of the response function
// thanks to Jari Oksanen, betasimu.c
func kSolve(tol, max, α, γ float64) (k float64) {
	t4 := tol / (α + γ)
	t6 := math.Pow(α*t4, α)
	t11 := math.Pow(γ*t4, γ)
	return max / t6 / t11
}

// HOF response function 
// Huisman, J., Olff, H. & Fresco, L.F.M. (1993) A hierarchical set of models for species response analysis. Journal of Vegetation Science, 4, 37-46. 
func hof(a, b, c, d, m, x float64, which byte) (y float64) {
	switch which {
	case 1: // model I
		y = m/(1+math.Exp(a))
	case 2: // model II
		y = m/(1+math.Exp(a+b*x))
	case 3: // model III
		y = m/((1+math.Exp(a+b*x)) * (1+math.Exp(c)))
	case 4: // model IV
		y = m/((1+math.Exp(a+b*x)) * (1+math.Exp(c-b*x)))
	case 5: // model IV
		y = m/((1+math.Exp(a+b*x)) * (1+math.Exp(c+d*x)))
	}
	return
}
/*

// SRF model
func srfModel(which byte) (y float64) {
	switch which {
	case 0:	// Gaussian
		srf := 



}

// Coenocline modeller
func Coenocline(nSpec, nSamp int, srfModel, optModel, abuModel, tolModel, spacing  byte, aa, ba, at, bt, abumax, tmax, alphamax, gammamax float64) (out *Matrix) {
	var lo float64
	out = NewMatrix(nSamp, nSpec)
	points := generate_points(nSamp, spacing)	// generate sampling points
	rngO := ....			// optima distribution model
	rngA := rndFn(abuModel, aa, ba)			// abundance distribution model
	rngT := rndFn(tolModel, at, bt)			// tolerance distribution model
	if srfModel < 0 || srfModel > 3 {
		panic("this SRF model is not defined")
	}
	switch {
	case srfModel == 0: // Gaussian
		for j := 0; j < nSpec; j++ {
			opt := rngO()			// optimum (point on the gradient)
			t := tmax * rngT()		// tolerance (range of acceptable gradient values)
			for t < 2/float64(nSamp) {	// if tolerance is too small, generate new value
				t = tmax * rand.Float64()
			}

			max := abumax * rngA()		// species' max abundance
			for i := 0; i < nSamp; i++ {
				x := points[i]
				y := gauss(x, opt, t, max)
				out.Set(i, j, y)
			}
		}
	case srfModel == 1: // Beta		
		for j := 0; j < nSpec; j++ {
			t := tmax * rngT()   // tolerance (range of acceptable gradient values)
			for t < 2/float64(nSamp) { // if tolerance is too small, generate new value
				t = tmax * rand.Float64()
			}
			lo := rand.Float64()   // lower end
			hi := lo + t   // upper end

			max := abumax * rngA()	// species' max abundance
			α := alphamax * rand.Float64() // shape parameters α, γ
			γ := gammamax * rand.Float64()
			for i := 0; i < nSamp; i++ {
				x := points[i]
				y := beta(max, lo, hi, α, γ, x)
				out.Set(i, j, y)
			}
		}
	case srfModel == 2: // Triangular
		for j := 0; j < nSpec; j++ {
			t := tmax * rngT()   // tolerance (range of acceptable gradient values)
			for t < 2/float64(nSamp) { // if tolerance is too small, generate new value
				t = tmax * rand.Float64()
			}

			max := abumax * rngA()	// species' max abundance
			e := rand.Float64()          // excentricity
			lo = rand.Float64()   // lower end

			opt := lo + e*t                // optimum
			for i := 0; i < nSamp; i++ {
				x := points[i]
				y := triang(x, max, e, opt, t)
				if y < 0 {
					y = 0
				}
				out.Set(i, j, y)
			}
		}
	case srfModel == 3: // HOF
		for j := 0; j < nSpec; j++ {
			t := tmax * rngT()   // tolerance (range of acceptable gradient values)
			for t < 2/float64(nSamp) { // if tolerance is too small, generate new value
				t = tmax * rand.Float64()
			}

/////  implement a, b, c, d, m, which generation

			for i := 0; i < nSamp; i++ {
				x := points[i]
				y := hof(a, b, c, d, m, x, which)
				if y < 0 {
					y = 0
				}
				out.Set(i, j, y)
			}
		}
	}
//	fmt.Println("MODEL: ", srfModel)

	return
}

*/
// Random variable generator for abundance and tolerance
func rndFn(which byte, a, b, c float64) func() (x float64) {
	return func() (x float64) {
		switch {
		case which == 0:	// flat
			x = rand.Float64()
		case which == 1:	// Gaussian
			x = NextNormal(a, b)
		case which == 2:	// Beta
			x = NextBeta(a, b)
		case which == 3:	// single-parameter Pareto
			x = NextParetoSing(a, b)
		case which == 3:	// Pareto I
			x = NextPareto(a)
		case which == 3:	// Pareto II
			x = NextParetoII(a, b)
		case which == 3:	// Pareto III
			x = NextParetoIII(a, b)
		case which == 3:	// Pareto IV
			x = NextParetoIV(a, b)
		case which == 4:	// Generalized Pareto
			x = NextParetoG(a, b, c)
		case which == 3:	// tapered Pareto
			x = NextParetoTap(a, b)
		case which == 5:	// Yule
			x = NextYule(a)
		case which == 6:	// Planck
			x = NextPlanck(a, b)
		case which == 7:	// Zeta
			x = NextZeta(a)
		}
		return
	}
}



