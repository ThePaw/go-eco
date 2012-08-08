package sim

import (
	"fmt"
	mtx "code.google.com/p/go-eco/eco"
	"testing"
)


// Canonical distance test against R:ade4
func TestCanonical(t *testing.T) {
	var (
		data, out, known *mtx.Matrix
	)

	fmt.Println("Canonical distance test against R:ade4")
	data = GetData()
	out = Canonical_D(data)

	//known distances
	dist := [...]float64{0.000000,3.801207,3.339846,3.674499,4.274500,3.563358,
3.801207,0.000000,5.075769,3.791670,3.469832,4.707330,
3.339846,5.075769,0.000000,4.744850,3.814911,5.023699,
3.674499,3.791670,4.744850,0.000000,3.575037,4.423284,
4.274500,3.469832,3.814911,3.575037,0.000000,3.940870,
3.563358,4.707330,5.023699,4.423284,3.940870,0.000000}

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


