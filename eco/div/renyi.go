// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package div

// Rényi entropy diversity.

import (
	"code.google.com/p/go-eco/eco/aux"
	"math"
)

// RenyiDiv returns vector of Rényi entropy diversities. 
func RenyiDiv(data *aux.Matrix, q float64) *aux.Vector {
	// The Rényi entropy is a generalization of the Shannon entropy to other values of q than unity. It can be expressed:
	//    {}^qH = \frac{1}{1-q} \; \log \sum_{i=1}^R p_i^q
	// which equals
	//    {}^qH = \log {1 \over \sqrt[q-1]{{\sum_{i=1}^R p_i p_i^{q-1}}}} = \log({}^q\!D)
	//This means that taking the logarithm of true diversity based on any value of q gives the Rényi entropy corresponding to the same value of q.
	rows := data.R
	cols := data.C
	div := aux.NewVector(rows)
	for i := 0; i < rows; i++ {
		sum := 0.0
		tot := 0.0 // total number of all individuals in the sample
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0 {
				tot += x
			}
		}

		for j := 0; j < cols; j++ {
			p := data.Get(i, j) / tot
			if p > 0 {
				sum += math.Pow(p, q)
			}
		}
		div.Set(i, 1/(1-q)*math.Log(sum))
	}
	return div
}
