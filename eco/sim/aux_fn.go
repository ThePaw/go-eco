package eco

import (
	"fmt"
//	"encoding/csv"
	. "gomatrix.googlecode.com/hg/matrix"
	"math"
	"os"
)

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


func ReadMatrix() *DenseMatrix {
	var (
		rows, cols int
		x          float64
		data       *DenseMatrix
	)

	fmt.Scanf("%d", &rows)
	fmt.Scanf("%d", &cols)

	data = Zeros(rows, cols)
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
func ReadMatrixCSV() *DenseMatrix {
	var (
		rows, cols int
		x          float64
		data       *DenseMatrix
	)

        table, err := csv.ReadAll(os.Stdin)
        if err != nil {
                panic(err.String())
        }
        // turn table to matrix... to be implemented
	return data
}
*/

func truncData(data *DenseMatrix) {
	rows := data.Rows()
	cols := data.Cols()
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

func warnIfNotBool(data *DenseMatrix) {
	rows := data.Rows()
	cols := data.Cols()
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
		fmt.Fprint(os.Stderr, "warning: data are 0/1, however, treated as boolean: 0 == false, otherwise true\n")
	}
	return

}

func warnIfDblZeros(data *DenseMatrix) {
	rows := data.Rows()
	cols := data.Cols()
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
func getABJPquad(data *DenseMatrix, i, j int) (aa, bb, jj, pp float64) {
	cols := data.Cols()

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
func getABJPmin(data *DenseMatrix, i, j int) (aa, bb, jj, pp float64) {
	cols := data.Cols()

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
func getABJPbool(data *DenseMatrix, i, j int) (aa, bb, jj, pp float64) {
	cols := data.Cols()

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
func sFromD(dis *DenseMatrix, which int) *DenseMatrix {
	rows := dis.Rows()
	out := Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			v := dis.Get(i, j)
			switch {
			case which == 0:
				v = 1 - v
			case which == 1:
				v = 1 / (v + 1)
			}
			out.Set(i, j, v)
		}
	}
	return out
}

// Calculates distance matrix from the similarity matrix
func dFromS(sim *DenseMatrix, which int) *DenseMatrix {
	rows := sim.Rows()
	out := Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			v := sim.Get(i, j)
			switch {
			case which == 0:
				v = 1 - v
			case which == 1:
				v = 1/v - 1
			}
			out.Set(i, j, v)
		}
	}
	return out
}

// recalculates data matrix to proportions. 
func recalcToProp(data *DenseMatrix) int {
	rows := data.Rows()
	cols := data.Cols()
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
func getABCD(data *DenseMatrix, row1, row2 int) (a, b, c, d float64) {
	cols := data.Cols()

	warnIfNotBool(data)

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
