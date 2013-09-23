package ser

import (
	"encoding/csv"
	//	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

// Matrix64 is 2D array of float64. Elements are stored in a single float64 slice and slices of each row are created.

type Matrix64 [][]float64

// NewMatrix64 creates a new IntMatrix instance with specified number of rows and columns.
func NewMatrix64(nRow, nCol int) Matrix64 {
	s := make([]float64, nCol*nRow)
	a := make(Matrix64, nRow)
	for i, p := 0, 0; i < nRow; i++ {
		a[i] = s[p : p+nCol]
		p += nCol
	}
	return a
}

// Dims return the dimensions of the matrix.
func (a Matrix64) Dims() (nRow, nCol int) {
	return len(a), len(a[0])
}

// Rows returns the number of rows. 
func (a Matrix64) Rows() int {
	return len(a)
}

// Cols returns the number of columns. 
func (a Matrix64) Cols() int {
	if len(a) == 0 {
		return 0
	}
	return len(a[0])
}

/*
// CopyTo copies matrix to an existing matrix. 
func (a Matrix64) CopyTo(targetMat Matrix64) {
	n := a.Rows() * a.Cols()
	if n > 0 {
		copy(targetMat[0][:n], a[0][:n])
	}
	return
}

// CopyFrom copies matrix from an existing matrix.
func (a Matrix64) CopyFrom(srcMat Matrix64) {
	n := a.Rows() * a.Cols()
	if n > 0 {
		copy(m, srcMat)
	}
	return
}
*/

// CopyTo copies matrix to an existing matrix. 
func (a Matrix64) CopyTo(targetMat Matrix64) {
	for i, row := range a {
		for j, _ := range row {
			targetMat[i][j] = a[i][j]
		}
	}
	return
}

// CopyFrom copies matrix from an existing matrix.
func (a Matrix64) CopyFrom(srcMat Matrix64) {
	for i, row := range a {
		for j, _ := range row {
			a[i][j] = srcMat[i][j]
		}
	}
	return
}

// CopyFromInt copies matrix from an existing integer matrix.
func (a Matrix64) CopyFromInt(srcMat IntMatrix) {
	for i, row := range a {
		for j, _ := range row {
			a[i][j] = float64(srcMat[i][j])
		}
	}
	return
}

// Clone clones a matrix.
func (a Matrix64) Clone() Matrix64 {
	clone := NewMatrix64(a.Dims())
	for i, row := range a {
		for j, val := range row {
			clone[i][j] = val
		}
	}
	return clone
}

// Swap rows i, j
func (a Matrix64) SwapRows(i int, j int) {
	nCol := a.Cols()
	for k := 0; k < nCol; k++ {
		x := a[i][k]
		a[i][k] = a[j][k]
		a[j][k] = x
	}
}

// Swap columns i, j
func (a Matrix64) SwapCols(i int, j int) {
	nRow := a.Rows()
	for k := 0; k < nRow; k++ {
		x := a[k][i]
		a[k][i] = a[k][j]
		a[k][j] = x
	}
}

// ReadCsvMatrix64  reads the matrix from an opened CSV file. There must be no spaces between values, just commas.
func ReadCsvMatrix64(f *os.File) (a Matrix64) {
	read := csv.NewReader(io.Reader(f))
	data, err := read.ReadAll()
	if err != nil {
		panic("Failed to read from the CSV File(Maybe the file does not comply to the CSV standard defined in RFC 4180)")
	}
	nRow := len(data)
	nCol := len(data[0])
	a = NewMatrix64(nRow, nCol)
	for i, row := range a {
		for j, _ := range row {
			x, err := strconv.ParseFloat(data[i][j], 64)
			if err != nil {
				fmt.Print(err)
				panic("could not convert the string")
			}
			a[i][j] = float64(x)
		}
	}
	return
}

/*
func (m IntMatrix) WriteCSV(f *os.File)  {	// to be implemented
	write := csv.NewWriter(io.Writer(f))
	records := // [][]string TO BE IMPLEMENTED
	nRow, cols := a.Dims()
	for i := 0; i < nRow; i++ {
		for j := 0; j < cols; j++ {
			records[i][j] = strconv.FormatInt(a[i][j], 10)
		}
	}
	err :=WriteAll(records)
	if err != nil {
		fmt.Println("Failed to write the CSV File")
	}
}
*/

func (a Matrix64) WriteCSV() {
	for i, row := range a {
		for j, _ := range row {
			if j == 0 {
				fmt.Print(a[i][j])
			} else {
				fmt.Print(",", a[i][j])
			}
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
}

func (a Matrix64) WriteCSV3() {
	for i, row := range a {
		for j, _ := range row {
			if j == 0 {
				fmt.Printf("%6.3f", a[i][j])
			} else {
				fmt.Printf(",%6.3f", a[i][j])
			}
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
}

func (a Matrix64) WriteGo() {
	fmt.Println("matrix := Matrix64{")
	for i, row := range a {
		fmt.Print("{")
		for j, _ := range row {
			if j == 0 {
				fmt.Print(a[i][j])
			} else {
				fmt.Print(",", a[i][j])
			}
		}
		fmt.Println("},")
	}
	fmt.Println("}")
}

func (a Matrix64) WriteGo3() {
	fmt.Println("matrix := Matrix64{")
	for i, row := range a {
		fmt.Print("{")
		for j, _ := range row {
			if j == 0 {
				fmt.Printf("%6.3f", a[i][j])
			} else {
				fmt.Printf(",%6.3f", a[i][j])
			}
		}
		fmt.Println("},")
	}
	fmt.Println("}")
	fmt.Println()
}

func (a Matrix64) Print() {
	for i, row := range a {
		for j, _ := range row {
			fmt.Print(a[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println()
}

// PrettyString returns a pretty string form of the matrix
func (a Matrix64) PrettyString() string {
	sa := make([][]string, 0, a.Rows())
	for _, row := range a {
		sr := make([]string, 0, len(row))
		for _, cell := range row {
			sr = append(sr, fmt.Sprint(cell))
		} // for cell
		sa = append(sa, sr)
	} // for row

	wds := make([]int, a.Cols())
	for j := 0; j < a.Cols(); j++ {
		for i := 0; i < a.Rows(); i++ {
			if len(sa[i][j]) > wds[j] {
				wds[j] = len(sa[i][j])
			} //  if
		} // for i
	} // for i

	res := ""
	for i, row := range sa {
		if i == 0 {
			res += "["
		} else {
			res += " "
		} // else
		res += "["
		for j, cell := range row {
			if j > 0 {
				res += " "
			} // if
			res += fmt.Sprintf(fmt.Sprintf("%%%ds", wds[j]), cell)
		} // for j, cell
		res += "]"
		if i == len(sa)-1 {
			res += fmt.Sprintf("](%dx%d)", a.Rows(), a.Cols())
		} // else
		res += "\n"
	} // for row

	return res
}

// Product computes the product of two matrices. 
func (a Matrix64) Product(b Matrix64) Matrix64 {
	nRow, nCol := a.Dims()
	bRow, bCol := b.Dims()
	if nRow != bCol || nCol != bRow {
		panic("bad dimensions")
	}
	c := NewMatrix64(nRow, nRow)
	for i, row := range c {
		for j, _ := range c {
			c[i][j] = 0
			for k, _ := range row {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return c
}

// CircleProduct computes circular product of two matrices. 
func (a Matrix64) CircleProduct(b Matrix64) Matrix64 {
	r1, c1 := a.Dims()
	r2, c2 := b.Dims()
	if r1 != c2 || c1 != r2 {
		panic("bad dimensions")
	}
	c := NewMatrix64(r1, r1)
	for i, row := range a {
		for j, _ := range a {
			c[i][j] = 0
			for k, _ := range row { // columns of a
				c[i][j] += min(a[i][k], b[k][j])
			}
		}
	}
	return c
}

// CirclePower computes circular power of a matrix. 
// See Kendall, 1971: 111, for definition.
func (a Matrix64) CirclePower(n int) Matrix64 {
	// S0
	w := a.Clone()
	t := w.Transpose()
	s := w.CircleProduct(t)
	for i := 1; i < n; i++ { // S1 ... Sn
		t = s.Transpose()
		s = s.CircleProduct(t)
	}
	return s
}

// Transpose returns transposed matrix. 
func (a Matrix64) Transpose() Matrix64 {
	nRow, nCol := a.Dims()
	c := NewMatrix64(nCol, nRow)
	for i, row := range c {
		for j, _ := range row {
			c[i][j] = a[j][i]
		}
	}
	return c
}

// IsSquare tests whether the matrix is square matrix. 
func (a Matrix64) IsSquare() bool {
	q := true
	nRow, nCol := a.Dims()
	if nRow != nCol {
		q = false
	}
	return q
}

// IsSymmetric tests whether the square matrix is symmetric. 
func (a Matrix64) IsSymmetric() bool {
	if !a.IsSquare() {
		panic("not a square matrix")
	}
	nRow, nCol := a.Dims()
	q := true
	for i := 0; i < nRow && q; i++ {
		for j := 0; j < nCol && q; j++ {
			if a[i][j] != a[j][i] {
				fmt.Println("i, j, a[i][j], a[j][i] :", i, j, a[i][j], a[j][i])
				q = false
			}
		}
	}
	return q
}

// IsQ tests whether the matrix is a Q-matrix (columnwise). 
// See Kendall, 1971, for definition.
func (a Matrix64) IsQ() bool {
	nRow, nCol := a.Dims()
	q := true
	for j := 0; j < nCol && q; j++ {
		//find peak in column j
		peak := -inf
		peakPos := 0
		for i := 0; i < nRow; i++ {
			if a[i][j] >= peak {
				peak = a[i][j]
				peakPos = i

			} else {
				break
			}
		}

		// now test whether further is nonincreasing
		for i := peakPos + 1; i < nRow && q; i++ {
			if a[i][j] > a[i-1][j] {
				q = false
			}
		}
	}
	return q
}

// IsR tests whether the square matrix is a R-matrix (Robinson, columnwise). 
// See Kendall, 1971, for definition.
func (a Matrix64) IsR() bool {
	q := a.IsSymmetric()
	if q {
		q = a.IsQ()
	}
	return q
}

// IsAR tests whether the square matrix is an Anti-R-matrix (Anti-Robinson, columnwise). 
// See Kendall, 1971, for definition.
func (a Matrix64) IsAR() bool {
	if !a.IsSymmetric() {
		return false
	}
	nRow, nCol := a.Dims()
	q := true
	for j := 0; j < nCol && q; j++ {
		//find minimum in column j
		minimum := inf
		minimumPos := 0
		for i := 0; i < nRow; i++ {
			if a[i][j] <= minimum {
				minimum = a[i][j]
				minimumPos = i

			} else {
				break
			}
		}

		// now test whether further is nonincreasing
		for i := minimumPos + 1; i < nRow && q; i++ {
			if a[i][j] < a[i-1][j] {
				q = false
			}
		}
	}
	return q
}

// IntRound clones the matrix into an integer matrix. 
func (a Matrix64) IntRound() IntMatrix {
	out := NewIntMatrix(a.Rows(), a.Cols())
	for i, row := range a {
		for j, _ := range row {
			out[i][j] = iRound(a[i][j])
		}
	}
	return out
}

// SimToDist converts similarity matrix to distance matrix, and vice versa (ad hoc !!!)
func (a Matrix64) SimToDist() {
	// find max value
	maxVal := -inf
	for _, row := range a {
		for _, val := range row {
			if val > maxVal {
				maxVal = val
			}
		}
	}
	// calc distance
	for i, row := range a {
		for j, val := range row {
			a[i][j] = maxVal - val
		}
	}
}

// DistToSim converts distance matrix to similarity matrix
func (a Matrix64) DistToSim(lambda float64) {
	//To convert distance to similarity value, we
	//adopt the formula inspired by Mochihashi, and
	//Matsumoto 2002.
	//similarity( x , y ) = exp{− λ ⋅ distance( x , y )}
	for i, row := range a {
		for j, val := range row {
			a[i][j] = math.Exp(-lambda * val)
		}
	}
}

func (a Matrix64) PermuteRows(p IntVector) {
	if p.Len() != a.Rows() || !p.IsPermutation() {
		panic("bad permutation vector")
	}
	b := a.Clone()
	for i, row := range a {
		for j, _ := range row {
			b[i][j] = a[p[i]][j]
		}
	}
	a.CopyFrom(b)
}

func (a Matrix64) PermuteCols(p IntVector) {
	if p.Len() != a.Cols() || !p.IsPermutation() {
		panic("bad permutation vector")
	}
	b := a.Clone()
	for i, row := range a {
		for j, _ := range row {
			b[i][j] = a[i][p[j]]
		}
	}
	a.CopyFrom(b)
}

func (a Matrix64) Permute(pRow, pCol IntVector) {
	if pRow.Len() != a.Rows() || !pRow.IsPermutation() {
		panic("bad row permutation vector")
	}
	if pCol.Len() != a.Cols() || !pCol.IsPermutation() {
		panic("bad col permutation vector")
	}
	b := a.Clone()
	for i, row := range a {
		for j, _ := range row {
			b[i][j] = a[pRow[i]][pCol[j]]
		}
	}
	a.CopyFrom(b)
}

func (a Matrix64) Equals(b Matrix64) bool {
	r1, c1 := a.Dims()
	r2, c2 := b.Dims()
	if r1 != r2 || c1 != c2 {
		return false
	}
	for i, row := range a {
		for j, val := range row {
			if b[i][j] != val {
				return false
			}
		}
	}
	return true
}

func (a Matrix64) ForceTo01() {
	// find maximum and minimum
	max := -inf
	min := inf
	for _, row := range a {
		for _, val := range row {
			if val < min {
				min = val
			}
			if val > max {
				max = val
			}
		}
	}

	// recalc
	span := max - min
	for i, row := range a {
		for j, val := range row {
			a[i][j] = (val - min) / span
		}
	}
	return
}
