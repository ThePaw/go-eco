package aux

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

// MatrixInt64 represents a dense matrix struct. 
type MatrixInt64 struct {
	R int
	C int
	A []int64
}

// NewMatrixInt64 returns a pointer to new instance of MatrixInt64. 
func NewMatrixInt64(rows, cols int) (m *MatrixInt64) {
	m = new(MatrixInt64)
	m.R = rows
	m.C = cols
	m.A = make([]int64, rows*cols)
	return m
}

// Set sets x_ij element of the matrix. 
func (m MatrixInt64) Set(i int, j int, x int64) {
	//	m.A[i+j*m.C] = x
	m.A[i*m.C+j] = x
}

// Get returns x_ij th element of the matrix. 
func (m MatrixInt64) Get(i int, j int) int64 {
	return m.A[i*m.C+j]
}

// PrintRC prints number of rows  and columns of  the matrix on stdout. 
func (m *MatrixInt64) PrintRC() {
			fmt.Printf("%d %d \n \n", m.R, m.C)
}

// Print prints the matrix on stdout. 
func (m *MatrixInt64) Print() {
	var i, j int
	for i = 0; i < m.R; i++ {
		for j = 0; j < m.C; j++ {
			fmt.Printf("%d ", m.Get(i, j))
		}
		fmt.Print("\n")
	}
		fmt.Print("\n")
}

// ReadCsvMatrixInt64  reads the matrix from an opened CSV file. 
 func ReadCsvMatrixInt64(f *os.File) (m *MatrixInt64) {
	read := csv.NewReader(io.Reader(f))
	data, err := read.ReadAll()
	if err != nil {
		fmt.Println("Failed to read from the CSV File(Maybe the file does not comply to the CSV standard defined in RFC 4180)")
	}
	rows := len(data)
	cols := len(data[0])
	m = NewMatrixInt64(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			x, _ := strconv.ParseInt(data[i][j], 10, 64)
			m.Set(i, j, x)
		}
	}
	return
}
