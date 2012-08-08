package div

import (
	"fmt"
	"testing"
	. "code.google.com/p/go-eco/eco"
)

// Koleff1 test against R:vegan:betadist
func TestKoleff1(t *testing.T) {
	fmt.Println("Koleff1 test against R:vegan:betadist")
	data := GetBoolData()
	out := Koleff1Beta(data)

	//known distances
	dist := [...]float64{0.0000000,0.5192308,0.4951456,0.5384615,0.4857143,0.4528302,
0.5192308,0.0000000,0.6000000,0.4791667,0.5051546,0.5714286,
0.4951456,0.6000000,0.0000000,0.5368421,0.4166667,0.4845361,
0.5384615,0.4791667,0.5368421,0.0000000,0.5876289,0.5306122,
0.4857143,0.5051546,0.4166667,0.5876289,0.0000000,0.5353535,
0.4528302,0.5714286,0.4845361,0.5306122,0.5353535,0.0000000}

	rows := data.R
	known := NewMatrix(rows, rows)
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


