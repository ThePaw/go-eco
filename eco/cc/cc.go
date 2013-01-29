// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package cc

// Coenocline modelling functions. 
// To do: 
// implement carrying capacity Austin 2006
// implement species interactions

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

type Models struct {
	SRF        string  // type of Species response Function
	Samp       int     // spacing of samples
	Opt        int     // distribution of optima
	OptLoc     float64 // its mean
	OptScale   float64 // and tolerance
	Max        int     // distribution of maximum values of population densities
	MaxLoc     float64 // its mean
	MaxScale   float64 // and tolerance
	Tol        int     // distribution of species tolerances
	TolLoc     float64 // its mean
	TolScale   float64 // and tolerance
	MaxTolRho  float64 // correlation between population size and tolerance
	Beta       int     // distribution of alpha and gamma for beta function
	AlphaLoc   float64 // its mean
	AlphaScale float64 // and tolerance
	GammaLoc   float64 // its mean
	GammaScale float64 // and tolerance
	Noise      float64 // amount of external noise
}

// SetUpModels fills in a structure that holds model types and params. 
func (m *Models) SetUpModels(srfModel string, sampModel, optModel, popModel, tolModel, betaParamModel int, μOpt, εOpt, μMax, εMax, μTol, εTol, ρMaxTol, μα, εα, μγ, εγ, εNoise float64) {
	// Loc = location, Scale = scale
	m.SRF = srfModel
	m.Samp = sampModel
	m.Opt = optModel
	m.OptLoc = μOpt
	m.OptScale = εOpt // relative to 1, the length of the gradient
	m.Max = popModel
	m.MaxLoc = μMax
	m.MaxScale = εMax // relative to MaxLoc
	m.Tol = tolModel
	m.TolLoc = μTol   // relative to  1, the length of the gradient
	m.TolScale = εTol // relative to  TolLoc
	m.MaxTolRho = ρMaxTol
	m.Beta = betaParamModel
	m.AlphaLoc = μα
	m.AlphaScale = εα
	m.GammaLoc = μγ
	m.GammaScale = εγ
	m.Noise = εNoise

	// because some ε are originally relative numbers, recalculate
	m.MaxScale *= m.MaxLoc
	m.TolScale *= m.TolLoc
	return
}

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

// GaussSRF  computes the Gaussian response function. 
func GaussSRF(x float64, par ...float64) float64 {
	// opt	 	optimum = modus = position of max. population size on the gradient (= mean for Gaussian)
	// tol		tolerance (= range in Austin, 2006)
	// x		point on the gradient
	opt, tol := par[0], par[1]
	σ := tol / 6 // 99.7% of population lies within μ±3*σ for Normal (Gaussian) distribution
	y := NormalPDFAt(opt, σ, opt)
	return NormalPDFAt(opt, σ, x) / y // normalize so that maximum value is 1.0
}

// BetaSRF computes the β-function of Austin, 1976
func BetaSRF(x float64, par ...float64) (y float64) {
	// Austin, M. P. (1976). On non-linear species response models in ordination. Vegetatio, 33(1), 33-41.
	// http://cc.oulu.fi/~jarioksa/softhelp/betasimu.htm

	// opt	optimum = modus = position of max. population size on the gradient (= mean for Gaussian)
	// tol	tolerance (= range in Austin, 1994)
	// α, γ	params of β-function
	// x		point on the gradient

	opt, tol, α, γ := par[0], par[1], par[2], par[3]

	// calculate lower and upper limit from optimum and tolerance
	// needs to be reimplemented: just a quick hack valid only for symmetric (α = γ) cases
	lo := opt - 0.5*tol
	hi := opt + 0.5*tol

	// Return zero if x is not in (lo,hi)
	if x <= lo || x >= hi {
		y = 0
	} else {
		// Solve k from the maximum height of the response function
		// thanks to Jari Oksanen, betasimu.c
		max := 1.0 // desired maximum value
		tol := hi - lo
		t4 := tol / (α + γ)
		t6 := math.Pow(α*t4, α)
		t11 := math.Pow(γ*t4, γ)
		k := max / t6 / t11
		// evaluate the beta-function at x
		t2 := math.Pow(x-lo, α)
		t3 := math.Pow(hi-x, γ)
		y = k * t2 * t3
	}
	return
}

//  External influences modelled as Gaussian "error". 
func noise(μ, ε float64) float64 {
	// μ	 	theoretical value of density according to model
	// ε		scale of noise, relative to μ=1

	σ := ε / 6
	y := NormalNext(μ, σ)
	if y < 0 {
		y = 0
	}
	return y
}

// Random variable generator for optima, population sizes and tolerances, and for α, γ params of the β-function
func rndFn(which int, μ, ε float64) func() (x float64) {
	// μ	mean or modus
	// ε	tolerance

	return func() (x float64) {
		switch which {
		case flat: // Uniform distribution
			lo := μ - 0.5*ε
			hi := μ + 0.5*ε
			x = UniformNext(lo, hi)
		case gaussian: // Gaussian distribution
			σ := ε / 6
			x = NormalNext(μ, σ)
			/*
						case levy: // Lévy (unshifted)
							x = LevyNext(6/ε)	// terrible ad-hockery, empirical
						case paretoI: // Pareto I
				minValue := 0.01
							x = ParetoNext(minValue, 6/ε)	// terrible ad-hockery, empirical
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

func responseFn(which string) func(x float64, par ...float64) float64 {
	fn := GaussSRF
	switch which {
	case "gauss": // Gaussian distribution
		fn = GaussSRF
	case "beta": // β-function of Austin, 1976
		fn = BetaSRF
	}
	return fn
}

// Coenocline returns a matrix of species responses along (environmental) gradient. 
func Coenocline(nSpec, nSamp int, m Models) (out *Matrix) {
	var (
		opt, tol, aMax float64
	)

	out = NewMatrix(nSamp, nSpec)

	// if maxima and tolerance models are both Gaussian, allow for covariance:  model is Multivariate Normal
	mvNorm := false
	if m.Max == gaussian && m.Tol == gaussian {
		mvNorm = true
	}

	srf := responseFn(m.SRF)
	points := generate_points(nSamp, m.Samp)   // generate sampling points
	rngO := rndFn(m.Opt, m.OptLoc, m.OptScale) // optima distribution model
	rngM := rndFn(m.Max, m.MaxLoc, m.MaxScale) // maxima distribution model
	rngT := rndFn(m.Tol, m.TolLoc, m.TolScale)
	rngA := rndFn(m.Beta, m.AlphaLoc, m.AlphaScale) // alpha distribution model
	rngG := rndFn(m.Beta, m.GammaLoc, m.GammaScale) // gamma distribution model
	α, γ := 0.0, 0.0                                // not used outside beta model, just dummy values

	// for every species: 
	for j := 0; j < nSpec; j++ {

		// generate optimum (point on the gradient)
		opt = rngO()

		// generate species' population size and tolerance
		if mvNorm { // population size and tolerance model is Multivariate Normal
			// func Zeros(rows, cols int) *DenseMatrix
			mu := mtx.Zeros(2, 1)
			mu.Set(0, 0, m.MaxLoc)
			mu.Set(1, 0, m.TolLoc)
			cov := mtx.Ones(2, 2)
			cov.Set(0, 0, m.MaxScale*m.MaxScale) // needs justification
			cov.Set(0, 1, m.MaxTolRho*m.MaxScale*m.TolScale)
			cov.Set(1, 0, m.MaxTolRho*m.MaxScale*m.TolScale)
			cov.Set(1, 1, m.TolScale*m.TolScale)
			// func MVNormalNext(opt *DenseMatrix, Σ *DenseMatrix) *DenseMatrix
			mvMat := MVNormalNext(mu, cov)
			aMax = mvMat.Get(0, 0)
			tol = mvMat.Get(1, 0)
		} else {
			aMax = rngM()
			tol = rngT() // generate tolerance (range of acceptable gradient values)
		}

		// force max population density and tolerance within some considerable limits
		lo := m.MaxLoc - 3*m.MaxScale
		if lo < m.MaxLoc * 0.05 {
			lo = m.MaxLoc * 0.05
		}
		if aMax < lo {
			aMax = lo
		}
		hi := m.MaxLoc + 3*m.MaxScale
		if aMax > hi {
			aMax = hi
		}
		lo = m.TolLoc - 3*m.TolScale
		if lo < 0.1 * m.TolScale {
			lo = 0.1 * m.TolScale
		}
		if tol < lo {
			tol = lo
		}
		hi = m.TolLoc + 3*m.TolScale
		if tol > hi {
			tol = hi
		}

		// skew beta towards gradient optimum (Austin 2006): swap α, γ if needed
		if m.SRF == "beta" {
			α = rngA()
			γ = rngG()
			if m.Opt == gaussian && opt > m.OptLoc {
				if α > γ {
					swap := α
					α = γ
					γ = swap
				}
			} else {
				if α > γ {
					swap := α
					α = γ
					γ = swap
				}
			}
		}

		for i := 0; i < nSamp; i++ {
			x := points[i]
			y := srf(x, opt, tol, α, γ)
			// add "noise", if required
			if m.Noise > 0 {
				y = noise(y, m.Noise)
			}
			// scale by max population density
			y *= aMax
			out.Set(i, j, y)
		}
	}
	return
}

// Something went wrong here... needs revision.  
// GeneralisedBeta response function. 
// Austin et al. (2006), p. 200. 
func GenBetaSRF(x float64, par ...float64) float64 {
	// Austin, M. P., Belbin, L., Meyers, J. A., Doherty, M. D., & Luoto, M. (2006). Evaluation of statistical models used for predicting plant species distributions: 
	// Role of artificial data and theory. Ecological Modelling, 199(2), 197-216. 

	// opt	 	optimum = modus = position of max. population size on the gradient (= mean for Gaussian)
	// tol		tolerance (=range in Austin, 2006)
	// α, γ	params of β-function
	// x		point on the gradient

	opt, tol, α, γ := par[0], par[1], par[2], par[3]

	b := α / (α + γ)
	d := math.Pow(b, α) * math.Pow(1-b, γ)
	e := (x-opt)/tol + b
	a := (math.Pow(e, α) * math.Pow(1-e, γ)) / d
	return a
}
