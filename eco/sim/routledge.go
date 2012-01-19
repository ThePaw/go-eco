// Routledge similarity matrices
// Routledge (1977), Magurran (1988), Wilson & Shmida (1984)


package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	"math"
)

// Routledge similarity matrix #1
// Routledge (1977), Magurran (1988)
func Routledge1Bool_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim           *DenseMatrix
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	sim = Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = getABCD(data, i, j)
			abc2 := (a+b+c) * (a+b+c)
			s:= abc2/(abc2 - 2*b*c) - 1
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}

// Routledge similarity matrix #2
// Routledge (1977), Wilson & Shmida (1984)
func Routledge2Bool_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim           *DenseMatrix
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	sim = Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = getABCD(data, i, j)
			s:= math.Log(2*a+b+c)-((1/(2*a+b+c))*2*a*math.Log(2))-((1/(2*a+b+c))*((a+b)*math.Log(a+b)+(a+c)*math.Log(a+c))) 
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}

// Routledge similarity matrix #3
// Routledge (1977)
func Routledge3Bool_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim           *DenseMatrix
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	sim = Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = getABCD(data, i, j)
			v:= math.Log(2*a+b+c)-((1/(2*a+b+c))*2*a*math.Log(2))-((1/(2*a+b+c))*((a+b)*math.Log(a+b)+(a+c)*math.Log(a+c))) 
			s := math.Exp(v) -1
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}


