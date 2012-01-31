// Pielou's evenness (J)

package eco

import (
	"math"
)

func Pielou(data *Matrix, base byte, corr bool) *Vector {
	rows := data.R
	cols := data.C
	hh := Shannon(data, base, corr)
	ss := Richness(data)
	j := NewVector(cols)

	for i := 0; i < rows; i++ {
s:=ss.Get(i)
h:=hh.Get(i)
		j.Set(i, h/math.Log(s))
	}
	return j
}
