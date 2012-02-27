// Pielou's evenness (J)

package div

import (
	. "go-eco.googlecode.com/hg/eco"
	. "go-eco.googlecode.com/hg/eco/rich"
	"math"
)

func Pielou_E(data *Matrix, base byte, corr bool) *Vector {
	rows := data.R
	hh := Shannon(data, base, corr)
	ss := SObs(data)
	j := NewVector(rows)

	for i := 0; i < rows; i++ {
		s := ss.Get(i)
		h := hh.Get(i)
		j.Set(i, h/math.Log(s))
	}
	return j
}
