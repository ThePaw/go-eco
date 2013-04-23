package ser

import (
	//	"code.google.com/p/go-eco/eco/aux"
	"fmt"
	"math/rand"
	"testing"
)

func TestBoltzmann_1(t *testing.T) {
	fmt.Println("Testing Boltzmann: perfect Robinson matrix: Psi")
	var seed int64 = 1
	rand.Seed(seed)
	a := [][]int{ // perfect Robinson matrix
		{81, 64, 49, 17, 16, 10, 9, 4, 3, 0},
		{64, 81, 64, 49, 17, 16, 10, 9, 4, 3},
		{49, 64, 81, 64, 49, 17, 16, 10, 9, 4},
		{17, 49, 64, 81, 64, 49, 17, 16, 10, 9},
		{16, 17, 49, 64, 81, 64, 49, 17, 16, 10},
		{10, 16, 17, 49, 64, 81, 64, 49, 17, 16},
		{9, 10, 16, 17, 49, 64, 81, 64, 49, 17},
		{4, 9, 10, 16, 17, 49, 64, 81, 64, 49},
		{3, 4, 9, 10, 16, 17, 49, 64, 81, 64},
		{0, 3, 4, 9, 10, 16, 17, 49, 64, 81},
	}
	known := IntVector{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := len(a)
	dataType := "similarity"
	energyFn := "Psi"
	improLagMax := 10000
	initPermute := true
	permuteCols := false
	rowPerm := NewIntVector(n)
	for i := 0; i < n; i++ {
		rowPerm[i] = i
	}
	colPerm := NewIntVector(n)
	for i := 0; i < n; i++ {
		colPerm[i] = i
	}

	e, bestRowPerm, _ := Boltzmann(a, dataType, rowPerm, colPerm, energyFn, improLagMax, initPermute, permuteCols)
	bestRowPerm.Increasing()
	if !bestRowPerm.IsIdentical(known) {
		fmt.Println("Best permutation: ")
		bestRowPerm.Print()
		fmt.Println("should be: ")
		known.Print()
		fmt.Println("Energy: ", e)
		t.Error()
	}

	fmt.Println("Testing Boltzmann: perfect Robinson matrix: Psi2")
	energyFn = "Psi2"
	e, bestRowPerm, _ = Boltzmann(a, dataType, rowPerm, colPerm, energyFn, improLagMax, initPermute, permuteCols)
	bestRowPerm.Increasing()
	if !bestRowPerm.IsIdentical(known) {
		fmt.Println("Best permutation: ")
		bestRowPerm.Print()
		fmt.Println("should be: ")
		known.Print()
		fmt.Println("Energy: ", e)
		t.Error()
	}

	fmt.Println("Testing Boltzmann: perfect Robinson matrix: Psi3")
	energyFn = "Psi3"
	e, bestRowPerm, _ = Boltzmann(a, dataType, rowPerm, colPerm, energyFn, improLagMax, initPermute, permuteCols)
	bestRowPerm.Increasing()
	if !bestRowPerm.IsIdentical(known) {
		fmt.Println("Best permutation: ")
		bestRowPerm.Print()
		fmt.Println("should be: ")
		known.Print()
		fmt.Println("Energy: ", e)
		t.Error()
	}

}

func TestBoltzmann_2(t *testing.T) {
	fmt.Println("Testing Boltzmann: permuted perfect Robinson matrix")
	var seed int64 = 7
	rand.Seed(seed)
	a := [][]int{ // permuted perfect Robinson matrix
		{81, 17, 49, 64, 64, 49, 17, 16, 10, 9},
		{17, 81, 64, 49, 16, 10, 9, 4, 3, 0},
		{49, 64, 81, 64, 17, 16, 10, 9, 4, 3},
		{64, 49, 64, 81, 49, 17, 16, 10, 9, 4},
		{64, 16, 17, 49, 81, 64, 49, 17, 16, 10},
		{49, 10, 16, 17, 64, 81, 64, 49, 17, 16},
		{17, 9, 10, 16, 49, 64, 81, 64, 49, 17},
		{16, 4, 9, 10, 17, 49, 64, 81, 64, 49},
		{10, 3, 4, 9, 16, 17, 49, 64, 81, 64},
		{9, 0, 3, 4, 10, 16, 17, 49, 64, 81},
	}
	known := IntVector{1, 2, 3, 0, 4, 5, 6, 7, 8, 9}
	n := len(a)
	dataType := "similarity"
	energyFn := "Psi"
	improLagMax := 10000
	initPermute := true
	permuteCols := false
	rowPerm := NewIntVector(n)
	for i := 0; i < n; i++ {
		rowPerm[i] = i
	}
	colPerm := NewIntVector(n)
	for i := 0; i < n; i++ {
		colPerm[i] = i
	}

	e, bestRowPerm, _ := Boltzmann(a, dataType, rowPerm, colPerm, energyFn, improLagMax, initPermute, permuteCols)
	bestRowPerm.Increasing()
	if !bestRowPerm.IsIdentical(known) {
		fmt.Println("Best permutation: ")
		bestRowPerm.Print()
		fmt.Println("should be: ")
		known.Print()
		fmt.Println("Energy: ", e)
		t.Error()
	}
}

func TestBoltzmann_3(t *testing.T) {
	fmt.Println("Testing Boltzmann: perfect columnwise Q-matrix: Psi")
	var seed int64 = 7
	rand.Seed(seed)
	b := [][]int{ // perfect columnwise Q-matrix 15*20
		{0, 0, 0, 1357, 0, 0, 0, 38, 3, 0, 0, 0, 4, 0, 0, 0, 1, 0, 1141, 0},
		{0, 1, 0, 1141, 0, 0, 0, 322, 0, 0, 0, 0, 63, 0, 0, 0, 8, 0, 768, 3},
		{0, 8, 0, 768, 3, 0, 0, 888, 0, 0, 0, 0, 276, 1, 0, 0, 29, 0, 413, 15},
		{0, 29, 0, 413, 15, 0, 2, 817, 0, 0, 0, 0, 358, 4, 1, 0, 81, 0, 177, 49},
		{0, 81, 0, 177, 49, 0, 4, 251, 0, 0, 0, 0, 137, 12, 7, 0, 172, 0, 61, 127},
		{0, 172, 0, 61, 127, 0, 4, 25, 0, 3, 0, 0, 15, 29, 26, 0, 280, 0, 16, 257},
		{0, 280, 0, 16, 257, 0, 3, 0, 0, 14, 0, 0, 0, 64, 78, 0, 351, 0, 3, 408},
		{0, 351, 0, 3, 408, 0, 1, 0, 0, 43, 0, 0, 0, 119, 185, 0, 340, 0, 0, 510},
		{0, 340, 0, 0, 510, 0, 0, 0, 0, 97, 0, 0, 0, 194, 354, 0, 253, 0, 0, 502},
		{0, 253, 0, 0, 502, 1, 0, 0, 0, 163, 1, 0, 0, 275, 547, 0, 145, 0, 0, 389},
		{0, 145, 0, 0, 389, 6, 0, 0, 0, 206, 5, 0, 0, 339, 683, 1, 64, 0, 0, 238},
		{1, 64, 0, 0, 238, 29, 0, 0, 0, 196, 14, 0, 0, 364, 690, 3, 21, 0, 0, 114},
		{3, 21, 0, 0, 114, 98, 0, 0, 0, 139, 26, 98, 0, 340, 563, 4, 5, 0, 0, 43},
		{4, 5, 0, 0, 43, 243, 0, 0, 0, 74, 33, 280, 0, 276, 372, 3, 1, 0, 0, 12},
		{3, 1, 0, 0, 12, 450, 0, 0, 0, 29, 29, 10, 0, 196, 198, 1, 0, 0, 0, 3},
	}
	known := IntVector{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	n := len(b)
	m := len(b[0])
	dataType := "adjacency"
	energyFn := "Psi"
	improLagMax := 100000
	initPermute := true
	permuteCols := false
	rowPerm := NewIntVector(n)
	for i, _ := range rowPerm {
		rowPerm[i] = i
	}
	colPerm := NewIntVector(m)
	for i, _ := range colPerm {
		colPerm[i] = i
	}
	e, bestRowPerm, _ := Boltzmann(b, dataType, rowPerm, colPerm, energyFn, improLagMax, initPermute, permuteCols)
	bestRowPerm.Increasing()
	if !bestRowPerm.IsIdentical(known) {
		fmt.Println("Best permutation: ")
		bestRowPerm.Print()
		fmt.Println("should be: ")
		known.Print()
		fmt.Println("Energy: ", e)
		t.Error()
	}

	fmt.Println("Testing Boltzmann: perfect columnwise Q-matrix: Psi2")
	energyFn = "Psi2"
	e, bestRowPerm, _ = Boltzmann(b, dataType, rowPerm, colPerm, energyFn, improLagMax, initPermute, permuteCols)
	bestRowPerm.Increasing()
	if !bestRowPerm.IsIdentical(known) {
		fmt.Println("Best permutation: ")
		bestRowPerm.Print()
		fmt.Println("should be: ")
		known.Print()
		fmt.Println("Energy: ", e)
		t.Error()
	}

	fmt.Println("Testing Boltzmann: perfect columnwise Q-matrix: Psi3")
	energyFn = "Psi3"
	e, bestRowPerm, _ = Boltzmann(b, dataType, rowPerm, colPerm, energyFn, improLagMax, initPermute, permuteCols)
	bestRowPerm.Increasing()
	if !bestRowPerm.IsIdentical(known) {
		fmt.Println("Best permutation: ")
		bestRowPerm.Print()
		fmt.Println("should be: ")
		known.Print()
		fmt.Println("Energy: ", e)
		t.Error()
	}
}

/*
	// perfect Anti-Robinson matrix
		{0, 3, 4, 9, 10, 16, 17, 49, 64, 81},
		{3, 0, 3, 4, 9, 10, 16, 17, 49, 64},
		{4, 3, 0, 3, 4, 9, 10, 16, 17, 49},
		{9, 4, 4, 0, 3, 4, 9, 10, 16, 17},
		{10, 9, 4, 3, 0, 3, 4, 9, 10, 16},
		{16, 10, 9, 4, 3, 0, 3, 4, 9, 10},
		{17, 16, 10, 9, 4, 3, 0, 3, 4, 9},
		{49, 17, 16, 10, 9, 4, 3, 0, 3, 4},
		{64, 49, 17, 16, 10, 9, 4, 3, 0, 3},
		{81, 64, 49, 17, 16, 10, 9, 4, 3, 0},

	// not-perfect Anti-Robinson matrix
	{0, 5, 4, 9, 81, 16, 15, 49, 64, 20},
	{5, 0, 3, 4, 9, 25, 90, 36, 49, 64},
	{4, 3, 0, 4, 5, 9, 16, 12, 36, 49},
	{9, 4, 4, 0, 6, 4, 9, 16, 25, 36},
	{81, 9, 5, 6, 0, 4, 4, 9, 16, 25},
	{16, 25, 9, 4, 4, 0, 5, 4, 9, 16},
	{15, 90, 16, 9, 4, 5, 0, 3, 4, 9},
	{49, 36, 12, 16, 9, 4, 3, 0, 5, 4},
	{64, 49, 36, 25, 16, 9, 4, 5, 0, 4},
	{20, 64, 49, 36, 25, 16, 9, 4, 4, 0},


	// perfect Robinson matrix
		{81,64, 49, 17, 16, 10, 9, 4, 3, 0},
		{64, 81,64, 49, 17, 16, 10, 9, 4, 3},
		{49, 64, 81,64, 49, 17, 16, 10, 9, 4},
		{17, 49, 64, 81,64, 49, 17, 16, 10, 9},
		{16, 17, 49, 64, 81,64, 49, 17, 16, 10},
		{10, 16, 17, 49, 64, 81,64, 49, 17, 16},
		{9, 10, 16, 17, 49, 64, 81,64, 49, 17},
		{4, 9, 10, 16, 17, 49, 64, 81,64, 49},
		{3, 4, 9, 10, 16, 17, 49, 64, 81,64},
		{0, 3, 4, 9, 10, 16, 17, 49, 64, 81},

		// perfect columnwise Q-matrix 15*20
		{0,0,0,1357,0,0,0,38,3,0,0,0,4,0,0,0,1,0,1141,0},
		{0,1,0,1141,0,0,0,322,0,0,0,0,63,0,0,0,8,0,768,3},
		{0,8,0,768,3,0,0,888,0,0,0,0,276,1,0,0,29,0,413,15},
		{0,29,0,413,15,0,2,817,0,0,0,0,358,4,1,0,81,0,177,49},
		{0,81,0,177,49,0,4,251,0,0,0,0,137,12,7,0,172,0,61,127},
		{0,172,0,61,127,0,4,25,0,3,0,0,15,29,26,0,280,0,16,257},
		{0,280,0,16,257,0,3,0,0,14,0,0,0,64,78,0,351,0,3,408},
		{0,351,0,3,408,0,1,0,0,43,0,0,0,119,185,0,340,0,0,510},
		{0,340,0,0,510,0,0,0,0,97,0,0,0,194,354,0,253,0,0,502},
		{0,253,0,0,502,1,0,0,0,163,1,0,0,275,547,0,145,0,0,389},
		{0,145,0,0,389,6,0,0,0,206,5,0,0,339,683,1,64,0,0,238},
		{1,64,0,0,238,29,0,0,0,196,14,0,0,364,690,3,21,0,0,114},
		{3,21,0,0,114,98,0,0,0,139,26,98,0,340,563,4,5,0,0,43},
		{4,5,0,0,43,243,0,0,0,74,33,280,0,276,372,3,1,0,0,12},
		{3,1,0,0,12,450,0,0,0,29,29,10,0,196,198,1,0,0,0,3},

*/
