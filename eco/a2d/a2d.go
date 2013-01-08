// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package a2d

// Posterior distribution of population density D inferred from abundance (=counts of individuals) A and sampling duration T. 
// T ~ N(μ, σ)
// A ~ Pois(T*D)

import (
	. "code.google.com/p/go-eco/eco/aux"
	bayes "code.google.com/p/probab/bayes"
	dst "code.google.com/p/probab/dst"
	"math"
	"math/rand"
)

// durationNext returns random number drawn from the distribution of duration (time-span represented by the sample). 
func durationNext(m, s float64) float64 {
	return dst.NormalNext(m, s)
}

// postDensityGNext returns random number drawn from the posterior distribution of population density 
// inferred from abundance and sampling duration, using conjugate
func postDensityNext(k int64, dur, m, s float64) float64 {
	// Use r=m^2/s^2, and v=m/s^2, if you summarize your prior belief (Gamma) with mean == m, and std == s.
	r := m * m / (s * s)
	v := m / (s * s)
	qtl := bayes.PoissonLambdaQtlGPri(k, 1, r, v) // using conjugate Gamma prior
	p := rand.Float64()
	lambda := qtl(p)
	// lambda = density*duration, thus density = lambda/duration
	return lambda / dur
}

// postDensityFNext returns random number drawn from the posterior distribution of population density inferred from abundance and sampling duration. 
func postDensityNext(k int64, dur float64) float64 {
	qtl := bayes.PoissonLambdaQtlFPri(k, 1) // using Flat prior
	p := rand.Float64()
	lambda := qtl(p)
	// lambda = density*duration, thus density = lambda/duration
	return lambda / dur
}

// postDensityJNext returns random number drawn from the posterior distribution of population density inferred from abundance and sampling duration. 
func postDensityNext(k int64, dur float64) float64 {
	qtl := bayes.PoissonLambdaQtlFPri(k, 1) // using Flat prior
	p := rand.Float64()
	lambda := qtl(p)
	// lambda = density*duration, thus density = lambda/duration
	return lambda / dur
}

// Densities returns a matrix of sampled posterior population densities. 
func Densities(counts, durations Matrix, prior byte) (out *Matrix) {
	// 'counts' : a nSamp*nSpec matrix of counts of individulals belonging to species j, in sample i.
	// 'durations' : a 2*nSamp matrix whose first row are means, second row are standard deviations of prior belief about duration, for every sample.

	const (
		flat = iota
		jeffreys
		gamma
	)
	nSamp := counts.R
	nSpec := counts.C
	out = NewMatrix(nSamp, nSpec) // sample of posterior pop. densities

	// for every sample: 
	for i := 0; i < nSamp; i++ {
		// generate duration mean and std
		m := durations.Get(0, i)
		s := durations.Get(1, i)
		dur := durationNext(m, s)

		// generate species' population densities
		for j := 0; j < nSpec; j++ {
			kf := counts.Get(i, j)
			k := int64(math.Floor(kf))
			switch prior {
			case flat:
				y := postDensityFNext(k, dur)
			case jeffreys:
				y := postDensityFNext(k, dur)
			case gamma:
				// just for now, needs to be reimplemented:
				mLambda := kf
				sLambda := 0.3 * mLambda
				y := postDensityGNext(k, dur, mLambda, sLambda)
			}
			out.Set(i, j, y)
		}
	}
	return
}
