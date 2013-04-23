package ser

import (
	"fmt"
	"testing"
)

func TestIRound(t *testing.T) {
x:= iRound(4.45)
	fmt.Println(x, " should be 4")
x= iRound(4.51)
	fmt.Println(x, " should be 5")
}
