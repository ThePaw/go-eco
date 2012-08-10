package sim

import (
	mtx "code.google.com/p/go-eco/eco"
	"fmt"

	"testing"
)

// Ruggiero test against R:vegan, big data
func TestRuggiero(t *testing.T) {
	var (
		data, out, known *mtx.Matrix
	)

	fmt.Println("Ruggiero test against R:vegan, big data")
	data = GetBoolData()
	out = RuggieroBool_S(data)

	//known distances
	dist := [...]float64{0.0000000, 0.5208333, 0.5531915, 0.5000000, 0.5510204, 0.58,
		0.5208333, 0.0000000, 0.4042553, 0.5208333, 0.4897959, 0.42,
		0.5531915, 0.4042553, 0.0000000, 0.4583333, 0.5714286, 0.50,
		0.5000000, 0.5208333, 0.4583333, 0.0000000, 0.4081633, 0.46,
		0.5510204, 0.4897959, 0.5714286, 0.4081633, 0.0000000, 0.46,
		0.5800000, 0.4200000, 0.5000000, 0.4600000, 0.4600000, 0.00}

	rows := data.R
	known = mtx.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			if !check(out.Get(i, j), known.Get(i, j)) {
				t.Error()
				fmt.Println(i, j, out.Get(i, j), known.Get(i, j))
			}

		}
	}
}
