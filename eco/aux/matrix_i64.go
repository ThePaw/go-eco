package aux

import (
	"bufio"
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

// Print prints the matrix on stdout. 
func (m *MatrixInt64) PrintCSV() {
	var i, j int
	for i = 0; i < m.R; i++ {
		fmt.Printf("%d", m.Get(i, 0))	// no leading comma
		for j = 1; j < m.C; j++ {
			fmt.Printf(", %d", m.Get(i, j))
		}
		fmt.Print("\n")
	}
}

func skip(rd *bufio.Reader) {
	var b byte = ' '
	var err error
	for b == ' ' || b == '\t' || b == '\n' {
		b, err = rd.ReadByte()
		if err != nil {
			return
		}
	}
	rd.UnreadByte()
}

func wskip(s string) string {
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' && s[i] != '\t' {
			return s[i:]
		}
	}
	return ""
}

func end(s string) (i int64) {
	for i = 0; i < int64(len(s)); i++ {
		if s[i] == ' ' || s[i] == '\t' || s[i] == '\n' {
			return i
		}
	}
	return 0
}

func readUint(s string) (int64, int64) {
	i := end(s)
	x, _ := strconv.ParseInt(s[:i], 10, 64)
	return int64(x), i
}

// ReadMatrixInt64 reads the matrix. 
func ReadMatrixInt64(rd *bufio.Reader, n, m int) *MatrixInt64 {
	M := NewMatrixInt64(n, m)
	for i := 0; i < n; i++ {
		skip(rd)
		line, _ := rd.ReadString('\n')
		for j := 0; j < m; j++ {
			line = wskip(line)
			x, p := readUint(line)
			M.Set(j, i, x)
			if p == 0 {
				panic("bad int")
			}
			line = line[p:]
		}
	}
	return M
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
