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

// Print prints the matrix on stdout. 
func (m *Matrix) Print() {
	var i, j int
	for i = 0; i < m.R; i++ {
		for j = 0; j < m.C; j++ {
			fmt.Printf("%f ", m.Get(i, j))
		}
		fmt.Print("\n")
	}
}

// ReadCsvMatrix opens a CSV file and reads the matrix from it. 
 func ReadCsvMatrix(s string) (m *Matrix) {
	f, err := os.Open(s)
	if err != nil {
		fmt.Println("Could not open the CSV File")
		return
	}
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