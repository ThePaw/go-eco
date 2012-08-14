// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Millar distance

import (
	"code.google.com/p/go-eco/eco/aux"
)

// Millar_D returns an Millar distance matrix for floating-point data. 
// Millar distance  == Binomial distance. 
func Millar_D(data *aux.Matrix) *aux.Matrix {
	return Binomial_D(data)
}
