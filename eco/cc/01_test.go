package cc

import (
	"fmt"
	//	"code.google.com/p/go-eco/eco/aux"
	"testing"
)

func TestGenBeta(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(GenBetaSRF(float64(i), 50.0, 50.0, 4.0, 2.0))
	}
}
