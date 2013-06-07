// Copyright 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package ser

// Ccline encapsulates the Coenocline model.
type Ccline struct {
	srfModel       string //model of SRF: gaussian, beta.
	sampModel      int
	optModel       int
	denModel       int
	tolModel       int
	betaParamModel int
	μOpt           float64
	εOpt           float64
	μMax           float64
	εMax           float64
	μTol           float64
	εTol           float64
	ρMaxTol        float64
	μα             float64
	εα             float64
	μγ             float64
	εγ             float64
	εNoise         float64
}
