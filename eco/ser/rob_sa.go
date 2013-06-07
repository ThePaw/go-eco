// Copyright 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package ser

// Sort the pre-(Anti)-Robinson matrix using Simulated Annealing.
// Use functions in obj_fn_sim.go for Robinson, obj_fn_dis.go for Anti-Robinson matrix.

import (
//	"fmt"
	"math"
	"math/rand"

//	"time"
)

type GenFn func(p IntVector)        // Generates a new solution from an old one
type CoolFn func(x float64) float64 // Cooling schedule: Generates a new temperature from the previous one.
//                  Any function that takes a scalar as input and
//                  returns a smaller but positive scalar as output. 

func cool(x float64) float64 {
	return 0.8 * x
}

// SimAnn minimizes a loss function using simulated annealing (Kirkpatrick et al., 1983)
func SimAnn(a Matrix64, loss ObjFn, isLoss bool, proposePerm GenFn, cool CoolFn) (bestEnergy float64, bestSolution IntVector) {
	//    Kirkpatrick, S., Gelatt, C.D., & Vecchi, M.P. (1983). Optimization by Simulated Annealing. Science, 220: 671-680.
	//    Based on "anneal.m" code by Joachim Vandekerckhove joachim.vandekerckhove@psy.kuleuven.be
	//    http://www.mathworks.com/matlabcentral/fileexchange/10548-general-simulated-annealing-algorithm/content/anneal.m

	var (
		seed      int64
		newEnergy float64
	)

	// params
	initT := 1.0 // The initial temperature, can be any positive number.
	minT := 1e-8 // Temperature at which to stop, can be any positive number
	//                  smaller than initT. 
	stopVal := -inf //Value at which to stop immediately, can be any output of
	//                  loss fn that is sufficiently low for you.
	maxConsRej := 1000 // Maximum number of consecutive rejections, can be any
	//                  positive number.
	maxTries := 300 // Maximum number of tries within one temperature, can be
	//                  any positive number.
	maxSuccess := 20 // Maximum number of successes within one temperature, can
	//                  be any positive number.
	k := 1.0 // Boltzmann constant.

	// init
	try := 0
	success := 0
	finished := false
	consec := 0
	temp := initT
	stopVal = -inf
	total := 0
	initialSolution := NewIntVector(a.Rows())
	initialSolution.Perm()
	bestSolution = initialSolution.Clone()
	newSolution := initialSolution.Clone()
	initialEnergy := loss(a, initialSolution)
	bestEnergy = initialEnergy

	// Create and seed the generator.
	// Typically a non-fixed seed should be used, such as time.Now().UnixNano().
	// Using a fixed seed will produce the same output on every run.
	// seed := time.Now().UnixNano()
	seed = 5
	r := rand.New(rand.NewSource(seed))

	for !finished {
		try++ // just an iteration counter
		newSolution.CopyFrom(bestSolution)

		//  Stop / decrement temp criteria
		if try >= maxTries || success >= maxSuccess {
			if temp < minT || consec >= maxConsRej {
				finished = true
				total = total + try
				break
			} else {
				temp = cool(temp) // decrease temp according to cooling schedule
				total = total + try
				try = 1
				success = 1
			}
		}

		proposePerm(newSolution)
		if isLoss {
			newEnergy = loss(a, newSolution)
		} else {
			newEnergy = -loss(a, newSolution)
		}
		if newEnergy < stopVal {
			bestSolution.CopyFrom(newSolution)
			bestEnergy = newEnergy
			break
		}

		if bestEnergy-newEnergy > 1e-6 {
			bestSolution.CopyFrom(newSolution)
			bestEnergy = newEnergy
			success++
			consec = 0
		} else {
			if r.Float64() < math.Exp((bestEnergy-newEnergy)/(k*temp)) {
				bestSolution.CopyFrom(newSolution)
				bestEnergy = newEnergy
				success++
			} else {
				consec++
			}
		}
	}

	return
}

func RobSA(a Matrix64, objFn ObjFn, isLoss bool) (bestEnergy float64, bestSolution IntVector) {
	bestEnergy, bestSolution = SimAnn(a, objFn, isLoss, proposePerm, cool)
	return
}

// ============= New versions =============

// SimAnn2 minimizes a loss function using simulated annealing (Kirkpatrick et al., 1983)
func SimAnn2(a Matrix64, initialSolution IntVector, loss ObjFn, isLoss bool, proposePerm GenFn, cool CoolFn) (bestEnergy float64, bestSolution IntVector) {
	//    Kirkpatrick, S., Gelatt, C.D., & Vecchi, M.P. (1983). Optimization by Simulated Annealing. Science, 220: 671-680.
	//    Based on "anneal.m" code by Joachim Vandekerckhove joachim.vandekerckhove@psy.kuleuven.be
	//    http://www.mathworks.com/matlabcentral/fileexchange/10548-general-simulated-annealing-algorithm/content/anneal.m

	var (
		seed      int64
		newEnergy float64
	)

	// params
	initT := 1.0 // The initial temperature, can be any positive number.
	minT := 1e-8 // Temperature at which to stop, can be any positive number
	//                  smaller than initT. 
	stopVal := -inf //Value at which to stop immediately, can be any output of
	//                  loss fn that is sufficiently low for you.
	maxConsRej := 1000 // Maximum number of consecutive rejections, can be any
	//                  positive number.
	maxTries := 300 // Maximum number of tries within one temperature, can be
	//                  any positive number.
	maxSuccess := 20 // Maximum number of successes within one temperature, can
	//                  be any positive number.
	k := 1.0 // Boltzmann constant.

	// init
	try := 0
	success := 0
	finished := false
	consec := 0
	temp := initT
	stopVal = -inf
	total := 0
	bestSolution = initialSolution.Clone()
	newSolution := initialSolution.Clone()
	initialEnergy := loss(a, initialSolution)
	bestEnergy = initialEnergy

	// Create and seed the generator.
	// Typically a non-fixed seed should be used, such as time.Now().UnixNano().
	// Using a fixed seed will produce the same output on every run.
	// seed := time.Now().UnixNano()
	seed = 5
	r := rand.New(rand.NewSource(seed))

	for !finished {
		try++ // just an iteration counter
		newSolution.CopyFrom(bestSolution)

		//  Stop / decrement temp criteria
		if try >= maxTries || success >= maxSuccess {
			if temp < minT || consec >= maxConsRej {
				finished = true
				total = total + try
				break
			} else {
				temp = cool(temp) // decrease temp according to cooling schedule
				total = total + try
				try = 1
				success = 1
			}
		}

		proposePerm(newSolution)
		if isLoss {
			newEnergy = loss(a, newSolution)
		} else {
			newEnergy = -loss(a, newSolution)
		}
		if newEnergy < stopVal {
			bestSolution.CopyFrom(newSolution)
			bestEnergy = newEnergy
			break
		}

		if bestEnergy-newEnergy > 1e-6 {
			bestSolution.CopyFrom(newSolution)
			bestEnergy = newEnergy
			success++
			consec = 0
		} else {
			if r.Float64() < math.Exp((bestEnergy-newEnergy)/(k*temp)) {
				bestSolution.CopyFrom(newSolution)
				bestEnergy = newEnergy
				success++
			} else {
				consec++
			}
		}
	}

	return
}

func RobSA2(a Matrix64, p IntVector, objFn ObjFn, isLoss bool) (bestEnergy float64, bestSolution IntVector) {
	bestEnergy, bestSolution = SimAnn2(a, p, objFn, isLoss, proposePerm, cool)
	return
}

// Segment4Opt modifies permutation by greedy optimization of all segments legth 4.
func Segment4Opt(a Matrix64, p IntVector, objFn ObjFn, isLoss bool) {

	perm4 := IntMatrix{
		{0, 1, 3, 2},
		{0, 2, 1, 3},
		{0, 2, 3, 1},
		{0, 3, 1, 2},
		{0, 3, 2, 1},
		{1, 0, 2, 3},
		{1, 0, 3, 2},
		{1, 2, 0, 3},
		{1, 2, 3, 0},
		{1, 3, 0, 2},
		{1, 3, 2, 0},
		{2, 0, 1, 3},
		{2, 0, 3, 1},
		{2, 1, 0, 3},
		{2, 1, 3, 0},
		{2, 3, 0, 1},
		{2, 3, 1, 0},
		{3, 0, 1, 2},
		{3, 0, 2, 1},
		{3, 1, 0, 2},
		{3, 1, 2, 0},
		{3, 2, 0, 1},
	}

	cost := objFn(a, p)

	if !isLoss {
		cost = -cost
	}
	best := cost
	seg := NewIntVector(4)
	n := p.Len() - 3
	for i := 0; i < n; i++ {
		for j := 0; j < 22; j++ {
			w := p.Clone()
			for k := 0; k < 4; k++ {
				seg[k] = w[i+perm4[j][k]]
			}
			for k := 0; k < 4; k++ {
				w[i+k] = seg[k]
			}
			cost = objFn(a, w)
			if !isLoss {
				cost = -cost
			}
			if cost < best {
//				fmt.Println("=== IMPROVED ===", cost, best, i, j)
				best = cost
				p.CopyFrom(w)
				j = 0
				if !p.IsPermutation() {
					seg.Print()
					p.Print()
					panic("not a permutation")
				}
			}
		}
	}
//	cost = objFn(a, p)
//					p.Print()
//				fmt.Println("=== Final cost:  ===", cost)
}
