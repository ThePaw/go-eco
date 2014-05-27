package ser

import (
	"fmt"
	"math/rand"
)

type Vector64 []float64

// NewVector64 creates a new Vector64 instance with specified number of elements
func NewVector64(nElem int) Vector64 {
	v := make([]float64, nElem)
	return v
}

// Len returns number of elements in the vector
func (v Vector64) Len() int {
	return len(v)
}

// Copy to an existing vector
func (v Vector64) Copy(w Vector64) {
	n := v.Len()
	if n > 0 {
		copy(w, v)
	}
}

// Clone to a new vector
func (v Vector64) Clone() Vector64 {
	n := v.Len()
	w := NewVector64(n)
	if n > 0 {
		copy(w, v)
	}
	return w
}

// Swap elements i, j
func (v Vector64) Swap(i, j int) {
	x := v[i]
	v[i] = v[j]
	v[j] = x
}

// Permute the vector randomly.
func (v Vector64) Permute() {
	n := v.Len()
	w := v.Clone()
	p := rand.Perm(n)
	for i := 0; i < n; i++ {
		v[i] = w[p[i]]
	}
}

func (v Vector64) ReadCSV() {
	// to be implemented
}

func (v Vector64) WriteCSV() {
	// to be implemented
}

func (v Vector64) Print() {
	for i := 0; i < len(v); i++ {
		fmt.Printf("%f ", v[i])
	}
	fmt.Print("\n")
}
