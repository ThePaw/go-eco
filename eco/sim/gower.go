/* 
Gower distance and similarity
Gower is like Manhattan, but data are standardized to range 0..1
for rows and distance is divided by the number of pairs with both non-missing values. 

dis[jk] = (1/M) sum (abs(x[ij]-x[ik])/(max(x[i])-min(x[i]))
where M is the number of columns (excluding missing values) 

Gower, J. C. (1971), “A general coefficient of similarity and some of its properties”. Biometrics, 27, 623–637.
Kaufman, L. and Rousseeuw, P.J. (1990), Finding Groups in Data: An Introduction to Cluster Analysis. Wiley, New York. 
*/

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	. "math"
)

// Gower distance for interval-scaled variables
func Gower_D(data *DenseMatrix) *DenseMatrix {
	var (
		dis *DenseMatrix
	)
	const missing float64 = -999 //code for missing values

	rows := data.Rows()
	cols := data.Cols()
	dis = Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		dis.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sum := 0.0
			count := 0
			maxx := 0.0
			minx := 0.0

			// columns are considered as interval-scaled variables and 
			// d_ijk = abs(x_ik - x_jk) / R_k
			// where R_k is the range of the kth variable. 

			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				maxx = Max(x, maxx)
				maxx = Max(y, maxx)
				minx = Min(x, minx)
				minx = Min(y, minx)
				if x != missing && y != missing {
					r := maxx - minx
					d := Abs(x-y) / r
					sum += d
					count++
				}
			}
			d := sum / float64(count)
			dis.Set(i, j, d)
			dis.Set(j, i, d)
		}
	}
	return dis
}

// Gower similarity for interval-scaled variables
// If d denotes Gower distance, similarity is s=1.00/(d+1), so that it is in [0, 1]
func Gower_S(data *DenseMatrix) *DenseMatrix {
	var (
		rows     int
		sim, dis *DenseMatrix
	)

	rows = data.Rows()
	dis = Gower_D(data)
	sim = Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		sim.Set(i, i, 1.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			s := 1.00 / (dis.Get(i, j) + 1.0)
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}

// Gower distance for ordered variables
// If kr == true, the extension of the Gower's dissimilarity measure proposed by Kaufman and Rousseeuw (1990) is used. 
// Otherwise, the original Gower's (1971) dissimilarity is considered. 
func GowerOrd_D(data *DenseMatrix, kr bool) *DenseMatrix {
	var (
		dis *DenseMatrix
	)
	const missing float64 = -999 //code for missing values

	rows := data.Rows()
	cols := data.Cols()
	dis = Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		dis.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sum := 0.0
			count := 0
			maxx := 0.0
			minx := 0.0

			// columns are considered as categorical ordinal variables and the values are substituted 
			// with the corresponding position index, r_ik in the factor levels. 
			// These position indexes (that are different from the output of the R function rank) are transformed in the following manner: 
			// z_ik = (r_ik - 1)/(max(r_ik) - 1)
			// These new values, z_ik, are treated as observations of an interval scaled variable. 

			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				maxx = Max(x, maxx)
				maxx = Max(y, maxx)
				minx = Min(x, minx)
				minx = Min(y, minx)
				if x != missing && y != missing {
					r := maxx - 1
					if kr {
						x = (x - 1) / r
						y = (y - 1) / r
					}
					d := Abs(x-y) / r
					sum += d
					count++
				}
			}
			d := sum / float64(count)
			dis.Set(i, j, d)
			dis.Set(j, i, d)
		}
	}
	return dis
}

// Gower similarity for ordered variables
// If d denotes Gower distance, similarity is s=1.00/(d+1), so that it is in [0, 1]
func GowerOrd_S(data *DenseMatrix) *DenseMatrix {
	var (
		rows     int
		sim, dis *DenseMatrix
	)

	rows = data.Rows()
	dis = Gower_D(data)
	sim = Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		sim.Set(i, i, 1.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			s := 1.00 / (dis.Get(i, j) + 1.0)
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}

// Gower distance for boolean data
func GowerBool_D(data *DenseMatrix) *DenseMatrix {
	var (
		dis        *DenseMatrix
		a, b, c, d int64
	)

	rows := data.Rows()
	cols := data.Cols()
	dis = Zeros(rows, rows)
	a = 0
	b = 0
	c = 0
	d = 0

	checkIfBool(data)

	for i := 0; i < rows; i++ {
		dis.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)

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
			d := float64(b+c) / float64(cols)
			dis.Set(i, j, d)
			dis.Set(j, i, d)
		}
	}
	return dis
}

// Citation needed
func GowerZBool_D(data *DenseMatrix) *DenseMatrix {
	var (
		dis        *DenseMatrix
		a, b, c, _ float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	dis = Zeros(rows, rows)
	checkIfBool(data)

	for i := 0; i < rows; i++ {
		dis.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			a, b, c, _ = getABCD(data, i, j)
			d := (b+c) / (a+b+c)
			dis.Set(i, j, d)
			dis.Set(j, i, d)
		}
	}
	return dis
}

// Gower similarity for boolean data
// Gower & Legendre (1986)
func GowerBool_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim           *DenseMatrix
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	sim = Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = getABCD(data, i, j)
			s:= (a - (b+c)+d) / (a+b+c+d)
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}





