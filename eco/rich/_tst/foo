package rich

import (
	"fmt"
	"testing"
)

// Chao test against R:fossil
func TestChao(t *testing.T) {
	fmt.Println("Chao test against R:fossil")
	data := GetCounts()
	out := Chao(data)
	known := [...]float64{14.5, 7, 11.5}
	rows := data.R
	// check
	for i := 0; i < rows; i++ {
		x := out.Get(i)
		y := known[i]

		if !check(x, y) {
			t.Error()
			fmt.Println(i, x, y)
		}
	}
}

// ChaoBool test against R:fossil
func TestChaoBool(t *testing.T) {
	fmt.Println("ChaoBool test against R:fossil")
	data := GetBoolData2()
	out := Chao(data)
	known := [...]float64{21, 10, 6, 21, 21, 15}
	rows := data.R
	// check
	for i := 0; i < rows; i++ {
		x := out.Get(i)
		y := known[i]

		if !check(x, y) {
			t.Error()
			fmt.Println(i, x, y)
		}
	}
}

// ACE test against R:fossil
func TestACE(t *testing.T) {
	fmt.Println("ACE test against R:fossil")
	data := GetCounts()
	out := ACE(data)
	known := [...]float64{11.09091, 7, 8.2}
	rows := data.R
	// check
	for i := 0; i < rows; i++ {
		x := out.Get(i)
		y := known[i]

		if !check(x, y) {
			t.Error()
			fmt.Println(i, x, y)
		}
	}
}

// ICE test against R:fossil
func TestICE(t *testing.T) {
	fmt.Println("ICE test against R:fossil")
	data := GetBoolData2()
	out := ICE(data)
	known := [...]float64{21, 10, 6, 21, 21, 15}
	rows := data.R
	// check
	for i := 0; i < rows; i++ {
		x := out.Get(i)
		y := known[i]

		if !check(x, y) {
			t.Error()
			fmt.Println(i, x, y)
		}
	}
}
