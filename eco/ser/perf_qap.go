package ser

import (
	"code.google.com/p/go-eco/eco/cc"
	"fmt"
	"math"
	"math/rand"
)

func QAPserPerformance(srfModel string, nSpec, nSamp, sampModel, optModel, denModel, tolModel, betaParamModel int, μOpt, εOpt, μMax, εMax, μTol, εTol, ρMaxTol, μα, εα, μγ, εγ, εNoise float64, nIter, trials, improLagMax, r int, exp, exp2, exp3 float64, di, fl int, seed int64) (rhoH IntVector, h, pOH IntMatrix, rhoMean, qMean float64, hitsMean float64) {

	var (
		mod                     cc.Models
		totalCost               int
		dist, flow, imtx, outCC IntMatrix
		hitsSum, rhoSum, qSum   float64
	)

	// init
	rand.Seed(seed)
	totalCost = 0
	hitsSum = 0
	rhoSum = 0
	qSum = 0
	empty := 0
	verbose := false

	// alloc slices
	rhoH = NewIntVector(20)          // rho histogram
	pOH = NewIntMatrix(nSamp, nSamp) // pair-order histogram
	h = NewIntMatrix(nSamp, nSamp)   // ranks histogram
	outCC = NewIntMatrix(nSamp, nSpec)

	// set up coenocline models
	mp := &mod // pointer to models
	mp.SetUpModels(srfModel, sampModel, optModel, denModel, tolModel, betaParamModel, μOpt, εOpt, μMax, εMax, μTol, εTol, ρMaxTol, μα, εα, μγ, εγ, εNoise)

	for it := 0; it < nIter; it++ {
		isEmpty := true
		for isEmpty {
			// generate the coenocline matrix
			m := cc.Coenocline(nSpec, nSamp, mod)

			rows := nSamp
			cols := nSpec

			//   ******* Remove when cc is rewritten to return Matrix64
			// unload Matrix to [][]float64
			mtx := NewMatrix64(rows, cols)
			for i := 0; i < rows; i++ {
				for j := 0; j < cols; j++ {
					mtx[i][j] = m.Get(i, j)
				}
			}

			// apply sampler
			imtx = TruncSampler(mtx)

			//check whether there are no empty samples
			isEmpty = false
			for i := 0; i < rows; i++ {
				isEmpty = true
				for j := 0; j < cols; j++ {
					if imtx[i][j] > 0 {
						isEmpty = false
						break
					}
				}
				if isEmpty {
					empty++
					break
				}
			}
		}
		// imtx.Print()
		// fmt.Println()
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
		/*
			if rr < 0.3 {


				fmt.Println("best solution: ")
				p.Print()
				fmt.Println("----coenocline---------")
				imtx.Print()
				fmt.Println()
				fmt.Println("rho: ", rr)
				fmt.Println("----permuted coenocline---------")
				for i := 0; i < nSamp; i++ {
					for j := 0; j < nSpec; j++ {
						fmt.Printf("%d, ", imtx[p[i]][j])
					}
					fmt.Println()
				}

			}
		*/

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
