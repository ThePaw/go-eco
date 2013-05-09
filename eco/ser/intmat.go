package ser

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

// IntMatrix is 2D array of integers. Elements are stored in a single int slice and slices of each row are created.

type IntMatrix [][]int

// NewIntMatrix creates a new IntMatrix instance with specified number of rows and columns
func NewIntMatrix(nRow, nCol int) IntMatrix {
	s := make([]int, nCol*nRow)
	a := make(IntMatrix, nRow)
	for i, p := 0, 0; i < nRow; i++ {
		a[i] = s[p : p+nCol]
		p += nCol
	}
	return a
}

// Dims return the dimensions of the matrix.
func (a IntMatrix) Dims() (int, int) {
	return len(a), len(a[0])
}

// Rows returns the number of rows
func (a IntMatrix) Rows() int {
	return len(a)
}

// Cols returns the number of columns
func (a IntMatrix) Cols() int {
	if len(a) == 0 {
		return 0
	}
	return len(a[0])
}

// Copy to an existing matrix
func (a IntMatrix) CopyTo(targetMat IntMatrix) {
	for i, row := range a {
		for j, _ := range row {
			targetMat[i][j] = a[i][j]
		}
	}
	return
}

// Copy from an existing matrix
func (a IntMatrix) CopyFrom(srcMat IntMatrix) {
	for i, row := range a {
		for j, _ := range row {
			a[i][j] = srcMat[i][j]
		}
	}
	return
}

// Clone clones an IntMatrix
func (a IntMatrix) Clone() IntMatrix {
	clone := NewIntMatrix(a.Dims())
	for i, row := range a {
		for j, val := range row {
			clone[i][j] = val
		}
	}
	return clone
}

// Swap rows i, j
func (a IntMatrix) SwapRows(i int, j int) {
	nCol := a.Cols()
	for k := 0; k < nCol; k++ {
		x := a[i][k]
		a[i][k] = a[j][k]
		a[j][k] = x
	}
}

// Swap columns i, j
func (a IntMatrix) SwapCols(i int, j int) {
	nRow := a.Rows()
	for k := 0; k < nRow; k++ {
		x := a[k][i]
		a[k][i] = a[k][j]
		a[k][j] = x
	}
}

// ReadCsvIntMatrix  reads the matrix from an opened CSV file. 
func ReadCsvIntMatrix(f *os.File) (a IntMatrix) {
	read := csv.NewReader(io.Reader(f))
	data, err := read.ReadAll()
	if err != nil {
		panic("Failed to read from the CSV File(Maybe the file does not comply to the CSV standard defined in RFC 4180)")
	}
	nRow := len(data)
	nCol := len(data[0])
	a = NewIntMatrix(nRow, nCol)
	for i, row := range a {
		for j, _ := range row {
			x, _ := strconv.ParseInt(data[i][j], 10, 32) // why this is not int??
			a[i][j] = int(x)
		}
	}
	return
}

/*
func (a IntMatrix) WriteCSV(f *os.File)  {
	// to be implemented
	write := csv.NewWriter(io.Writer(f))
	records := // [][]string TO BE IMPLEMENTED
	nRow, cols := m.Dims()
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

func (a IntMatrix) WriteCSV() {
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
}

func (a IntMatrix) WriteGo() {
	fmt.Println("matrix := IntMatrix{")
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

func (a IntMatrix) Print() {
	for i, row := range a {
		for j, _ := range row {
			fmt.Printf("%d ", a[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

// PrettyString returns a pretty string form of the matrix
func (a IntMatrix) PrettyString() string {
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
func (a IntMatrix) Product(b IntMatrix) IntMatrix {
	nRow, nCol := a.Dims()
	bRow, bCol := b.Dims()
	if nRow != bCol || nCol != bRow {
		panic("bad dimensions")
	}
	c := NewIntMatrix(nRow, nRow)
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
func (a IntMatrix) CircleProduct(b IntMatrix) IntMatrix {
	r1, c1 := a.Dims()
	r2, c2 := b.Dims()
	if r1 != c2 || c1 != r2 {
		panic("bad dimensions")
	}
	c := NewIntMatrix(r1, r1)
	for i, row := range a {
		for j, _ := range a {
			c[i][j] = 0
			for k, _ := range row { // columns of a
				c[i][j] += imin(a[i][k], b[k][j])
			}
		}
	}
	return c
}

// CirclePower computes circular power of a matrix. 
// See Kendall, 1971: 111, for definition.
func (a IntMatrix) CirclePower(n int) IntMatrix {
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
func (a IntMatrix) Transpose() IntMatrix {
	nRow, nCol := a.Dims()
	c := NewIntMatrix(nCol, nRow)
	for i, row := range c {
		for j, _ := range row {
			c[i][j] = a[j][i]
		}
	}
	return c
}

// IsSquare tests whether the matrix is square matrix. 
func (a IntMatrix) IsSquare() bool {
	q := true
	nRow, nCol := a.Dims()
	if nRow != nCol {
		q = false
	}
	return q
}

// IsSymmetric tests whether the square matrix is symmetric. 
func (a IntMatrix) IsSymmetric() bool {

	if !a.IsSquare() {
		panic("not a square matrix")
	}

	nRow, nCol := a.Dims()
	q := true
	for i := 0; i < nRow && q; i++ {
		for j := 0; j < nCol && q; j++ {
			if a[i][j] != a[j][i] {
				q = false
			}
		}
	}
	return q
}

// IsQ tests whether the matrix is a Q-matrix (columnwise). 
// See Kendall, 1971, for definition.
func (a IntMatrix) IsQ() bool {
	nRow, nCol := a.Dims()
	q := true
	for j := 0; j < nCol && q; j++ {
		//find peak in column j
		peak := -999
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
func (a IntMatrix) IsR() bool {
	q := a.IsSymmetric()
	if q {
		q = a.IsQ()
	}
	return q
}

// IsAR tests whether the square matrix is an Anti-R-matrix (Anti-Robinson, columnwise). 
// See Kendall, 1971, for definition.
func (a IntMatrix) IsAR() bool {
	if !a.IsSymmetric() {
		return false
	}
	nRow, nCol := a.Dims()
	q := true
	for j := 0; j < nCol && q; j++ {
		//find minimum in column j
		minimum := iInf
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

// PermuteRows returns rearranged matrix according to a row permutation vector. 
func (a IntMatrix) PermuteRows(p IntVector) {
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

// PermuteCols returns rearranged matrix according to a row permutation vector. 
func (a IntMatrix) PermuteCols(p IntVector) {
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

// Permute returns rearranged matrix according to row and column permutation vectors. 
func (a IntMatrix) Permute(pRow, pCol IntVector) {
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

// IsIdentical tests whether the matrices are the same. 
func (a IntMatrix) IsIdentical(b IntMatrix) bool {
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
