// Sorgenfrei similarity matrix
// Sorgenfrei (1959)

package sim

import (
	"code.google.com/p/go-eco/eco/aux"
)

// Sorgenfrei similarity matrix
func SorgenfreiBool_S(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = aux.GetABCD(data, i, j)
			v := a * a / ((a + b) * (a + c))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
