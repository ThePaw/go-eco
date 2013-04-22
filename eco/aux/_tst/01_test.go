package aux

import (
	"fmt"
	"testing"
)

func TestOpenCsvFloat64(t *testing.T) {
	mtx := FetchCsvMatrix("matrix.dat")
	fmt.Println(mtx.Get(1, 3))
	mtx.Print()
}
