package aux

import (
	"encoding/csv"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
)

// Vector represents a dense vector struct. 
type Vector struct {
	A []float64 // data
	L int       // length
}

// NewVector returns a pointer to new instance of Vector. 
func NewVector(length int) (v *Vector) {
	v = new(Vector)
	v.L = length
	v.A = make([]float64, length)
	return v
}

// Set sets x_ij element of the vector. 
func (v Vector) Set(i int, x float64) {
	v.A[i] = x
}

// Get returns x_i th element of the vector. 
func (v Vector) Get(i int) float64 {
	return v.A[i]
}

// Swap swaps ith and jth element of a vector. 
func (p Vector) Swap(i int, j int) {
	x := p.A[i]
	p.A[i] = p.A[j]
	p.A[j] = x
}

// Len returns number of elements  in a vector (= length). 
func (v Vector) Len() int {
	return v.L
}

// Copy copies a vector to another one. 
func (v Vector) Copy(w Vector) {
	for i := 0; i < v.L; i++ {
		v.A[i] = w.A[i]
	}
}

// Print prints the vector on stdout. 
func (v Vector) Print() {
	for i := 0; i < v.L; i++ {
		fmt.Printf("%f ", v.A[i])
	}
	fmt.Print("\n")
}

// Perm permutes a vector.  
func Perm(p Vector) {
	n := p.L
	var i int
	for i = 0; i < n; i++ {
		p.A[i] = float64(i)
	}
	for i = 0; i < n; i++ {
		p.Swap(i, i+rand.Intn(n-i))
	}
}

// FetchCsvVector opens a CSV file and reads the vector from it. 
func FetchCsvVector(s string) (v *Vector) {
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
	v = NewVector(cols)
	for i := 0; i < cols; i++ {
		x, _ := strconv.ParseFloat(data[0][i], 64)
		v.Set(i, x)
	}
	return
}
