// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package cc

// Coenocline modelling functions. 
// To do: model for α, γ 

import (
	. "code.google.com/p/go-eco/eco/aux"
	. "code.google.com/p/probab/dst"
	mtx "github.com/skelterjohn/go.matrix"
	"math"
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

// Generates sampling points along the gradient.
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

// Gaussian response function. 
func GaussSRF(x float64, par ...float64) float64 {
	// opt	 	optimum = modus = position of max. population size on the gradient (= mean for Gaussian)
	// tol		tolerance (=range in Austin, 2006)
	// x		point on the gradient
	opt, tol := par[0], par[1]
	σ := tol / 6 // 99.7% of population lies within μ±3*σ for Normal (Gaussian) distribution
	y := NormalPDFAt(opt, σ, opt)
	return NormalPDFAt(opt, σ, x) / y
}

// GeneralisedBeta response function. 
// Austin (2006), p. 200. 
func GenBetaSRF(x float64, par ...float64) float64 {
	// opt	 	optimum = modus = position of max. population size on the gradient (= mean for Gaussian)
	// tol		tolerance (=range in Austin, 2006)
	// α, γ	params of β-function
	// x		point on the gradient
	opt, tol, α, γ := par[0], par[1], par[2], par[3]
	b := α / (α + γ)
	d := math.Pow(b, α) * math.Pow(1-b, γ)
	e := (x-opt)/tol + b
	a := 1 / d * (math.Pow(e, α) * math.Pow(1-e, γ))
	return a
}

//  External influences modelled as Gaussian "error" (in terms of Fisherian statistics). 
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

type Models struct {
	SRF       string  // type of Species response Function
	Samp      int     // spacing of samples
	Opt       int     // distribution of optima
	OptMean   float64 // its mean
	OptSigma  float64 // and variance
	Pop       int     // distribution of population sizes
	PopMean   float64 // its mean
	PopSigma  float64 // and variance
	Tol       int     // distribution of species tolerances
	TolMean   float64 // its mean
	TolSigma  float64 // and variance
	PopTolRho float64 // correlation between population size and tolerance
	Noise     float64 // amount of external noise
}

// SetUpModels fills in a structure that holds model types and params. 
func (m *Models) SetUpModels(srfModel string, sampModel, optModel, popModel, tolModel int, μOpt, σOpt, μPop, σPop, μTol, σTol, ρPopTol, σNoise float64) {
	m.SRF = srfModel
	m.Samp = sampModel
	m.Opt = optModel
	m.OptMean = μOpt
	m.OptSigma = σOpt
	m.Pop = popModel
	m.PopMean = μPop
	m.PopSigma = σPop
	m.Tol = tolModel
	m.TolMean = μTol
	m.TolSigma = σTol
	m.PopTolRho = ρPopTol
	m.Noise = σNoise

	m.PopSigma *= m.PopMean // because σPop is originally relative number
	m.TolSigma *= m.TolMean // same for tolerance
	return
}

/*
// Coenocline returns a matrix of species responses along (environmental) gradient. 
func Coenocline(nSpec, nSamp int, m Models) (out *Matrix) {
	var (
		μ, σ, popSize float64
	)

	out = NewMatrix(nSamp, nSpec)

	srf := responseFn(m.SRF)
	points := generate_points(nSamp, m.Samp)    // generate sampling points
	rngO := rndFn(m.Opt, m.OptMean, m.OptSigma) // optima distribution model

	// if population size and tolerance models are both Gaussian, allow for covariance:  model is Multivariate Normal
	mvNorm := false
	if m.Pop == gaussian && m.Tol == gaussian {
		mvNorm = true
	}

	rngA := rndFn(m.Pop, m.PopMean, m.PopSigma) // population size distribution model
	rngT := rndFn(m.Tol, m.TolMean, m.TolSigma) // tolerance distribution model

	// for every species: 
	for j := 0; j < nSpec; j++ {
		// generate optimum (point on the gradient)
		μ = rngO()

		// generate species' population size and tolerance
		if mvNorm { // population size and tolerance model is Multivariate Normal
			// func Zeros(rows, cols int) *DenseMatrix
			mu := mtx.Zeros(2, 1)
			mu.Set(0, 0, m.PopMean)
			mu.Set(1, 0, m.TolMean)
			cov := mtx.Ones(2, 2)
			cov.Set(0, 0, m.PopSigma*m.PopSigma) // needs justification
			cov.Set(0, 1, m.PopTolRho*m.PopSigma*m.TolSigma)
			cov.Set(1, 0, m.PopTolRho*m.PopSigma*m.TolSigma)
			cov.Set(1, 1, m.TolSigma*m.TolSigma)
			// func MVNormalNext(μ *DenseMatrix, Σ *DenseMatrix) *DenseMatrix
			mvMat := MVNormalNext(mu, cov)
			popSize = mvMat.Get(0, 0)
			σ = mvMat.Get(1, 0)
		} else {
			popSize = rngA()
			σ = rngT() // generate tolerance (range of acceptable gradient values)
		}

		// force population size and tolerance within some considerable limits
		lo := m.PopMean - 3*m.PopSigma
		if lo < 0 {
			lo = 0
		}
		if popSize < lo {
			popSize = lo
		}
		hi := m.PopMean + 3*m.PopSigma
		if popSize > hi {
			popSize = hi
		}
		lo = m.TolMean - 3*m.TolSigma
		if lo < 0 {
			lo = 0.1 * m.TolSigma
		}
		if σ < lo {
			σ = lo
		}
		hi = m.TolMean + 3*m.TolSigma
		if σ > hi {
			σ = hi
		}
		for i := 0; i < nSamp; i++ {
			x := points[i]
			y := srf(μ, σ, x)
			// add "noise", if required
			if m.Noise > 0 {
				y = noise(y, m.Noise)
			}
			// scale by population size; needs reimplementation so that area under the curve is multiplied by popSize (maybe it is already OK? Maybe, yes.)
			y *= popSize
			out.Set(i, j, y)
		}
	}
	return
}
*/
func responseFn(which string) func(x float64, par ...float64) float64 {
	fn := GaussSRF
	switch which {
	case "gauss": // Gaussian distribution
		fn = GaussSRF
	case "beta": // Generalised Beta of Austin, 2006
		fn = GenBetaSRF
	}
	return fn
}

// Coenocline returns a matrix of species responses along (environmental) gradient. 
func Coenocline(nSpec, nSamp int, m Models) (out *Matrix) {
	var (
		opt, tol, aMax float64
	)

	out = NewMatrix(nSamp, nSpec)

	srf := responseFn(m.SRF)
	points := generate_points(nSamp, m.Samp)    // generate sampling points
	rngO := rndFn(m.Opt, m.OptMean, m.OptSigma) // optima distribution model

	// if population size and tolerance models are both Gaussian, allow for covariance:  model is Multivariate Normal
	mvNorm := false
	if m.Pop == gaussian && m.Tol == gaussian {
		mvNorm = true
	}

	rngA := rndFn(m.Pop, m.PopMean, m.PopSigma) // population size distribution model
	rngT := rndFn(m.Tol, m.TolMean, m.TolSigma) // tolerance distribution model

	// for every species: 
	for j := 0; j < nSpec; j++ {
		α, γ := 1.0, 1.0 // params of beta: to be reimplemented

		// generate optimum (point on the gradient)
		opt = rngO()

		// generate species' population size and tolerance
		if mvNorm { // population size and tolerance model is Multivariate Normal
			// func Zeros(rows, cols int) *DenseMatrix
			mu := mtx.Zeros(2, 1)
			mu.Set(0, 0, m.PopMean)
			mu.Set(1, 0, m.TolMean)
			cov := mtx.Ones(2, 2)
			cov.Set(0, 0, m.PopSigma*m.PopSigma) // needs justification
			cov.Set(0, 1, m.PopTolRho*m.PopSigma*m.TolSigma)
			cov.Set(1, 0, m.PopTolRho*m.PopSigma*m.TolSigma)
			cov.Set(1, 1, m.TolSigma*m.TolSigma)
			// func MVNormalNext(opt *DenseMatrix, Σ *DenseMatrix) *DenseMatrix
			mvMat := MVNormalNext(mu, cov)
			aMax = mvMat.Get(0, 0)
			tol = mvMat.Get(1, 0)
		} else {
			aMax = rngA()
			tol = rngT() // generate tolerance (range of acceptable gradient values)
		}

		// force max abundance and tolerance within some considerable limits
		lo := m.PopMean - 3*m.PopSigma
		if lo < 0 {
			lo = 0
		}
		if aMax < lo {
			aMax = lo
		}
		hi := m.PopMean + 3*m.PopSigma
		if aMax > hi {
			aMax = hi
		}
		lo = m.TolMean - 3*m.TolSigma
		if lo < 0 {
			lo = 0.1 * m.TolSigma
		}
		if tol < lo {
			tol = lo
		}
		hi = m.TolMean + 3*m.TolSigma
		if tol > hi {
			tol = hi
		}
		for i := 0; i < nSamp; i++ {
			x := points[i]
			y := srf(x, opt, tol, α, γ)
			// add "noise", if required
			if m.Noise > 0 {
				y = noise(y, m.Noise)
			}
			// scale by max abundance
			y *= aMax
			out.Set(i, j, y)
		}
	}
	return
}
