package ser

import (
	"fmt"
	"testing"
//	. "code.google.com/p/probab/dst"
	"code.google.com/p/probab/stat"
)

func TestA2D(t *testing.T) {
var α, θ float64
α = 800
θ = 1
nIter := 1000000
y:=NewVector64(nIter)
	for i := 0; i < nIter; i++ {

y[i] = postDensityFNext(int64(α) , θ)

}
μ , σ :=stat.SampleMeanVar(y)
fmt.Println("Mean, var: ",μ , σ)
}
