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
	m := make(IntMatrix, nRow)
	for i, p := 0, 0; i < nRow; i++ {
		m[i] = s[p : p+nCol]
		p += nCol
	}
	return m
}

// Dims return the dimensions of the matrix.
func (m IntMatrix) Dims() (int, int) {
	return len(m), len(m[0])
}

// Rows returns the number of rows
func (m IntMatrix) Rows() int {
	return len(m)
}

// Cols returns the number of columns
func (m IntMatrix) Cols() int {
	if len(m) == 0 {
		return 0
	}
	return len(m[0])
}

// Copy to an existing matrix
func (m IntMatrix) CopyTo(targetMat IntMatrix) {
	n := m.Rows() * m.Cols()
	if n > 0 {
		//		copy(targetMat[0][:n], m[0][:n])
		copy(targetMat, m)
	}
	return
}

// Copy from an existing matrix
func (m IntMatrix) CopyFrom(srcMat IntMatrix) {
	n := m.Rows() * m.Cols()
	if n > 0 {
		copy(m, srcMat)
	}
	return
}

// Clone clones an IntMatrix
func (m IntMatrix) Clone() IntMatrix {
	clone := NewIntMatrix(m.Rows(), m.Cols())
	n := m.Rows() * m.Cols()
	if n > 0 {
		//		copy(clone[0][:n], m[0][:n])
		copy(clone, m)
	}
	return clone
}

// Swap rows i, j
func (m IntMatrix) SwapRows(i int, j int) {
	nCol := m.Cols()
	for k := 0; k < nCol; k++ {
		x := m[i][k]
		m[i][k] = m[j][k]
		m[j][k] = x
	}
}

// Swap columns i, j
func (m IntMatrix) SwapCols(i int, j int) {
	nRow := m.Rows()
	for k := 0; k < nRow; k++ {
		x := m[k][i]
		m[k][i] = m[k][j]
		m[k][j] = x
	}
}

// ReadCsvIntMatrix  reads the matrix from an opened CSV file. 
func ReadCsvIntMatrix(f *os.File) (m IntMatrix) {
	read := csv.NewReader(io.Reader(f))
	data, err := read.ReadAll()
	if err != nil {
		panic("Failed to read from the CSV File(Maybe the file does not comply to the CSV standard defined in RFC 4180)")
	}
	nRow := len(data)
	nCol := len(data[0])
	m = NewIntMatrix(nRow, nCol)
	for i, row := range m {
		for j, _ := range row {
			x, _ := strconv.ParseInt(data[i][j], 10, 32) // why this is not int??
			m[i][j] = int(x)
		}
	}
	return
}

/*
func (m IntMatrix) WriteCSV(f *os.File)  {
	// to be implemented
	write := csv.NewWriter(io.Writer(f))
	records := // [][]string TO BE IMPLEMENTED
	nRow, cols := m.Dims()
	for i := 0; i < nRow; i++ {
		for j := 0; j < cols; j++ {
			records[i][j] = strconv.FormatInt(m[i][j], 10)
		}
	}
	err :=WriteAll(records)
	if err != nil {
		fmt.Println("Failed to write the CSV File")
	}
}
*/

func (m IntMatrix) WriteCSV() {
	for i, row := range m {
		for j, _ := range row {
			if j == 0 {
				fmt.Print(m[i][j])
			} else {
				fmt.Print(",", m[i][j])
			}
		}
		fmt.Println()
	}
}

func (m IntMatrix) WriteGo() {
	fmt.Println("matrix := IntMatrix{")
	for i, row := range m {
		fmt.Print("{")
		for j, _ := range row {
			if j == 0 {
				fmt.Print(m[i][j])
			} else {
				fmt.Print(",", m[i][j])
			}
		}
		fmt.Println("},")
	}
	fmt.Println("}")
}

func (m IntMatrix) Print() {
	for i, row := range m {
		for j, _ := range row {
			fmt.Printf("%d ", m[i][j])
		}
		fmt.Println()
	}
}

// PrettyString returns a pretty string form of the matrix
func (m IntMatrix) PrettyString() string {
	sa := make([][]string, 0, m.Rows())
	for _, row := range m {
		sr := make([]string, 0, len(row))
		for _, cell := range row {
			sr = append(sr, fmt.Sprint(cell))
		} // for cell
		sa = append(sa, sr)
	} // for row

	wds := make([]int, m.Cols())
	for j := 0; j < m.Cols(); j++ {
		for i := 0; i < m.Rows(); i++ {
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
			res += fmt.Sprintf("](%dx%d)", m.Rows(), m.Cols())
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

// RearrangeRows returns rearranged matrix according to a row permutation vector. 
func (a IntMatrix) RearrangeRows(r IntVector) {
	// r - rows permutation vector
	nRow, nCol := a.Dims()
	if r.Len() != nRow {
		panic("bad dimensions")
	}
	newMat := NewIntMatrix(nRow, nCol)
	for i, row := range a {
		for j, _ := range row {
			newMat[i][j] = a[r[i]][j]
		}
	}
	// Copy back
	a.CopyFrom(newMat)
	return
}

// RearrangeCols returns rearranged matrix according to a row permutation vector. 
func (a IntMatrix) RearrangeCols(c IntVector) {
	// c - columns permutation vector
	nRow, nCol := a.Dims()
	if c.Len() != nCol {
		panic("bad dimensions")
	}
	newMat := NewIntMatrix(nRow, nCol)
	for i, row := range a {
		for j, _ := range row {
			newMat[i][j] = a[i][c[j]]
		}
	}
	// Copy back
	a.CopyFrom(newMat)
	return
}

// Rearrange returns rearranged matrix according to row and column permutation vectors. 
func (a IntMatrix) Rearrange(r, c IntVector) {
	// r - rows permutation vector
	// c - columns permutation vector

	nRow, nCol := a.Dims()
	if r.Len() != nRow || c.Len() != nCol {
		panic("bad dimensions")
	}
	newMat := NewIntMatrix(nRow, nCol)
	for i, row := range a {
		for j, _ := range row {
			newMat[i][j] = a[r[i]][c[j]]
		}
	}
	// Copy back
	//	a.CopyFrom(newMat)
	a.CopyFrom(newMat)
	return
}

// IsEqual tests whether the matrices are the same. 
func (a IntMatrix) IsEqual(b IntMatrix) bool {
	nRowA, nColA := a.Dims()
	nRowB, nColB := b.Dims()
	if nRowA != nRowB || nColA != nColB {
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
