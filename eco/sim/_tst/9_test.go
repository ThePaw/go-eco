package sim

import (
	"fmt"
	//	"code.google.com/p/go-eco/eco/aux"
	"testing"
)

func TestHornMorisitaBool(t *testing.T) {
	// HornMorisitaBool test 1 vs.2 
	fmt.Println("HornMorisitaBool test 1 vs.2 ")
	data := GetBoolData2()
	out1 := HornMorisita_S(data)
	out2 := MorisitaHorn2_S(data)
	rows := data.R

	// check
	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			if !check(out1.Get(i, j), out2.Get(i, j)) {
				t.Error()
				fmt.Println(i, j, out1.Get(i, j), out2.Get(i, j))
			}

		}
	}
}
