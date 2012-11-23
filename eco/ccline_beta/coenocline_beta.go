// Coenocline modeller: Beta SRF (inspired by Jari Oksanen).
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

// Beta response function 
// Austin, M.P., 1976. On non-linear species responses models in ordination. Vegetatio 33, 33-41. DOI: 10.1007/BF00055297
// Austin, M.P., Gaywood, M.J., 1994. Current problems of environmental gradients and species response curves in relation to continuum theory. J. Veg. Sci. 5, 473-482. DOI: 10.2307/3235973
// thanks to Jari Oksanen, betasimu.c

func betaSRF(k, lo, hi, max, α, γ, x float64) (y float64) {
	// Return zero if x is not in (lo,hi)
	if x <= lo || x >= hi {
		y = 0
	} else {
		// Otherwise evaluate the beta-function at x
		k := kSolve(lo, hi, α, γ, max)
		t2 := math.Pow(x-lo, α)
		t3 := math.Pow(hi-x, γ)
		y = k * t2 * t3
	}
	return
}

// Solve k from the maximum height of the response function
// thanks to Jari Oksanen, betasimu.c
func kSolve(lo, hi, α, γ, max float64) (k float64) {
	tol := hi-lo
	t4 := tol / (α + γ)
	t6 := math.Pow(α*t4, α)
	t11 := math.Pow(γ*t4, γ)
	return max / t6 / t11
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
func Coenocline(nSpec, nSamp, sampModel, optModel, abuModel, tolModel, αγModel int, μOpt, σOpt, μAbu, σAbu, μTol, σTol, μαγ, σαγ float64) (out *Matrix) {
	out = NewMatrix(nSamp, nSpec)
	points := generate_points(nSamp, sampModel) // generate sampling points
	rngO := rndFn(optModel, μOpt, σOpt)       // tolerance-interval center distribution model
	rngA := rndFn(abuModel, μAbu, σAbu)       // abundance distribution model
	rngT := rndFn(tolModel, μTol, σTol)       // tolerance distribution model
	rngαγ := rndFn(αγModel, μαγ, σαγ)       // distribution model of α, γ params
	for j := 0; j < nSpec; j++ {              // for every species: 
		μ := rngO()                // generate center of the tolerance-interval
		σ := rngT()                // generate width of the tolerance-interval
		for σ < 1/float64(nSamp) { // if tolerance is too small, generate new value
			σ = rngT()
		}
		lo := μ-3*σ
		hi := μ+3*σ
		maxDensity := rngA() // generate species' maxiumum population density
		for maxDensity <= 0.02*μAbu { // if population density is too small, generate new value
			maxDensity = rngA()
		}
		α := rngαγ() // generate α
		γ := rngαγ() // generate γ
		for i := 0; i < nSamp; i++ {
			x := points[i]
			k := kSolve(lo, hi, α, γ, maxDensity)
			y := betaSRF(k, lo, hi, maxDensity, α, γ, x)
			out.Set(i, j, y)
		}
	}
	return
}

// Main function 
func main() {
	help := flag.Bool("h", false, "Coenocline modeller\nUsage: ccline_beta  [-nmso[om[os]]a[am[as]]t[tm[ts]]z]")
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
	αγModel := flag.Int("q", 1, "tolerance model: 0 - flat, 1 - Gaussian, 2 - Beta...")
	μαγ := flag.Float64("qm", 0.003, "mean tolerance")
	σαγ := flag.Float64("qs", 0.2, "standard deviation of tolerance")
	seed := flag.Int64("z", 1, "seed of random number generator")

	flag.Parse()
	rand.Seed(*seed)

	if *help {
		flag.PrintDefaults()
		return
	}
	// compute the coenocline matrix
	mtx := Coenocline(*nSpec, *nSamp, *sampModel, *optModel, *abuModel, *tolModel, *αγModel, *μOpt, *σOpt, *μAbu, *σAbu, *μTol, *σTol, *μαγ, *σαγ)
	// and write it out as CSV, transposed so that rows are species, columns are sampling points (to be reimplemented using csv.WriteAll)
	for i := 0; i < *nSpec; i++ {
		for j := 0; j < *nSamp; j++ {
			if j == 0 {
				fmt.Print(mtx.Get(j, i))
			} else {
				fmt.Print(",", mtx.Get(j, i))
			}
		}
		fmt.Println()
	}
}

