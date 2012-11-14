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
	. "code.google.com/p/probab/dst"
	"math"
	"math/rand"
	"sort"
)

// Generate sampling points along the gradient
func generate_points(k int, spacing byte) (arr []float64) {
	const offset = 0.2
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
		λ := 1 / k
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
		arr[i] = arr[i]*(1-2*offset) + offset // sampling starts at 'offset' and ends at '1-offset'
	}
	return
}

// Triangular response function of a taxon on a gradient. 
func triangSRF(opt, tol, max, exc, x float64) (y float64) {
	// opt	optimum = modus = position of max. abundance on the gradient
	// tol	tolerance, range of nonzero values of abundance (fraction of gradient length)
	// max	amplitude, maximum abundance = abundance at modus
	// exc	excentricity = left-right; is zero if symmetric, -1 or +1 if extremely asymmetric
	// x		point on the gradient

	// exc = left-right; tol = left-right; thus:
	right := (tol - exc) / 2 // segment above optimum
	left := exc + right      // segment below optimum
	lo := opt - left         // lower tolerance bound
	hi := opt + right        // lower tolerance bound

	if x <= lo || x >= hi {
		y = 0

	} else if x <= opt {
		a := max / left
		y = a * (x - lo)
	} else { // x > opt
		a := -max / right
		y = a * (x - opt)
	}
	return
}

// Gaussian response function 
func gaussSRF(opt, tol, max, exc, x float64) (y float64) {
	// opt	 optimum = modus = position of max. abundance on the gradient (= mean for Gaussian)
	// tol	tolerance, range of nonzero values of abundance (fraction of gradient length)
	// max	amplitude, maximum abundance = abundance at modus
	// exc	excentricity: not applicable, as Gaussian fn is symmetric; just for compatibility with other SRF's
	// x		point on the gradient

	exc = 0
	spanZ := 2 * 2.326348 // span between 2% (arbitrarily chosen) tails of Z distribution
	maxZ := 0.3989423     // value of Z at 0 (=mean=mode)
	x -= opt
	x *= spanZ / tol
	y = (max / maxZ) / math.Sqrt(2*math.Pi) * math.Exp(-x*x/2)
	return
}

// SRF to be used
func sRF(which byte) func() (x float64) {
	return func() (y float64) {
		const (
			Gaussian = iota
			Triangular
		)
		srf := gaussSRF //default
		switch which {
		case Gaussian:
			srf = gaussSRF
		case Triangular:
			srf = triangSRF
		}
		return 
	}
}

// Random variable generator for abundance and tolerance
func rndFn(which byte, a, b, c float64) func() (x float64) {
	return func() (x float64) {
		const (
			Flat = iota
			Gaussian
			Beta
			ParetoSing
			ParetoI
			ParetoII
			ParetoIII
			ParetoIV
			ParetoG
			ParetoTap
			Yule
			Planck
			Zeta
		)

		switch which {
		case Flat: // Flat
			x = rand.Float64()
		case Gaussian: // Gaussian
			x = NextNormal(a, b)
		case Beta: // Beta
			x = NextBeta(a, b)
		case ParetoSing: // single-parameter Pareto
			x = NextParetoSing(a, b)
/*
		case ParetoI: // Pareto I
			x = NextPareto(a)
*/
		case ParetoII: // Pareto II
			x = NextParetoII(a, b)
/*
		case ParetoIII: // Pareto III
			x = NextParetoIII(a, b)
		case ParetoIV: // Pareto IV
			x = NextParetoIV(a, b)
*/
		case ParetoG: // Generalized Pareto
			x = NextParetoG(a, b, c)
		case ParetoTap: // tapered Pareto
			x = NextParetoTap(a, b, c)  
/*
		case Yule: // Yule
			x = NextYule(a)
		case Planck: // Planck
			x = NextPlanck(a, b)
		case Zeta: // Zeta
			x = NextZeta(a)
*/
		}
		return
	}
}

// Coenocline modeller
func Coenocline(nSpec, nSamp int, srfModel, optModel, abuModel, tolModel, spacing  byte, aa, ba, at, bt, abumax, tmax, alphamax, gammamax float64) (out *Matrix) {
	var lo float64
	out = NewMatrix(nSamp, nSpec)
	points := generate_points(nSamp, spacing)	// generate sampling points
	rngO := rndFn(optModel, ao, bo)			// optima distribution model
	rngA := rndFn(abuModel, aa, ba)			// abundance distribution model
	rngT := rndFn(tolModel, at, bt)				// tolerance distribution model
	if srfModel < 0 || srfModel > 3 {
		panic("this SRF model is not defined")
	}
/*
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
*/

	return
}

