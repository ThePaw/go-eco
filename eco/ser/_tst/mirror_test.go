package ser

import (
	"fmt"
	"math"
	"testing"
)

func TestMirrorLoss(t *testing.T) {
	fmt.Println("TestMirrorLoss() #1")
	mtx := Matrix64{
		// perfect columnwise Q-matrix 15*20
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
	tol := 1e-4

	pRow := IntVector{14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	pCol := IntVector{18, 3, 7, 12, 17, 6, 2, 15, 0, 8, 10, 11, 5, 9, 13, 14, 4, 19, 1, 16}
	known := 32089.0
	loss := MirrorLoss(mtx, pRow, pCol, false)
	if math.Abs(loss-known) > tol {
		fmt.Println("loss:  ", loss, "known :", known)
		t.Error()
	}

	fmt.Println("TestMirrorLoss() #2")
	pRow = IntVector{14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	pCol = IntVector{18, 3, 7, 12, 2, 6, 0, 17, 15, 8, 10, 11, 5, 9, 13, 14, 4, 19, 1, 16}
	known = 32105.0
	loss = MirrorLoss(mtx, pRow, pCol, false)
	if math.Abs(loss-known) > tol {
		fmt.Println("loss:  ", loss, "known :", known)
		t.Error()
	}

	fmt.Println("TestMirrorLoss() #3")
	pRow = IntVector{12, 4, 2, 13, 14, 8, 5, 7, 11, 3, 0, 1, 9, 10, 6}
	pCol = IntVector{14, 15, 11, 0, 8, 6, 5, 10, 3, 7, 2, 18, 4, 19, 12, 9, 17, 16, 1, 13}
	known = 75039.0
	loss = MirrorLoss(mtx, pRow, pCol, false)
	if math.Abs(loss-known) > tol {
		fmt.Println("loss:  ", loss, "known :", known)
		t.Error()
	}

	fmt.Println("TestMirrorLoss() #4")
	pRow = IntVector{12, 4, 2, 10, 9, 13, 14, 8, 5, 6, 7, 11, 3, 0, 1}
	pCol = IntVector{14, 15, 11, 0, 8, 12, 2, 18, 16, 1, 4, 19, 9, 6, 5, 10, 3, 7, 17, 13}
	known = 67922.0
	loss = MirrorLoss(mtx, pRow, pCol, false)
	if math.Abs(loss-known) > tol {
		fmt.Println("loss:  ", loss, "known :", known)
		t.Error()
	}

	fmt.Println("TestMirrorLoss() #5")
	pRow = IntVector{1, 0, 3, 2, 8, 5, 6, 7, 9, 10, 11, 12, 13, 14, 4}
	pCol = IntVector{13, 17, 7, 3, 11, 0, 8, 18, 2, 12, 16, 1, 4, 19, 9, 6, 5, 10, 15, 14}
	known = 59182.0
	loss = MirrorLoss(mtx, pRow, pCol, false)
	if math.Abs(loss-known) > tol {
		fmt.Println("loss:  ", loss, "known :", known)
		t.Error()
	}
}
