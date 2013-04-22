package ser

import (
	"fmt"
	"math/rand"
)

type IntVector []int

// NewIntVector creates a new IntVector instance with specified number of elements. 
func NewIntVector(nElem int) IntVector {
	v := make([]int, nElem)
	return v
}

// Len returns number of elements in the vector. 
func (v IntVector) Len() int {
	return len(v)
}

// Copy from an existing vector
func (v IntVector) CopyFrom(w IntVector) {
	n := v.Len()
	if n > 0 {
		copy(v, w)
	}
}

// Clone to a new vector
func (v IntVector) Clone() IntVector {
	n := v.Len()
	w := NewIntVector(n)
	if n > 0 {
		copy(w, v)
	}
	return w
}

// Swap elements i, j
func (v IntVector) Swap(i, j int) {
	x := v[i]
	v[i] = v[j]
	v[j] = x
}

// Create ordered sequence 0 .. n. 
func (v IntVector) Order() {
	n := v.Len()
	for i := 0; i < n; i++ {
		v[i] = i
	}
}

// Fill in a pseudo-random permutation of the integers [0,n).
func (v IntVector) Perm() {
	n := v.Len()
	copy(v, rand.Perm(n))
	return
}

// Permute the vector randomly. 
func (v IntVector) Permute() {
	n := v.Len()
	for j := 0; j < 20; j++ { // the '20' constant is ugly, but it works for me now
		for i := 0; i < n; i++ {
			v.Swap(i, i+rand.Intn(n-i))
		}
	}
}

func (v IntVector) ReadCSV() {
	// to be implemented
}

func (v IntVector) WriteCSV() {
	// to be implemented
}

func (v IntVector) Print() {
	for i := 0; i < len(v); i++ {
		fmt.Printf("%d ", v[i])
	}
	fmt.Print("\n")
}

func (v IntVector) Increasing() {
	smp := len(v)
	half := smp / 2

	// sum labels up to half-length
	sum1 := 0
	for i := 0; i < half; i++ {
		sum1 += v[i]
	}

	sum2 := 0
	for i := half; i < smp; i++ {
		sum2 += v[i]
	}

	if sum2 < sum1 {
		// reverse the series
		for i := 0; i < half; i++ {
			tmp := v[i]
			v[i] = v[smp-i-1]
			v[smp-i-1] = tmp
		}
	}
	return
}

func (v IntVector) IsIdentical(w IntVector) bool {
	nElem := v.Len()
	if w.Len() != nElem {
		return false
	}

	for i, val := range v {
		if w[i] != val {
			return false
		}
	}
	return true
}
