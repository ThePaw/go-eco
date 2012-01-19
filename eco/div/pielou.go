// Pielou's evenness (J)

package eco

import (
	"math"
)

func Pielou(data *Matrix) *Vector {
	rows := data.R
	cols := data.C
	h := Shannon(data)
	s := richness(data)
	j := NewVector(cols)


	for i := 0; i < rows; i++ {
		j.Set(i, h/math.Log(s))
	}
	return j
}
