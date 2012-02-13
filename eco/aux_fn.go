// Auxiliary functions

package eco

import (
	"fmt"
	//	"encoding/csv"
	"math"
	"os"
)
/*
type Vector struct {
	A []float64 // data
	L int       // length
}

func NewVector(length int) (v *Vector) {
	v = new(Vector)
	v.L = length
	v.A = make([]float64, length)
	return v
}
*/

func ReadMatrix() *Matrix {
	var (
		rows, cols int
		x          float64
		data       *Matrix
	)

	fmt.Scanf("%d", &rows)
	fmt.Scanf("%d", &cols)

	data = NewMatrix(rows, cols)
	for i := 0; i < rows; i++ {
		for j := i + 1; j < cols; j++ {
			fmt.Scanf("%f", &x)
			data.Set(i, j, x)
		}
	}
	return data
}

/*
// Reads CSV from stdin and dumps it back to stdout.
func ReadMatrixCSV() *Matrix {
	var (
		rows, cols int
		x          float64
		data       *Matrix
	)

        table, err := csv.ReadAll(os.Stdin)
        if err != nil {
                panic(err.String())
        }
        // turn table to matrix... to be implemented
	return data
}
*/

func TruncData(data *Matrix) {
	rows := data.R
	cols := data.C
	eps := 1e-3
	warning := false

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			newX := math.Floor(data.Get(i, j))
			if !warning {
				if x-newX > eps {
					warning = true
				}
			}
			data.Set(i, j, newX) // truncate data, anyway
		}
	}
	if warning {
		fmt.Fprint(os.Stderr, "warning: data were not integers, so truncated\n")
	}
	return

}

func WarnIfNotBool(data *Matrix) {
	rows := data.R
	cols := data.C
	eps := 1e-3
	warning := false

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			newX := math.Floor(data.Get(i, j))
			if !warning {
				if x-newX > eps {
					warning = true
				}
				if newX != 0 {
					newX = 1
				}
			}
			data.Set(i, j, newX) // truncate data, anyway
		}
	}
	if warning {
		fmt.Fprint(os.Stderr, "warning: data are not 0/1, however, will be treated as boolean: 0 == false, otherwise true\n")
	}
	return

}

func WarnIfNotCounts(data *Matrix) {
	rows := data.R
	cols := data.C
	eps := 1e-3
	warning := false
	warning2 := true

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			newX := math.Floor(data.Get(i, j))
			if !warning {
				if x-newX > eps {
					warning = true
				}
			}
			if x != 0 &&  x != 0 {
				warning2 = false
			}
			data.Set(i, j, newX) // truncate data, anyway
		}
	}
	if warning2 {
		fmt.Fprint(os.Stderr, "warning: data seem to be boolean, not counts as required\n")
	}

	if warning {
		fmt.Fprint(os.Stderr, "warning: data are not counts, will be truncated to integers\n")
	}
	return

}

func WarnIfDblNewMatrix(data *Matrix) {
	rows := data.R
	cols := data.C
	warning := false
L:
	for j := 0; j < cols; j++ {
		colSum := 0
		for i := 0; i < rows; i++ {
			if data.Get(i, j) > 0.0 {
				colSum++
			}
		}
		if colSum == 0 {
			warning = true
			break L
		}
	}
	if warning {
		fmt.Fprint(os.Stderr, "warning: data have empty species which influence the results\n")
	}
	return
}

// Calculates A, B, J, and P values from two rows of boolean data matrix, "quadratic variant". 
// See R:vegan:vegdist
// "quadratic" terms are J = sum(x*y), A = sum(x^2), B = sum(y^2)
func GetABJPquad(data *Matrix, i, j int) (aa, bb, jj, pp float64) {
	cols := data.C

	jj = 0.0
	aa = 0.0
	bb = 0.0
	for k := 0; k < cols; k++ {
		x := data.Get(i, k)
		y := data.Get(j, k)
		jj += x * y
		aa += x * x
		bb += y * y
	}
	pp = float64(cols)
	return
}

// Calculates A, B, J, and P values from two rows of boolean data matrix, "minimum variant". 
// See R:vegan:vegdist
// "minimum" terms are J = sum(min(x,y)), A = sum(x) and B = sum(y)
func GetABJPmin(data *Matrix, i, j int) (aa, bb, jj, pp float64) {
	cols := data.C

	jj = 0.0
	aa = 0.0
	bb = 0.0
	for k := 0; k < cols; k++ {
		x := data.Get(i, k)
		y := data.Get(j, k)
		jj += math.Min(x, y)
		aa += x
		bb += y
	}
	pp = float64(cols)
	return
}

// Calculates A, B, J, and P values from two rows of boolean data matrix. 
// See R:vegan:vegdist
func GetABJPbool(data *Matrix, i, j int) (aa, bb, jj, pp float64) {
	cols := data.C

	jj = 0.0
	t1 := 0
	t2 := 0
	for k := 0; k < cols; k++ {
		x := data.Get(i, k)
		y := data.Get(j, k)

		if x > 0.0 && y > 0.0 {
			jj++
		}
		if x > 0.0 {
			t1++
		}
		if y > 0.0 {
			t2++
		}
	}
	if t1 < t2 {
		aa = float64(t1)
		bb = float64(t2)

	} else {
		aa = float64(t2)
		bb = float64(t1)
	}
	pp = float64(cols)
	return
}

// Calculates similarity matrix from the distance matrix. 
func SFromD(dis *Matrix, which int) *Matrix {
	rows := dis.R
	out := NewMatrix(rows, rows)

	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			v := dis.Get(i, j)
			switch {
			case which == 0:	// D is from [0, 1]
				v = 1 - v
			case which == 1:	// D is from [0, +inf]
				v = 1 / (v + 1)
			}
			out.Set(i, j, v)
		}
	}
	return out
}

// Calculates distance matrix from the similarity matrix
func DFromS(sim *Matrix, which int) *Matrix {
	rows := sim.R
	out := NewMatrix(rows, rows)

	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			v := sim.Get(i, j)
			switch {
			case which == 0:	// S is from [0, 1]
				v = 1 - v
			case which == 1:	// S is from [0, 1]
				v = math.Sqrt(1 - v)
			case which == 2:	// S is from [0, 1]
				v = math.Sqrt(1 - v*v)
			case which == 3:	// S is from [0, +inf]
				v = 1/v - 1
			}
			out.Set(i, j, v)
		}
	}
	return out
}

// recalculates data matrix to proportions. 
func RecalcToProp(data *Matrix) int {
	rows := data.R
	cols := data.C
	// calculate row sums (sum of all abundances of all species in the sample)
	for i := 0; i < rows; i++ {
		rowSum := 0.0
		for j := 0; j < cols; j++ {
			rowSum += data.Get(i, j)
		}
		// recalculate to proportiond
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			data.Set(i, j, x/rowSum)
		}
	}
	return 0
}

// For boolean data, calculates a, b, c, d values of the contingency table. 
// To be use for calculation of similarity indices. 
func GetABCD(data *Matrix, row1, row2 int) (a, b, c, d float64) {
	cols := data.C

	WarnIfNotBool(data)

	a = 0
	b = 0
	c = 0
	d = 0

	for k := 0; k < cols; k++ {
		x := data.Get(row1, k)
		y := data.Get(row2, k)

		switch {
		case x != 0 && y != 0:
			a++
		case x != 0 && y == 0:
			b++
		case x == 0 && y != 0:
			c++
		case x == 0 && y == 0:
			d++
		}
	}
	return a, b, c, d
}
