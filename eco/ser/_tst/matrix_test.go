package ser

import (
//	. "code.google.com/p/ser"
	"fmt"
	"testing"
)

func TestMatrixRearrangement1(t *testing.T) {
	fmt.Println("Test IntMatrix: RearrangeRows()")
	matrix := IntMatrix{
		{11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25, 26, 27, 28, 29, 30},
		{31, 32, 33, 34, 35, 36, 37, 38, 39, 40},
		{41, 42, 43, 44, 45, 46, 47, 48, 49, 50},
		{51, 52, 53, 54, 55, 56, 57, 58, 59, 60},
		{61, 62, 63, 64, 65, 66, 67, 68, 69, 70},
		{71, 72, 73, 74, 75, 76, 77, 78, 79, 80},
		{81, 82, 83, 84, 85, 86, 87, 88, 89, 90},
		{91, 92, 93, 94, 95, 96, 97, 98, 99, 100},
	}
	perm := IntVector{0, 1, 2, 4, 3, 5, 6, 7, 8}

	matrix.RearrangeRows(perm)
	matrix.WriteGo()
}

func TestMatrixRearrangement2(t *testing.T) {
	fmt.Println("Test IntMatrix: RearrangeCols()")
	matrix := IntMatrix{
		{11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25, 26, 27, 28, 29, 30},
		{31, 32, 33, 34, 35, 36, 37, 38, 39, 40},
		{41, 42, 43, 44, 45, 46, 47, 48, 49, 50},
		{51, 52, 53, 54, 55, 56, 57, 58, 59, 60},
		{61, 62, 63, 64, 65, 66, 67, 68, 69, 70},
		{71, 72, 73, 74, 75, 76, 77, 78, 79, 80},
		{81, 82, 83, 84, 85, 86, 87, 88, 89, 90},
		{91, 92, 93, 94, 95, 96, 97, 98, 99, 100},
	}
	perm := IntVector{0, 1, 2, 4, 3, 5, 6, 7, 8, 9}

	matrix.RearrangeCols(perm)
	matrix.WriteGo()
}

/*
func TestMatrixRearrangement3(t *testing.T) {
	fmt.Println("Test IntMatrix: Rearrange()")
	matrix := IntMatrix{
		{11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25, 26, 27, 28, 29, 30},
		{31, 32, 33, 34, 35, 36, 37, 38, 39, 40},
		{41, 42, 43, 44, 45, 46, 47, 48, 49, 50},
		{51, 52, 53, 54, 55, 56, 57, 58, 59, 60},
		{61, 62, 63, 64, 65, 66, 67, 68, 69, 70},
		{71, 72, 73, 74, 75, 76, 77, 78, 79, 80},
		{81, 82, 83, 84, 85, 86, 87, 88, 89, 90},
		{91, 92, 93, 94, 95, 96, 97, 98, 99, 100},
	}
//	rowPerm := IntVector{0, 1, 2, 4, 3, 5, 6, 7, 8}
//	colPerm := IntVector{0, 1, 2, 4, 3, 5, 6, 7, 8, 9}
	v := matrix.Clone()	

	rowPerm := NewIntVector(matrix.Rows())
	rowPerm.Perm()
	rowPerm.Print()

	colPerm := NewIntVector(matrix.Cols())
	colPerm.Perm()
	colPerm.Print()


	matrix.Rearrange(rowPerm, colPerm)
	matrix.WriteGo()
	w := matrix.Clone()	
	w.Rearrange(rowPerm, colPerm)
	if !w.IsEqual(v) {
		fmt.Println("Equality test #3 failed")
		v.Print()
		fmt.Println()
		w.Print()
		t.Error()
	}
}
*/

func TestIsEqual(t *testing.T) {
	fmt.Println("Test IntMatrix: IsEqual()")
	// perfect Robinson matrix
	a := IntMatrix{
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
	if !a.IsEqual(a) {
		fmt.Println("Equality test #1 failed")
		t.Error()
	}
}

func TestIsEqual2(t *testing.T) {
	fmt.Println("Test IntMatrix: IsEqual2()")
	// perfect Robinson matrix
	a := IntMatrix{
		{41, 42, 43, 44, 45, 46, 47, 48, 49, 50},
		{51, 52, 53, 54, 55, 56, 57, 58, 59, 60},
		{61, 62, 63, 64, 65, 66, 67, 68, 69, 70},
		{71, 72, 73, 74, 75, 76, 77, 78, 79, 80},
		{81, 82, 83, 84, 85, 86, 87, 88, 89, 90},
		{91, 92, 93, 94, 95, 96, 97, 98, 99, 100},
	}
	if !a.IsEqual(a) {
		fmt.Println("Equality test #2 failed")
		t.Error()
	}
}

func TestMatrixRearrangement4(t *testing.T) {
	fmt.Println("Test IntMatrix: Rearrange() using RobFAntK()")
	// perfect Robinson matrix
	sim := IntMatrix{
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
	//	rowPerm:=IntVector{0, 1, 2, 4, 3, 5, 9, 7, 8, 6}

	rowPerm := NewIntVector(sim.Rows())
	rowPerm.Perm()
	rowPerm.Print()
	sim.Rearrange(rowPerm, rowPerm)

	trials := 2
	improLagMax := 60
	r := 5
	_, backPerm := RobFAntK(sim, trials, improLagMax, r)
	sim.Rearrange(backPerm, backPerm)
	if !sim.IsR() {
		sim.WriteGo()
		t.Error()
	}
}
