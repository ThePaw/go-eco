package ser

import (
	"math"
)

func HOFFit(model int, a, b, c, d, x float64) float64 {
	maxPar := inf - 1
	t1 := 0.0
	t2 := 1.0
	t4 := 0.0
	t5 := 0.0
	fit := 1.0
	switch model {
	case 4:
		t1 = -d * x
		fallthrough
	case 3:
		t2 = c + t1
		if t2 > maxPar {
			t2 = maxPar
		}
		t2 = 1.0 / (1 + math.Exp(t2))
		fallthrough
	case 2:
		t4 = b * x
		fallthrough
	case 1:
		t5 = a + t4
		if t5 > maxPar {
			t5 = maxPar
		}
		t5 = 1 + math.Exp(t5)
		fit = t2 / t5
	default:
		panic("invalid model")
	}
	return fit
}

// Sample mean and unbiased (Bessel correction) variance estimates for a data vector.
func SampleMeanVar(x []float64) (μ, σ2 float64) {
	// Arguments: 
	// x - vector of observations
	//
	// Returns: 
	// μ - mean estimator 
	// σ2 - variance estimator 

	var n int
	var m, m2 float64
	μ = 0.0  // sample mean
	σ2 = 0.0 // sample variance unbiased
	m = 0.0
	m2 = 0.0

	for _, val := range x {
		n++
		μ += val
		delta := val - m
		m += delta / float64(n)
		m2 += delta * (val - m)
	}

	σ2 = m2 / float64(n-1)
	μ /= float64(len(x))
	return
}
