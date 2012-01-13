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

func sFromD(dis *DenseMatrix, which int) *DenseMatrix {
	var s float64
	rows := dis.Rows()
	sim := Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			d := dis.Get(i, j)
			switch {
			case which == 0:
				s = 1 - d
			case which == 1:
				s = 1 / (d + 1)
			}
			sim.Set(i, j, s)
		}
	}
	return sim
}

func dFromS(sim *DenseMatrix, which int) *DenseMatrix {
	var d float64
	rows := sim.Rows()
	dis := Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			s := sim.Get(i, j)
			switch {
			case which == 0:
				d = 1-s
			case which == 1:
				d = 1/s - 1
			}
			dis.Set(i, j, d)
		}
	}
	return dis
}

func recalcToProp(data *DenseMatrix) int{
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
			x:= data.Get(i, j)
			data.Set(i, j, x/rowSum)
		}
	}
	return 0
}

