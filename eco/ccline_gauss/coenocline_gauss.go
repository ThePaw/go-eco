// Coenocline modeller: Gaussian SRF (species population density response function). 
// Models species population densities along (environmental) gradient sampled at distinct points. 
// Density model is Gaussian (Normal) PDF, where, for every species,  μ is its optimum on the gradient, 
// σ is a measure of tolerance,  99.7% of population lies within μ±3*σ. 
// μ and σ themselves, and total population size have meta-models: uniform | Gaussian | Beta. 
// Uniform and Beta distributions are reparametrized to use mean and standard deviation as parameters, for consistency.
// Influences external to the model are modelled as Gaussian "noise".
// Sample spacing model is uniform | exponential (Poisson process). 
// Gradient spans from 0 to 1. 
package main


// To do: -p 2 does not work: panic: σ too big, α, β out of range

import (
	. "code.google.com/p/go-eco/eco/aux"
	. "code.google.com/p/probab/dst"
	"flag"
	"fmt"
	m "github.com/skelterjohn/go.matrix"
	"math/rand"
	"sort"
)

const (
	flat = iota
	gaussian
	beta
	paretoI
	paretoII
	paretoIII
	paretoIV
	paretoG
	paretoTap
	paretoSing
	yule
	planck
	zeta
)

// Generate sampling points along the gradient
func generate_points(k, spacing int) (pts []float64) {
	pts = make([]float64, k)
	switch spacing {
	default: // regular spacing
		for i := 0; i < k; i++ {
			pts[i] = float64(i) / float64(k-1)
		}
	case 1: // uniform random spacing
		for i := 0; i < k; i++ {
			pts[i] = rand.Float64()
		}
		sort.Float64s(pts) // sort in increasing order
	case 2: // exponential random spacing (spacing at points of a Poisson process)
		λ := 1 / float64(k) // mean interval length
		pts[0] = 0
		for i := 1; i < k; i++ {
			pts[i] = pts[i-1] + ExponentialNext(λ)
		}
		// scale to [0..1]
		for i := 1; i < k; i++ {
			pts[i] /= pts[k-1]
		}
	}
	return
}

// Gaussian response function 
func gaussSRF(μ, σ, x float64) float64 {
	// μ	 	optimum = modus = position of max. population size on the gradient (= mean for Gaussian)
	// σ		measure of tolerance, 99.7% of population lies within μ±3*σ
	// x		point on the gradient
	return NormalPDFAt(μ, σ, x)
}

// Gaussian external influences ("error" in terms of Fisherian statistics)
func noise(μ, σ float64) float64 {
	// μ	 	theoretical value of density according to model
	// σ		standard deviation due to external (to the model) factors
	y := NormalNext(μ, σ)
	if y < 0 {
		y = 0
	}
	return y
}

// Random variable generator for optima, population sizes and tolerances
func rndFn(which int, μ, σ float64) func() (x float64) {
	return func() (x float64) {

		switch which {
		case flat: // Uniform distribution
			a, b := UniformReparamMeanStd(μ, σ)
			x = UniformNext(a, b)
		case gaussian: // Gaussian distribution
			x = NormalNext(μ, σ)
		case beta: // Beta
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
func Coenocline(nSpec, nSamp, sampModel, optModel, popModel, tolModel int, μOpt, σOpt, μPop, σPop, μTol, σTol, ρPopTol, σNoise float64) (out *Matrix) {
	var (
		μ, σ, popSize float64
	)

	σPop *= μPop	// because σPop is originally relative number
	σTol *= μTol	// same for tolerance

	out = NewMatrix(nSamp, nSpec)
	points := generate_points(nSamp, sampModel) // generate sampling points
	rngO := rndFn(optModel, μOpt, σOpt)         // optima distribution model

	// if population size and tolerance models are both Gaussian, allow for covariance:  model is Multivariate Normal
	mvNorm := false
	if popModel == gaussian && tolModel == gaussian {
		mvNorm = true
	}

	rngA := rndFn(popModel, μPop, σPop) // population size distribution model
	rngT := rndFn(tolModel, μTol, σTol) // tolerance distribution model

	// for every species: 
	for j := 0; j < nSpec; j++ {
		// generate optimum (point on the gradient)
		μ = rngO()

		// generate species' population size and tolerance
		if mvNorm { // population size and tolerance model is Multivariate Normal
			// func Zeros(rows, cols int) *DenseMatrix
			mu := m.Zeros(2, 1)
			mu.Set(0, 0, μPop)
			mu.Set(1, 0, μTol)
			cov := m.Ones(2, 2)
			cov.Set(0, 0, σPop*σPop) // needs justification
			cov.Set(0, 1, ρPopTol*σPop*σTol)
			cov.Set(1, 0, ρPopTol*σPop*σTol)
			cov.Set(1, 1, σTol*σTol)
			// func MVNormalNext(μ *DenseMatrix, Σ *DenseMatrix) *DenseMatrix
			mvMat := MVNormalNext(mu, cov)
			popSize = mvMat.Get(0, 0)
			σ = mvMat.Get(1, 0)
		} else {
			popSize = rngA()
			σ = rngT() // generate tolerance (range of acceptable gradient values)
		}

		// force population size and tolerance within some considerable limits
		lo := μPop - 3*σPop
		if lo < 0 {
			lo = 0
		}
		if popSize < lo {
			popSize = lo
		}
		hi := μPop + 3*σPop
		if popSize > hi {
			popSize = hi
		}
		lo = μTol - 3*σTol
		if lo < 0 {
			lo = 0.1*σTol
		}
		if σ < lo {
			σ = lo
		}
		hi = μTol + 3*σTol
		if σ > hi {
			σ = hi
		}

		for i := 0; i < nSamp; i++ {
			x := points[i]
			y := gaussSRF(μ, σ, x)
			// add "noise", if required
			if σNoise > 0 {
				y = noise(y, σNoise)
			}
			// scale by population size; needs reimplementation so that area under the curve is multiplied by popSize (maybe it is already OK? Maybe, yes.)
			y *= popSize
			out.Set(i, j, y)
		}
	}
	return
}

// Main function 
func main() {
	help := flag.Bool("h", false, "print usage")
	nSpec := flag.Int("n", 20, "number of species")
	nSamp := flag.Int("m", 200, "number of samples")
	sampModel := flag.Int("s", 0, "sampling model: 0 - regular, 1 - uniform, 2 - Poisson")
	optModel := flag.Int("o", 0, "optima model: 0 - flat, 1 - Gaussian, 2 - Beta...")
	μOpt := flag.Float64("om", 0.5, "optimum mean")
	σOpt := flag.Float64("os", 0.35, "optimum standard deviation")
	popModel := flag.Int("p", 0, "population size model: 0 - flat, 1 - Gaussian, 2 - Beta...")
	μPop := flag.Float64("pm", 10.0, "mean population size")
	σPop := flag.Float64("ps", 0.2, "standard deviation of population size, relative to population size = 1")
	tolModel := flag.Int("t", 1, "tolerance model: 0 - flat, 1 - Gaussian, 2 - Beta...")
	μTol := flag.Float64("tm", 0.08, "mean tolerance")
	σTol := flag.Float64("ts", 0.4, "standard deviation of tolerance, relative to tolerance = 1")
	ρPopTol := flag.Float64("r", 0.0, "correlation between ")
	σNoise := flag.Float64("e", 0.0, "outer noise as standard deviation")
	seed := flag.Int64("z", 1, "seed of random number generator")

	flag.Parse()
	rand.Seed(*seed)

	if *help {
		flag.PrintDefaults()
		return
	}
	// fmt.Println("μPop", *μPop, "σPop", *σPop, "μTol", *μTol, "σTol", *σTol)
	// compute the coenocline matrix
	mtx := Coenocline(*nSpec, *nSamp, *sampModel, *optModel, *popModel, *tolModel, *μOpt, *σOpt, *μPop, *σPop, *μTol, *σTol, *ρPopTol, *σNoise)
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
