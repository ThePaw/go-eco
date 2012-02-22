// Pielou's evenness (J)

package div

import (
	"math"
	. "go-eco.googlecode.com/hg/eco"
	. "go-eco.googlecode.com/hg/eco/rich"
)

func Pielou_E(data *Matrix, base byte, corr bool) *Vector {
	rows := data.R
	cols := data.C
	hh := Shannon(data, base, corr)
	ss := SObs(data)
	j := NewVector(cols)

	for i := 0; i < rows; i++ {
		s:=ss.Get(i)
		h:=hh.Get(i)
		j.Set(i, h/math.Log(s))
	}
	return j
}
