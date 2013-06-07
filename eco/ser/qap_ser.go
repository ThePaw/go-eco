package ser

// Distance and flow functions specific to QAP seriation.

import (
	"math"
)

// distances calculates distance matrix for the QAP. 
func distances(smp int) (dist IntMatrix) {
	dist = NewIntMatrix(smp, smp)
	for i := 0; i < smp; i++ {
		for j := 0; j < smp; j++ {
			dist[i][j] = iAbs(i - j)
		}
	}
	return
}

// distances2 calculates distance matrix for the QAP. 
func distances2(smp int, exp float64) (dist IntMatrix) {
	dist = NewIntMatrix(smp, smp)
	for i := 0; i < smp; i++ {
		for j := 0; j < smp; j++ {
			x := math.Abs(float64(i - j))
			x = math.Pow(x, exp)
			dist[i][j] = int(x)
		}
	}
	return
}

// distances3 calculates distance matrix for the QAP. 
func distances3(smp int, root float64) (dist IntMatrix) {
	dist = NewIntMatrix(smp, smp)
	for i := 0; i < smp; i++ {
		for j := 0; j < smp; j++ {
			x := math.Abs(float64(i - j))
			x = math.Pow(root, x)
			dist[i][j] = int(x)
		}
	}
	return
}

// distances4 calculates distance matrix for the QAP. 
func distances4(smp int, c float64) (dist IntMatrix) {
	dist = NewIntMatrix(smp, smp)
	for i := 0; i < smp; i++ {
		for j := 0; j < smp; j++ {
			x := math.Abs(float64(i - j))
			x = math.Exp(c * x)
			dist[i][j] = int(x)
		}
	}
	return
}

// calculates flow matrix for the QAP. 
func flows(dat IntMatrix, exp float64) (flow IntMatrix) {
	const scale float64 = 1000
	var y float64
	smp := dat.Rows()
	spc := dat.Cols()
	flow = NewIntMatrix(smp, smp)
	if exp == 1 {
		for i := 0; i < smp; i++ {
			for j := 0; j < smp; j++ {
				x := 0
				if i != j {
					for k := 0; k < spc; k++ { // every species contributes
						x += dat[i][k] * dat[j][k]
					}
				}
				flow[i][j] = x
			}
		}
	} else {
		for i := 0; i < smp; i++ {
			for j := 0; j < smp; j++ {
				y = 0
				if i != j {
					for k := 0; k < spc; k++ { // every species contributes
						y += math.Pow(float64(dat[i][k]*dat[j][k]), 1.0/exp)
					}
				}
				// convert float64 values to int
				flow[i][j] = int(y)
			}
		}
	}
	return
}

// calculates flow matrix for the QAP  (generalized Minkowski metric).
func flows2(dat IntMatrix, exp, exp2 float64) (flow IntMatrix) {
	const scale float64 = 1000
	var (
		x float64
	)
	smp := dat.Rows()
	spc := dat.Cols()
	flow = NewIntMatrix(smp, smp)
	for i := 0; i < smp; i++ {
		for j := 0; j < smp; j++ {
			x = 0
			if i != j {
				for k := 0; k < spc; k++ { // every species contributes
					y := iAbs(dat[i][k] - dat[j][k])
					x += math.Pow(float64(y), exp)
				}
			}
			// convert float64 values to int
			x = math.Pow(x, 1.0/exp2)
			flow[i][j] = int(scale / x)
		}
	}
	return
}

// calculates just another flow matrix for the QAP.
func flows3(dat IntMatrix, exp, exp2 float64) (flow IntMatrix) {
	const scale float64 = 1000
	var (
		x, maxVal float64
	)

	scaleFlows := false
	smp := dat.Rows()
	spc := dat.Cols()
	flow = NewIntMatrix(smp, smp)
	maxVal = 0

	for i := 0; i < smp; i++ {
		for j := 0; j < smp; j++ {
			x = 0
			if i != j {
				for k := 0; k < spc; k++ { // every species contributes
					a := dat[i][k]
					b := dat[j][k]
					if a*b > 0 {
						x += math.Pow(float64(a*b), 1.0/exp)
					}
				}
			}
			x = math.Pow(x, exp2)
			if x > maxVal {
				maxVal = x
			}
			flow[i][j] = int(x)
		}
	}

	if scaleFlows {
		// scale to maxVal
		for i := 0; i < smp; i++ {
			for j := 0; j < smp; j++ {
				x = float64(flow[i][j])
				flow[i][j] = int(scale * x / maxVal)
			}
		}
	}
	return
}
