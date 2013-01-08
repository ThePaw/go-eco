// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package div

// Pielou's evenness (J)

import (
	"code.google.com/p/go-eco/eco/aux"
	. "code.google.com/p/go-eco/eco/rich"
	"math"
)

// PielouEq returns vector of Pielou equitabilities. 
func PielouEq(data *aux.Matrix, base byte, corr bool) *aux.Vector {
	rows := data.R
	hh := ShannonDiv(data, base, corr)
	ss := SObs(data)
	j := aux.NewVector(rows)

	for i := 0; i < rows; i++ {
		s := ss.Get(i)
		h := hh.Get(i)
		j.Set(i, h/math.Log(s))
	}
	return j
}
