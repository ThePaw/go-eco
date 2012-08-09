// Smith and Wilson's evenness index 1-D 
// Smith & Wilson, 1996

package div

import (
	"code.google.com/p/go-eco/eco/aux"
	. "code.google.com/p/go-eco/eco/rich"
	"math"
)

// Smith and Wilson's evenness index 1-D 
// needs to be verified !
func SmithWilson1_E(data *aux.Matrix, which byte, small bool) *Vector {
	rows := data.R
	cols := data.C
	dd := Simpson(data, which, small)
	ss := SObs(data)
	out := NewVector(cols)

	for i := 0; i < rows; i++ {
		s := ss.Get(i)
		d := dd.Get(i)
		v := (1 - d) / (1 - 1/s)
		out.Set(i, v)
	}
	return out
}

// Smith and Wilson's evenness index B
func SmithWilson2_E(data *aux.Matrix) *Vector {
	rows := data.R
	cols := data.C
	out := NewVector(cols)

	for i := 0; i < rows; i++ {
		s := 0.0 // number of species
		sum1 := 0.0
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0.0 {
				s++
				sum1 += math.Log(x)
			}
		}
		sum1 /= s
		sum2 := 0.0
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0.0 {
				y := math.Log(x) - sum1
				sum2 += y * y
			}
		}
		v := 1.0 - (2.0 / (math.Pi * math.Atan(sum2)))
		out.Set(i, v)
	}
	return out
}
