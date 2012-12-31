package aux

import (
	"encoding/csv"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
)

// VectorInt64 represents a dense vector struct. 
type VectorInt64 struct {
	A []int64 // data
	L int       // length
}

// NewVectorInt64 returns a pointer to new instance of VectorInt64. 
func NewVectorInt64(length int) (v *VectorInt64) {
	v = new(VectorInt64)
	v.L = length
	v.A = make([]int64, length)
	return v
}

// Set sets x_ij element of the vector. 
func (v VectorInt64) Set(i int, x int64) {
	v.A[i] = x
}

// Get returns x_i th element of the vector. 
func (v VectorInt64) Get(i int) int64 {
	return v.A[i]
}

// Swap swaps ith and jth element of a vector. 
func (p VectorInt64) Swap(i int, j int) {
	x := p.A[i]
	p.A[i] = p.A[j]
	p.A[j] = x
}

// Len returns number of elements  in a vector (= length). 
func (v VectorInt64) Len() int {
	return v.L
}

// Copy copies a vector to another one. 
func (v VectorInt64) Copy(w VectorInt64) {
	for i := 0; i < v.L; i++ {
		v.A[i] = w.A[i]
	}
}

// Print prints the vector on stdout. 
func (v VectorInt64) Print() {
	for i := 0; i < v.L; i++ {
		fmt.Printf("%f ", v.A[i])
	}
	fmt.Print("\n")
}

// Perm permutes a vector.  
func PermInt64(p VectorInt64) {
	n := p.L
	var i int
	for i = 0; i < n; i++ {
		p.A[i] = int64(i)
	}
	for i = 0; i < n; i++ {
		p.Swap(i, i+rand.Intn(n-i))
	}
}

// FetchCsvVectorInt64 opens a CSV file and reads the vector from it. 
func FetchCsvVectorInt64(s string) (v *VectorInt64) {
	f, err := os.Open(s)
	if err != nil {
		fmt.Println("Could not Open the CSV File")
		return
	}
	read := csv.NewReader(io.Reader(f))
	data, err := read.ReadAll()
	if err != nil {
		fmt.Println("Failed to read from the CSV File(Maybe the file does not comply to the CSV standard defined in RFC 4180)")
	}
	cols := len(data[0])
	v = NewVectorInt64(cols)
	for i := 0; i < cols; i++ {
		x, _ := strconv.ParseInt(data[0][i], 10, 64)
		v.Set(i, x)
	}
	return
}
