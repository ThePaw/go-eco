package ser

import (
	"fmt"
	"math"
	"math/rand"
)

// BayesSer does a Bayesian seriation of a species abundance matrix.
func BayesSer(counts IntMatrix, durations Matrix64, prior byte, nIter int, seed int64) (rhoH IntVector, h, pOH IntMatrix, rhoMean, qMean float64, hitsMean float64) {

	var (
		totalCost             int
		dist, flow, outCC     IntMatrix
		hitsSum, rhoSum, qSum float64
	)

	// init
	rand.Seed(seed)
	totalCost = 0
	hitsSum = 0
	rhoSum = 0
	qSum = 0
	empty := 0
	exp := 5.34
	exp2 := 1.34
	exp3 := 3.5
	di := 1
	fl := 2
	r := 5
	trials := 1
	improLagMax := 500
	verbose := false

	// alloc slices
	nSamp, nSpec := counts.Dims()
	rhoH = NewIntVector(20)          // rho histogram
	pOH = NewIntMatrix(nSamp, nSamp) // pair-order histogram
	h = NewIntMatrix(nSamp, nSamp)   // ranks histogram
	outCC = NewIntMatrix(nSamp, nSpec)

	for it := 0; it < nIter; it++ {
		// sample densities from abundances
		den := Densities(counts, durations, 1)

		// round to integers
		imtx := den.IntRound()
		//imtx.Print()
		//				fmt.Println("-------------------------------")

		// prepare distance and flow matrices
		switch di {
		case 0:
			dist = distances(nSamp)
		case 1:
			dist = distances2(nSamp, exp3)
		case 2:
			dist = distances3(nSamp, exp3)
		case 3:
			dist = distances4(nSamp, exp3)
		default:
			dist = distances3(nSamp, exp3)
		}

		switch fl {
		case 0:
			flow = flows(imtx, exp)
		case 1:
			flow = flows2(imtx, exp, exp2)
		case 2:
			flow = flows3(imtx, exp, exp2)
		default:
			flow = flows3(imtx, exp, exp2)
		}

		// solve QAP
		cc, p := QAP_fant(dist, flow, trials, improLagMax, r)
		totalCost += cc

		// reverse, if needed
		rho := reverseIfNeeded2(p)
		if verbose {
			if math.Abs(rho) < 0.5 {
				fmt.Println("rho: ", rho)
				fmt.Println("data: ")
				imtx.Print()
				fmt.Println("best solution: ")
				//	p.Print()
				fmt.Println("-------------------------------")
			}
		}

		// rank correlation
		rr := math.Abs(rho)
		rhoSum += rr

		// add pair-orders to pair-order histogram
		addToPairOrderHistogram(p, pOH)

		// add ranks to histogram
		addToRankHistogram(p, h)

		// add rho to histogram
		addToRhoHistogram(rr, rhoH)

		// update perfect hits
		if rr == 1 {
			hitsSum++
		}

		// is sorted Coenocline a Q-matrix?
		for i := 0; i < nSamp; i++ {
			for j := 0; j < nSpec; j++ {
				outCC[i][j] = imtx[p[i]][j]
			}
		}
		if outCC.IsQ() {
			qSum++
		}

	}

	// turn sums to proportions
	hitsMean = hitsSum / float64(nIter)
	rhoMean = rhoSum / float64(nIter)
	qMean = qSum / float64(nIter)
	fmt.Println("Empty samples: ", empty)
	fmt.Println()

	return
}
