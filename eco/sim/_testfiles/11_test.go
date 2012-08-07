package sim

import (
	"fmt"
	. "go-eco.googlecode.com/hg/eco"
	"testing"
)

// Chao - Jaccard similarity test against R:fossil #1
func TestChaoJaccard(t *testing.T) {
	fmt.Println("Chao - Jaccard similarity test against R:fossil #1")
	data := GetCounts()
	out := ChaoJaccard_S(data)

	//known values
	dist := [...]float64{1,0.8399285,0.7548488,
0.8399285,1,1,
0.7548488,1,1}

	rows := data.R
	known := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			if !check(out.Get(i, j), known.Get(i, j)) {
				t.Error()
				fmt.Println(i, j, out.Get(i, j), known.Get(i, j))
			}
		}
	}
}

// Chao - Jaccard similarity test against R:fossil  #2
func TestChaoJaccard2(t *testing.T) {
	fmt.Println("Chao - Jaccard similarity test against R:fossil #2")
	data := GetCounts2()
	out := ChaoJaccard_S(data)

	//known values
	dist := [...]float64{1,0.6866585,0.6866585,1}

	rows := data.R
	known := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			if !check(out.Get(i, j), known.Get(i, j)) {
				t.Error()
				fmt.Println(i, j, out.Get(i, j), known.Get(i, j))
			}
		}
	}
}

// Chao - Sorensen similarity test against R:fossil
func TestChaoSorensen(t *testing.T) {
	fmt.Println("Chao - Sorensen similarity test against R:fossil")
	data := GetCounts()
	out := ChaoSorensen_S(data)

	//known values
	dist := [...]float64{1,0.9129781,0.8602875,
0.9129781,1,1,
0.8602875,1,1}

	rows := data.R
	known := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			if !check(out.Get(i, j), known.Get(i, j)) {
				t.Error()
				fmt.Println(i, j, out.Get(i, j), known.Get(i, j))
			}
		}
	}
}


