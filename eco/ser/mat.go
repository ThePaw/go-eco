package ser

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Matrix64 is 2D array of integers. Elements are stored in a single int slice and slices of each row are created.

type Matrix64 [][]float64

// NewMatrix64 creates a new IntMatrix instance with specified number of rows and columns.
func NewMatrix64(nRow, nCol int) Matrix64 {
	s := make([]float64, nCol*nRow)
	m := make(Matrix64, nRow)
	for i, p := 0, 0; i < nRow; i++ {
		m[i] = s[p : p+nCol]
		p += nCol
	}
	return m
}

// Dims return the dimensions of the matrix.
func (m Matrix64) Dims() (nRow, nCol int) {
	return len(m), len(m[0])
}

// Rows returns the number of rows. 
func (m Matrix64) Rows() int {
	return len(m)
}

// Cols returns the number of columns. 
func (m Matrix64) Cols() int {
	if len(m) == 0 {
		return 0
	}
	return len(m[0])
}

// CopyTo copies matrix to an existing matrix. 
func (m Matrix64) CopyTo(targetMat Matrix64) {
	n := m.Rows() * m.Cols()
	if n > 0 {
		copy(targetMat[0][:n], m[0][:n])
	}
	return
}

// CopyFrom copies matrix from an existing matrix.
func (m Matrix64) CopyFrom(srcMat Matrix64) {
	n := m.Rows() * m.Cols()
	if n > 0 {
		copy(m[0][:n], srcMat[0][:n])
	}
	return
}

// Clone clones a matrix.
func (m Matrix64) Clone() Matrix64 {
	clone := NewMatrix64(m.Rows(), m.Cols())
	n := m.Rows() * m.Cols()
	if n > 0 {
		copy(clone[0][:n], m[0][:n])
	}
	return clone
}

// Swap rows i, j
func (m Matrix64) SwapRows(i int, j int) {
	nCol := m.Cols()
	for k := 0; k < nCol; k++ {
		x := m[i][k]
		m[i][k] = m[j][k]
		m[j][k] = x
	}
}

// Swap columns i, j
func (m Matrix64) SwapCols(i int, j int) {
	nRow := m.Rows()
	for k := 0; k < nRow; k++ {
		x := m[k][i]
		m[k][i] = m[k][j]
		m[k][j] = x
	}
}

// ReadCsvMatrix64  reads the matrix from an opened CSV file. 
func ReadCsvMatrix64(f *os.File) (m Matrix64) {
	read := csv.NewReader(io.Reader(f))
	data, err := read.ReadAll()
	if err != nil {
		panic("Failed to read from the CSV File(Maybe the file does not comply to the CSV standard defined in RFC 4180)")
	}
	nRow := len(data)
	nCol := len(data[0])
	m = NewMatrix64(nRow, nCol)
	for i, row := range m {
		for j, _ := range row {
			x, _ := strconv.ParseFloat(data[i][j], 10)
			m[i][j] = float64(x)
		}
	}
	return
}

/*
func (m IntMatrix) WriteCSV(f *os.File)  {	// to be implemented
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

func (m Matrix64) WriteCSV() {
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

func (m Matrix64) Print() {
	for i, row := range m {
		for j, _ := range row {
			fmt.Printf("%d ", m[i][j])
		}
		fmt.Println()
	}
}

// PrettyString returns a pretty string form of the matrix
func (m Matrix64) PrettyString() string {
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
	nRow, nCol := a.Dims()
	c := NewMatrix64(nRow, nCol)
	for i, row := range c {
		for j, _ := range row {
			c[i][j] = 0
			for k, _ := range row {
				c[i][j] += min(a[i][k], b[k][j])
			}
		}
	}
	return c
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
		peak := -Inf
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

// IntRound clones the matrix into an integer matrix. 
func (m Matrix64) IntRound() IntMatrix {
	out := NewIntMatrix(m.Rows(), m.Cols())
	for i, row := range m {
		for j, _ := range row {
			out[i][j] = iRound(m[i][j])
		}
	}
	return out
}
