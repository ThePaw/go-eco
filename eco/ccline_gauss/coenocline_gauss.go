// Coenocline modeller: Gaussian SRF

package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"sort"
	. "code.google.com/p/go-eco/eco/aux"
	. "code.google.com/p/probab/dst"
)

// Generate sampling points along the gradient
func generate_points(k, spacing int) (arr []float64) {
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
			arr[i] = arr[i-1] + ExponentialNext(float64(λ))
		}
		// scale to [0..1]
		for i := 1; i < k; i++ {
			arr[i] /= arr[k-1]
		}
	}
	return
}

// Gaussian response function 
func gaussSRF(opt, tol, max, x float64) (y float64) {
	// opt	 optimum = modus = position of max. abundance on the gradient (= mean for Gaussian)
	// tol	tolerance, range of nonzero values of abundance (fraction of gradient length)
	// max	amplitude, maximum abundance = abundance at modus
	// x		point on the gradient

	spanZ := 2 * 2.326348 // span between 2% (arbitrarily chosen) tails of Z distribution
	maxZ := 0.3989423     // value of Z at 0 (=mean=mode)
	x -= opt
	x *= spanZ / tol
	y = (max / maxZ) / math.Sqrt(2*math.Pi) * math.Exp(-x*x/2)
	return
}

// Random variable generator for abundance and tolerance
func rndFn(which int, a, b float64) func() (x float64) {
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
			x = a*rand.Float64() + b
		case Gaussian: // Gaussian
			x = NormalNext(a, b)
		case Beta: // Beta
			x = BetaNext(a, b)
		case ParetoSing: // single-parameter Pareto
			x = ParetoSingNext(a, b)
			/*
				case ParetoI: // Pareto I
					x = ParetoNext(a)
			*/
		case ParetoII: // Pareto II
			x = ParetoIINext(a, b)
			/*
				case ParetoIII: // Pareto III
					x = ParetoIIINext(a, b)
				case ParetoIV: // Pareto IV
					x = ParetoIVNext(a, b)
		case ParetoG: // Generalized Pareto
			x = ParetoGNext(a, b, c)
		case ParetoTap: // tapered Pareto
			x = ParetoTapNext(a, b, c)
				case Yule: // Yule
					x = YuleNext(a)
				case Planck: // Planck
					x = PlanckNext(a, b)
				case Zeta: // Zeta
					x = ZetaNext(a)
			*/
		}
		return
	}
}

// Coenocline modeller
func Coenocline(nSpec, nSamp, optModel, abuModel, tolModel, spacing int, aOpt, bOpt, aAbu, bAbu, aTol, bTol, abuMax, tolMax float64) (out *Matrix) {
	out = NewMatrix(nSamp, nSpec)
	points := generate_points(nSamp, spacing) // generate sampling points
	rngO := rndFn(optModel, aOpt, bOpt)           // optima distribution model
	rngA := rndFn(abuModel, aAbu, bAbu)           // abundance distribution model
	rngT := rndFn(tolModel, aTol, bTol)           // tolerance distribution model
	for j := 0; j < nSpec; j++ {
		opt := rngO()				// optimum (point on the gradient)
		t := tolMax * rngT()			// tolerance (range of acceptable gradient values)
		for t < 2/float64(nSamp) {		// if tolerance is too small, generate new value
			t = tolMax * rand.Float64()
		}

		max := abuMax * rngA() // species' max abundance
		for i := 0; i < nSamp; i++ {
			x := points[i]
			y := gaussSRF(opt, t, max, x)
			out.Set(i, j, y)
		}
	}
	return
}

func main() {
	help := flag.Bool("h", false, "Coenocline modeller\nUsage: coenocline  [-xysoat] [-ao -bo] [-aa -ba] [-at -bt] ")
	nSpec := flag.Int("x", 20, "number of species")
	nSamp := flag.Int("y", 30, "number of samples")
	sampModel := flag.Int("s", 0, "sampling model: 0 - regular, 1 - uniform, 2 - Poisson")
	optModel := flag.Int("o", 0, "optima model: 0 - flat, 1 - Gaussian, 2 - Beta...")
	abuModel := flag.Int("a", 0, "abundance model: 0 - flat, 1 - Gaussian, 2 - Beta...")
	tolModel := flag.Int("t", 0, "tolerance model: 0 - flat, 1 - Gaussian, 2 - Beta...")

aOpt := 1.0
bOpt := 0.0
aAbu :=  1.0
bAbu := 0.0
aTol :=  1.0
bTol :=  0.0
abuMax := 100.0
tolMax := 0.5

	flag.Parse()

	if *help {
		flag.PrintDefaults()
	} else {
		mtx := Coenocline(*nSpec, *nSamp, *sampModel, *optModel, *abuModel, *tolModel, aOpt, bOpt, aAbu, bAbu, aTol, bTol, abuMax, tolMax)
		for i := 0; i < *nSamp; i++ {
			for j := 0; j < *nSpec; j++ {
				fmt.Print(mtx.Get(i, j), ",")
			}
			fmt.Println()
		}
	}
}

