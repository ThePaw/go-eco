package ser

import (
	"fmt"
	"math/rand"
	"sort"
)

type IntVector64 []int

// NewIntVector64 creates a new IntVector64 instance with specified number of elements. 
func NewIntVector64(nElem int) IntVector64 {
	v := make([]int, nElem)
	return v
}

// Len returns number of elements in the vector. 
func (v IntVector64) Len() int {
	return len(v)
}

// Copy from an existing vector
func (v IntVector64) CopyFrom(w IntVector64) {
	n := v.Len()
	if n > 0 {
		copy(v, w)
	}
}

// Clone to a new vector
func (v IntVector64) Clone() IntVector64 {
	n := v.Len()
	w := NewIntVector64(n)
	if n > 0 {
		copy(w, v)
	}
	return w
}

// Swap elements i, j
func (v IntVector64) Swap(i, j int) {
	x := v[i]
	v[i] = v[j]
	v[j] = x
}

// Fill-in ordered sequence 0 .. n-1. 
func (v IntVector64) Order() {
	n := v.Len()
	for i := 0; i < n; i++ {
		v[i] = i
	}
}

// Fill in a pseudo-random permutation of the integers [0,n).
func (v IntVector64) Perm() {
	n := v.Len()
	copy(v, rand.Perm(n))
	return
}

func (v IntVector64) ReadCSV() {
	// to be implemented
}

func (v IntVector64) WriteCSV() {
	for i, val := range v {
		if i == 0 {
			fmt.Print(val)
		} else {
			fmt.Print(",", val)
		}
	}
	fmt.Println()
}

func (v IntVector64) WriteGo() {
	fmt.Println("vector := IntVector64{")
	for i, val := range v {
		if i == 0 {
			fmt.Print(val)
		} else {
			fmt.Print(",", val)
		}
	}
	fmt.Println("}")
	fmt.Println()
}

func (v IntVector64) Print() {
	for i := 0; i < len(v); i++ {
		fmt.Printf("%d ", v[i])
	}
	fmt.Print("\n")
}

// Increasing reverses the order in case that lower half sum > upper half sum
func (v IntVector64) Increasing() {
	smp := len(v)
	half := smp / 2
	sum1 := 0
	if isOdd(smp) == false {
		// sum elements up to half-length
		for i := 0; i < half; i++ {
			sum1 += v[i]
		}
	} else {
		// sum elements up to half-length, including the mid-element
		for i := 0; i <= half; i++ {
			sum1 += v[i]
		}

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

func (v IntVector64) Equals(w IntVector64) bool {
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

func (v IntVector64) IsPermutation() bool {
	w := v.Clone()
	w.Order()
	z := v.Clone()
	sort.Ints(z)
	if !w.Equals(z) {
		return false
	}
	return true
}
