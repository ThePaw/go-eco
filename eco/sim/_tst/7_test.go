package sim

import (
	mtx "code.google.com/p/go-eco/eco"
	"fmt"

	"testing"
)

// Raup-Crick1 test against R:vegan, smaller data
func TestRaupCrick1(t *testing.T) {
	var (
		data, out, known *mtx.Matrix
	)

	fmt.Println("Raup-Crick1 test against R:vegan, smaller data")
	data = GetBoolData2()
	out = RaupCrick1Bool_S(data)

	//known distances
	dist := [...]float64{0.00000000, 0.45238095, 0.66666667, 0.11904762, 1.00000000, 0.97619048,
		0.45238095, 0.00000000, 0.83333333, 0.07142857, 0.88095238, 0.97619048,
		0.66666667, 0.83333333, 0.00000000, 0.66666667, 0.66666667, 1.00000000,
		0.11904762, 0.07142857, 0.66666667, 0.00000000, 1.00000000, 0.97619048,
		1.00000000, 0.88095238, 0.66666667, 1.00000000, 0.00000000, 0.73809524,
		0.97619048, 0.97619048, 1.00000000, 0.97619048, 0.73809524, 0.00000000}

	rows := data.R
	known = mtx.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			if i != j && !check(1-out.Get(i, j), known.Get(i, j)) {
				t.Error()
				fmt.Println(1-out.Get(i, j), known.Get(i, j))
			}

		}
	}
}

// Raup-Crick2 test against R:vegan, smaller data, flat prior
func TestRaupCrick2Bool(t *testing.T) {
	var (
		data, out, known *mtx.Matrix
	)

	fmt.Println("Raup-Crick2 test against R:vegan, smaller data")
	data = GetBoolData2()
	// p is flat:
	p := []float64{0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1}

	out = RaupCrick2Bool_S(data, p)

	//known distances
	dist := [...]float64{0.00000000, 0.45238095, 0.66666667, 0.11904762, 1.00000000, 0.97619048,
		0.45238095, 0.00000000, 0.83333333, 0.07142857, 0.88095238, 0.97619048,
		0.66666667, 0.83333333, 0.00000000, 0.66666667, 0.66666667, 1.00000000,
		0.11904762, 0.07142857, 0.66666667, 0.00000000, 1.00000000, 0.97619048,
		1.00000000, 0.88095238, 0.66666667, 1.00000000, 0.00000000, 0.73809524,
		0.97619048, 0.97619048, 1.00000000, 0.97619048, 0.73809524, 0.00000000}

	rows := data.R
	known = mtx.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			if i != j && !check(1-out.Get(i, j), known.Get(i, j)) {
				t.Error()
				fmt.Println(1-out.Get(i, j), known.Get(i, j))
			}

		}
	}
}
