package main

/*ToDo: k --> lo, hi
// Beta response function 
// Austin, M.P., 1976. On non-linear species responses models in ordination. Vegetatio 33, 33-41. DOI: 10.1007/BF00055297
// Austin, M.P., Gaywood, M.J., 1994. Current problems of environmental gradients and species response curves in relation to continuum theory. J. Veg. Sci. 5, 473-482. DOI: 10.2307/3235973
// This is NOT the Beta PDF !
// thanks to Jari Oksanen, betasimu.c

func betaSRF(opt, tol, max, α, γ, x float64) (y float64) {
	// opt is where first derivative is zero
	// solve lo, hi
	/////gnuplot> f(x)=k*(x-l)**a * (h-x)**g

	lo := max-tol
	// Return zero if x is not in (lo,hi)
	if x <= lo || x >= hi {
		y = 0
	} else {
		// Otherwise evaluate the beta-function at x
		k := kSolve(tol, α, γ, max)
		t2 := math.Pow(x-lo, α)
		t3 := math.Pow(hi-x, γ)
		y = k * t2 * t3
	}
	return
}

// Solve k from the maximum height of the response function
// thanks to Jari Oksanen, betasimu.c
func kSolve(tol, max, α, γ float64) (k float64) {
	t4 := tol / (α + γ)
	t6 := math.Pow(α*t4, α)
	t11 := math.Pow(γ*t4, γ)
	return max / t6 / t11
}
*/
