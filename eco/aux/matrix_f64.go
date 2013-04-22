package aux

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Matrix represents a dense matrix struct. 
type Matrix struct {
	R int
	C int
	A []float64
}

// NewMatrix returns a pointer to new instance of Matrix. 
func NewMatrix(rows, cols int) (m *Matrix) {
	m = new(Matrix)
	m.R = rows
	m.C = cols
	m.A = make([]float64, rows*cols)
	return m
}

// Set sets x_ij element of the matrix. 
func (m Matrix) Set(i int, j int, x float64) {
	//	m.A[i+j*m.C] = x
	m.A[i*m.C+j] = x
}

// Get returns x_ij th element of the matrix. 
func (m Matrix) Get(i int, j int) float64 {
	return m.A[i*m.C+j]
}

// PrintRC prints number of rows  and columns of  the matrix on stdout. 
func (m *Matrix) PrintRC() {
	fmt.Printf("%d %d \n \n", m.R, m.C)
}

// Print prints the matrix on stdout. 
func (m *Matrix) Print() {
	var i, j int
	for i = 0; i < m.R; i++ {
		for j = 0; j < m.C; j++ {
			fmt.Printf("%f ", m.Get(i, j))
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

// Transpose returns a pointer to transposed matrix. 
func (m *Matrix) Transpose() *Matrix {
	b := NewMatrix(m.C, m.R)
	for i := 0; i < m.R; i++ {
		for j := 0; j < m.C; j++ {
			b.Set(j, i, m.Get(i, j))
		}
	}
	return b
}

// Scale scales the matrix (in place) to the specified maximum value. 
func (m *Matrix) Scale(scale float64) {
	var maxVal float64 = 0
	rows := m.R
	cols := m.C
	// find max value
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if m.Get(i, j) > maxVal {
				maxVal = m.Get(i, j)
			}
		}
	}
	coeff := scale / maxVal
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			m.Set(i, j, coeff*m.Get(i, j))
		}
	}
}

// ReadCsvMatrix  reads the matrix from an opened CSV file. 
func ReadCsvMatrix(f *os.File) (m *Matrix) {
	read := csv.NewReader(io.Reader(f))
	data, err := read.ReadAll()
	if err != nil {
		fmt.Println("Failed to read from the CSV File(Maybe the file does not comply to the CSV standard defined in RFC 4180)")
	}
	rows := len(data)
	cols := len(data[0])
	m = NewMatrix(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			x, _ := strconv.ParseFloat(data[i][j], 64)
			m.Set(i, j, x)
		}
	}
	return
}

/*
func WriteCsvMatrix(s string) () {
	f, err := os.Open(s)
	if err != nil {
		fmt.Println("Could not open the CSV File")
		return
	}
	write := csv.NewWriter(io.Writer(f))
	data, err := write.WriteAll()
	return
}

func (r *Reader) ReadAll() (records [][]string, err error)
func (w *Writer) WriteAll(records [][]string) (err error)
*/
