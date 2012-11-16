// Coenocline modeller: Gaussian SRF

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sort"
	. "code.google.com/p/go-eco/eco/aux"
	. "code.google.com/p/probab/dst"
)

// Generate sampling points along the gradient
func generate_points(k, spacing int) (arr []float64) {
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
		λ := 1 / k // mean interval length
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
func gaussSRF(μ, σ, popSize, x float64) float64 {
	// μ	 	optimum = modus = position of max. abundance on the gradient (= mean for Gaussian)
	// σ		measure of tolerance, range of nonzero values of abundance (fraction of gradient length)
	// popSize	total size of population, along the whole gradient
	// x		point on the gradient

	return popSize * NormalPDFAt(μ, σ, x)
}

// Random variable generator for optima, abundances and tolerances
func rndFn(which int, μ, σ float64) func() (x float64) {
	return func() (x float64) {
		const (
			Flat = iota
			Gaussian
			Beta
			ParetoI
			ParetoII
			ParetoIII
			ParetoIV
			ParetoG
			ParetoTap
			ParetoSing
			Yule
			Planck
			Zeta
		)

		switch which {
		case Flat: // Uniform distribution
			a, b := UniformReparamMeanStd(μ, σ)
			x = UniformNext(a, b)
		case Gaussian: // Gaussian distribution
			x = NormalNext(μ, σ)
		case Beta: // Beta
			a, b := BetaReparamMeanStd(μ, σ)
			x = BetaNext(a, b)
		/*
		case ParetoI: // Pareto I
			x = ParetoNext(a, b)
		case ParetoII: // Pareto II distribution
			x = ParetoIINext(a, b)
		case ParetoIII: // Pareto III distribution
			x = ParetoIIINext(a, b)
		case ParetoIV: // Pareto IV distribution
			x = ParetoIVNext(a, b)
		case ParetoG: // Generalized Pareto distribution
			x = ParetoGNext(a, b, c)
		case ParetoTap: // tapered Pareto distribution
			x = ParetoTapNext(a, b, c)
		case ParetoSing: // single-parameter Pareto distribution
			x = ParetoSingNext(a, b)
		case Yule: // Yule distribution
			x = YuleNext(a)
		case Planck: // Planck distribution
			x = PlanckNext(a, b)
		case Zeta: // Zeta distribution
			x = ZetaNext(a)
		*/
		}
		return
	}
}

// Coenocline modeller 
func Coenocline(nSpec, nSamp, sampModel, optModel, abuModel, tolModel int, μOpt, σOpt, μAbu, σAbu, μTol, σTol float64) (out *Matrix) {
	out = NewMatrix(nSamp, nSpec)
	points := generate_points(nSamp, sampModel) // generate sampling points
	rngO := rndFn(optModel, μOpt, σOpt)       // optima distribution model
	rngA := rndFn(abuModel, μAbu, σAbu)       // abundance distribution model
	rngT := rndFn(tolModel, μTol, σTol)       // tolerance distribution model
	for j := 0; j < nSpec; j++ {              // for every species: 
		μ := rngO()                // generate optimum (point on the gradient)
		σ := rngT()                // generate tolerance (range of acceptable gradient values)
		for σ < 1/float64(nSamp) { // if tolerance is too small, generate new value
			σ = rngT()
		}

		popSize := rngA() // generate species' population size (abundance)
		for popSize <= 0.02*μAbu { // if population size is too small, generate new value
			popSize = rngA()
		}
		for i := 0; i < nSamp; i++ {
			x := points[i]
			y := gaussSRF(μ, σ, popSize, x)
			out.Set(i, j, y)
		}
	}
	return
}

// Main function 
func main() {
	help := flag.Bool("h", false, "Coenocline modeller\nUsage: ccline_gauss  [-nmso[om[os]]a[am[as]]t[tm[ts]]z]")
	nSpec := flag.Int("n", 20, "number of species")
	nSamp := flag.Int("m", 50, "number of samples")
	sampModel := flag.Int("s", 0, "sampling model: 0 - regular, 1 - uniform, 2 - Poisson")
	optModel := flag.Int("o", 0, "optima model: 0 - flat, 1 - Gaussian, 2 - Beta...")
	μOpt := flag.Float64("om", 0.5, "optimum mean")
	σOpt := flag.Float64("os", 0.5, "optimum standard deviation")
	abuModel := flag.Int("a", 0, "abundance model: 0 - flat, 1 - Gaussian, 2 - Beta...")
	μAbu := flag.Float64("am", 1.0, "mean population size")
	σAbu := flag.Float64("as", 0.03, "standard deviation of population size")
	tolModel := flag.Int("t", 1, "tolerance model: 0 - flat, 1 - Gaussian, 2 - Beta...")
	μTol := flag.Float64("tm", 0.003, "mean tolerance")
	σTol := flag.Float64("ts", 0.2, "standard deviation of tolerance")
	seed := flag.Int64("z", 1, "seed of random number generator")

	flag.Parse()
	rand.Seed(*seed)

	if *help {
		flag.PrintDefaults()
	} else {
		// compute the coenocline matrix
		mtx := Coenocline(*nSpec, *nSamp, *sampModel, *optModel, *abuModel, *tolModel, *μOpt, *σOpt, *μAbu, *σAbu, *μTol, *σTol)
		// and write it out as CSV (to be reimplemented using csv.WriteAll)
		for i := 0; i < *nSamp; i++ {
			for j := 0; j < *nSpec; j++ {
				fmt.Print(mtx.Get(i, j), ",")
			}
			fmt.Println()
		}
	}
}
