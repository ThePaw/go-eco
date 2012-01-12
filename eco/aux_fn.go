package eco

import (
	"fmt"
	. "gomatrix.googlecode.com/hg/matrix"
	"math"
	"os"
)

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

func checkIfBool(data *DenseMatrix) {
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
