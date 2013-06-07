package ser

import (
//	"fmt"
	"math"
//	"math/rand"
)

func ObjFnPerformance(sim Matrix64, objFn ObjFn, isLoss, isDistFn  bool,  optMethod OptMethod, nIter int) (rhoH IntVector, rankH, pOH IntMatrix, rhoMean, rhoStDev, rProp, hitsProp float64) {

	// init
	hitsSum := 0.0
	rhoSum := 0.0
	rhoM := 0.0
	rhoStDev = 0.0
	rSum := 0.0
	nSamp := sim.Rows()

	// alloc slices
	rhoH = NewIntVector(20)            // rho histogram
	pOH = NewIntMatrix(nSamp, nSamp)   // pair-order histogram
	rankH = NewIntMatrix(nSamp, nSamp) // ranks histogram
	aa := NewMatrix64(nSamp, nSamp)    // sorted similarity/distance matrix
	pKnown := NewIntVector(nSamp) // known permutation
	pKnown.Order()

	for it := 0; it < nIter; it++ {
		a := sim.Clone() // essential, because input matrix may be converted to distances!
//		a.ForceTo01()

		// if objFn is distance-based, convert to distances
		if isDistFn {
			a.SimToDist()
		}

		// solve for best permutation
		_, p := optMethod(a, objFn, isLoss)
		Segment4Opt(a, p, objFn, isLoss)


		// reverse, if needed
		rho := reverseIfNeeded2(p)
if !p.Equals(pKnown) {
p.Print()
}		// rank correlation
		rr := math.Abs(rho)

		// rho sample mean and unbiased (Bessel correction) variance estimates
		rhoSum += rr
		rhoDelta := rr - rhoM
		rhoM += rhoDelta / float64(it+1)
		rhoStDev += rhoDelta * (rr - rhoM)

		// add pair-orders to pair-order histogram
		addToPairOrderHistogram(p, pOH)

		// add ranks to histogram
		addToRankHistogram(p, rankH)

		// add rho to histogram
		addToRhoHistogram(rr, rhoH)

		// update perfect hits
		if p.Equals(pKnown) {
			hitsSum++
		}

		// is sorted similarity/distance matrix (A)R-matrix?
		for i := 0; i < nSamp; i++ {
			for j := 0; j < nSamp; j++ {
				aa[i][j] = a[p[i]][p[j]]
			}
		}

		if isDistFn {
			if aa.IsAR() {
				rSum++
			}
		} else {
			if aa.IsR() {
				rSum++
			}
		}

	}

	// calc mean and st. deviation
	rhoMean = rhoSum / float64(nIter)

	rhoStDev /= float64(nIter - 1)
	rhoStDev = math.Sqrt(rhoStDev)

	// calc proportions
	hitsProp = hitsSum / float64(nIter)
	rProp = rSum / float64(nIter)
	return
}


// ==================
func ObjFnPerformance2(sim Matrix64, objFn ObjFn, isLoss, isDistFn  bool, nIter int) (rhoH IntVector, rankH, pOH IntMatrix, rhoMean, rhoStDev, rProp, hitsProp float64) {

	// init
	hitsSum := 0.0
	rhoSum := 0.0
	rhoM := 0.0
	rhoStDev = 0.0
	rSum := 0.0
	nSamp := sim.Rows()

	// alloc slices
	rhoH = NewIntVector(20)            // rho histogram
	pOH = NewIntMatrix(nSamp, nSamp)   // pair-order histogram
	rankH = NewIntMatrix(nSamp, nSamp) // ranks histogram
	aa := NewMatrix64(nSamp, nSamp)    // sorted similarity/distance matrix
	pKnown := NewIntVector(nSamp) // known permutation
	pKnown.Order()
	p:= pKnown.Clone()
	p.Perm()

	for it := 0; it < nIter; it++ {
		a := sim.Clone() // essential, because input matrix may be converted to distances!
//		a.ForceTo01()

		// if objFn is distance-based, convert to distances
		if isDistFn {
			a.SimToDist()
		}

		// solve for best permutation using SA followed by FA
		_, p = RobSA2(a, p, objFn, isLoss)
		_, p = RobFA2(a, p, objFn, isLoss)

		// reverse, if needed
		rho := reverseIfNeeded2(p)
//p.Print()
		// rank correlation
		rr := math.Abs(rho)

		// rho sample mean and unbiased (Bessel correction) variance estimates
		rhoSum += rr
		rhoDelta := rr - rhoM
		rhoM += rhoDelta / float64(it+1)
		rhoStDev += rhoDelta * (rr - rhoM)

		// add pair-orders to pair-order histogram
		addToPairOrderHistogram(p, pOH)

		// add ranks to histogram
		addToRankHistogram(p, rankH)

		// add rho to histogram
		addToRhoHistogram(rr, rhoH)

		// update perfect hits
		if p.Equals(pKnown) {
			hitsSum++
		}

		// is sorted similarity/distance matrix (A)R-matrix?
		for i := 0; i < nSamp; i++ {
			for j := 0; j < nSamp; j++ {
				aa[i][j] = a[p[i]][p[j]]
			}
		}

		if isDistFn {
			if aa.IsAR() {
				rSum++
			}
		} else {
			if aa.IsR() {
				rSum++
			}
		}

	}

	// calc mean and st. deviation
	rhoMean = rhoSum / float64(nIter)

	rhoStDev /= float64(nIter - 1)
	rhoStDev = math.Sqrt(rhoStDev)

	// calc proportions
	hitsProp = hitsSum / float64(nIter)
	rProp = rSum / float64(nIter)
	return
}
