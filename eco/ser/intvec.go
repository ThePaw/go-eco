package ser

import (
	"fmt"
	"math/rand"
	"sort"
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

// Swap swaps elements i, j
func (v IntVector) Swap(i, j int) {
	// exchange mutation operator (Banzhaf 1990)
	// swap mutation operator (Oliver et al. 1987)
	// point mutation operator (Ambati et al. 1991)
	// reciprocal exchange mutation operator (Michalewicz 1992)
	// order based mutation operator (Syswerda 1991)
	x := v[i]
	v[i] = v[j]
	v[j] = x
}

// Fill-in ordered sequence 0 .. n-1.
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

// Increasing reverses the order in case that lower half sum > upper half sum
func (v IntVector) Increasing() {
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

func (v IntVector) Equals(w IntVector) bool {
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

func (v IntVector) IsPermutation() bool {
	w := v.Clone()
	w.Order()
	z := v.Clone()
	sort.Ints(z)
	if !w.Equals(z) {
		return false
	}
	return true
}

func (v IntVector) Invert() {
	n := v.Len()
	w := v.Clone()

	for i, val := range v {
		w[n-i-1] = val
	}
	v.CopyFrom(w)
}

// InvertFromTo implements simple inversion mutation operator (Holland 1975; Grefenstette 1987).
func (v IntVector) InvertFromTo(a, b int) {
	// simple inversion mutation operator (Holland 1975; Grefenstette 1987)
	n := v.Len()
	if a > n || a < 0 || b > n || b < 0 {
		panic(" bad params")
	}
	if a == b {
		return
	}
	if a > b { //swap
		c := a
		a = b
		b = c
	}
	w := v.Clone()

	for i := 0; i < b-a+1; i++ {
		w[a+i] = v[b-i]
	}
	v.CopyFrom(w)
}

// InvertTo inverts the head of a vector from the beginning to a specified position.
func (v IntVector) InvertTo(a int) {
	n := v.Len()
	if a >= n || a < 0 {
		panic(" bad params")
	}
	w := v.Clone()

	for i := 0; i <= a; i++ {
		w[a-i] = v[i]
	}
	v.CopyFrom(w)
}

// InvertFrom inverts the tail of a vector from specified position to the end.
func (v IntVector) InvertFrom(a int) {
	n := v.Len()
	if a >= n || a < 0 {
		panic(" bad params")
	}
	w := v.Clone()

	for i := 0; i < n-a; i++ {
		w[a+i] = v[n-i-1]
	}
	v.CopyFrom(w)
}

// InsideOut inverts both head and tail of a vector at specified position.
func (v IntVector) InsideOut(a int) {
	n := v.Len()
	if a >= n || a < 0 {
		panic(" bad params")
	}
	w := v.Clone()

	for i := 0; i < a; i++ {
		w[a-i] = v[i]
	}

	for i := 0; i < n-a; i++ {
		w[a+i] = v[n-i-1]
	}
	v.CopyFrom(w)
}

// InsertAt implements the insertion mutation operator (Fogel 1988; Michalewicz 1992).
func (v IntVector) InsertAt(a, b int) {
	// insertion mutation operator (Fogel 1988; Michalewicz 1992)
	// position based mutation operator (Syswerda 1991)
	n := v.Len()
	if a > n || a < 0 || b > n || b < 0 {
		panic(" bad params")
	}
	if a == b {
		return
	}
	if a > b { //swap
		c := a
		a = b
		b = c
	}
	w := v.Clone()

	x := v[a]                // cut
	for i := a; i < b; i++ { //  shift
		w[i] = v[i+1]
	}
	w[b] = x // insert

	v.CopyFrom(w)
}

// InvertHeadAndTail inverts head and tail of a vector, specified by two positions.
func (v IntVector) InvertHeadAndTail(a, b int) {
	n := v.Len()
	if a > n || a < 0 || b > n || b < 0 {
		panic(" bad params")
	}
	if a == b {
		return
	}
	if a > b { //swap
		c := a
		a = b
		b = c
	}
	w := v.Clone()

	// invert head
	for i := 0; i <= a; i++ {
		w[a-i] = v[i]
	}

	// invert tail
	for i := 0; i < n-b; i++ {
		w[b+i] = v[n-i-1]
	}

	v.CopyFrom(w)
}

// Displace implements a mutation modified from displacement mutation operator (Michalewicz 1992) ==  cut mutation (Banzhaf 1990).
func (v IntVector) Displace(a, b, c int) {
	n := v.Len()
	if a > n || a < 0 || b > n || b < 0 || c > n || c < 0 {
		panic(" bad params")
	}
	if a == b || b == c || c == a {
		return
	}
	// sort ascending
	if a > b { //swap
		x := a
		a = b
		b = x
	}
	if a > c { //swap
		x := c
		c = b
		b = x
	}
	if a > b { //swap
		x := a
		a = b
		b = x
	}
	if b > c { //swap
		x := b
		b = c
		c = x
	}
	w := v.Clone()
	left := NewIntVector(b - a)
	m := left.Len()
	for i := 0; i < m; i++ { //  copy left block
		left[i] = v[i+a]
	}
	for i := 0; i < c-b; i++ { //  shift right block
		w[i+a] = v[i+b]
	}
	for i := 0; i < m; i++ { //  insert left block
		w[i+a+c-b] = left[i]
	}
	v.CopyFrom(w)
}

// DisplaceInv works like Displace, but inverts the left block.
func (v IntVector) DisplaceInv(a, b, c int) {
	// inversion mutation (Fogel 1990, 1993)
	// cut-inverse mutation operator (Banzhaf 1990)

	n := v.Len()
	if a > n || a < 0 || b > n || b < 0 || c > n || c < 0 {
		panic(" bad params")
	}
	if a == b || b == c || c == a {
		return
	}
	// sort ascending
	if a > b { //swap
		x := a
		a = b
		b = x
	}
	if a > c { //swap
		x := c
		c = b
		b = x
	}
	if a > b { //swap
		x := a
		a = b
		b = x
	}
	if b > c { //swap
		x := b
		b = c
		c = x
	}
	w := v.Clone()
	left := NewIntVector(b - a)
	m := left.Len()
	for i := 0; i < m; i++ { //  copy left block
		left[i] = v[i+a]
	}

	left.Invert() //  invert left block

	for i := 0; i < c-b; i++ { //  shift right block
		w[i+a] = v[i+b]
	}
	for i := 0; i < m; i++ { //  insert left block
		w[i+a+c-b] = left[i]
	}
	v.CopyFrom(w)
}

//TwoPointSwapInv performs the Two point Swapped Inversion,  inspired by Sallabi (2009).
func (v IntVector) TwoPointSwapInv(a, b int) {
	n := v.Len()
	if a > n || a < 0 || b > n || b < 0 {
		panic(" bad params")
	}
	if a == b {
		return
	}
	if a > b { //swap
		c := a
		a = b
		b = c
	}
	head := NewIntVector(a)
	mid := NewIntVector(b - a)
	tail := NewIntVector(n - b)

	// invert head
	for i := 0; i < a; i++ {
		head[i] = v[i]
	}
	head.Invert()

	// copy middle
	for i := 0; i < b-a; i++ {
		mid[i] = v[a+i]
	}

	// invert tail
	for i := 0; i < n-b; i++ {
		tail[i] = v[b+i]
	}
	head.Invert()

	//assemble
	for i := 0; i < n-b; i++ {
		v[i] = tail[i]
	}
	for i := n - b; i < n-a; i++ {
		v[i] = mid[i-n+b]
	}
	for i := n - a; i < n; i++ {
		v[i] = head[i-n+a]
	}
}

// Scramble scrambles a segment of a vector, specified by two positions.
func (v IntVector) Scramble(a, b int) {
	n := v.Len()
	if a > n || a < 0 || b > n || b < 0 {
		panic(" bad params")
	}
	if a == b {
		return
	}
	if a > b { //swap
		c := a
		a = b
		b = c
	}
	mid := NewIntVector(b - a)
	p := NewIntVector(b - a)
	p.Perm()

	// copy and scramble middle segment
	for i := 0; i < b-a; i++ {
		mid[i] = v[a+p[i]]
	}

	// paste back
	for i := 0; i < b-a; i++ {
		v[i+a] = mid[i]
	}
}

// Scramble3 scrambles a segment of length 3 of a vector.
func (v IntVector) Scramble3(a int) {
	n := v.Len()
	if a > n || a < 0 {
		panic(" bad params")
	}
	if a > n-3 {
		a = n - 3
	}

	mid := NewIntVector(3)
	p := NewIntVector(3)
	p.Perm()

	// copy and scramble middle segment
	for i := 0; i < 3; i++ {
		mid[i] = v[a+p[i]]
	}

	// paste back
	for i := 0; i < 3; i++ {
		v[i+a] = mid[i]
	}
}

// Scramble4 scrambles a segment of length 4 of a vector.
func (v IntVector) Scramble4(a int) {
	n := v.Len()
	if a > n || a < 0 {
		panic(" bad params")
	}
	if a > n-4 {
		a = n - 4
	}

	mid := NewIntVector(4)
	p := NewIntVector(4)
	p.Perm()

	// copy and scramble middle segment
	for i := 0; i < 4; i++ {
		mid[i] = v[a+p[i]]
	}

	// paste back
	for i := 0; i < 4; i++ {
		v[i+a] = mid[i]
	}
}

// ThreePointExch is inspired by 3-opt move for the TSP.
func (v IntVector) ThreePointExch(a, b, c int) {
	n := v.Len()
	if a > n || a < 0 || b > n || b < 0 || c > n || c < 0 {
		panic(" bad params")
	}
	w := v.Clone()
	// leave unchanged up to a inclusive
	// go from a to c and inverse-continue to b+1
	j := a + 1
	for i := c; i > b; i-- {
		w[j] = v[i]
		j++
	}
	// go from b+1 to a+1 and continue to b
	for i := a + 1; i <= b; i++ {
		w[j] = v[i]
		j++
	}
	// go from c+1 to end
	for i := c + 1; i < n; i++ {
		w[j] = v[i]
		j++
	}
	v.CopyFrom(w)
}

// FourPointExch is inspired by non-sequential 4-change for the TSP.
func (v IntVector) FourPointExch(a, b, c, d int) {
	n := v.Len()
	if a > n || a < 0 || b > n || b < 0 || c > n || c < 0 || d > n || d < 0 {
		panic(" bad params")
	}
	w := v.Clone()
	// leave unchanged up to a inclusive
	j := a + 1
	// go from a to c+1 and continue to d
	for i := c + 1; i <= d; i++ {
		w[j] = v[i]
		j++
	}
	// go from d to b+1 and continue to c
	for i := b + 1; i <= c; i++ {
		w[j] = v[i]
		j++
	}
	// go from c to a+1 and continue to b
	for i := a + 1; i <= b; i++ {
		w[j] = v[i]
		j++
	}
	// go from b to d+1 and continue to end
	for i := d + 1; i < n; i++ {
		w[j] = v[i]
		j++
	}
	v.CopyFrom(w)
}

// ========================================================================
/*
Lexicographic order and finding the next permutation
Permutation f precedes a permutation g in the lexicographic (alphabetic) order iff for the minimum value of k such that f(k)≠ g(k), we have f(k) < g(k). Starting with the identical permutation f(i) = i for all i, the second algorithm generates sequentially permutaions in the lexicographic order. The algorithm is described in [Dijkstra, p. 71].
E. W. Dijkstra, A Discipline of Programming, Prentice-Hall, 1997.
*/

func (p IntVector) NextPermLex() {
	var i, j int
	n := p.Len()
	i = n - 1
	for p[i-1] >= p[i] {
		i = i - 1
	}
	j = n
	for p[j-1] <= p[i-1] {
		j = j - 1
	}
	p.Swap(i-1, j-1) // swap values at positions (i-1) and (j-1)

	i++
	j = n
	for i < j {
		p.Swap(i-1, j-1)
		i++
		j--
	}
}

func visit(p IntVector, level, n, k int) {
	level++
	p[k] = level
	if level == n {
		//      AddItem();     // to the list box
		p.Print()
	} else {
		for i := 0; i < n; i++ {
			if p[i] == 0 {
				visit(p, level, n, i)
			}
		}
	}
	level = level - 1
	p[k] = 0
}

func AllPerms(n int) {
	p := NewIntVector(n)
	level := -1
	visit(p, level, n, 0)
}

//    Heap's short and elegant algorithm is implemented as a recursive method HeapPermute [Levitin, p. 179]. It is invoked with HeapPermute(n).
// A. Levitin, Introduction to The Design & Analysis of Algorithms, Addison Wesley, 2003.
func heapPermute(p IntVector, n int) {
	if n == 1 {
		//        AddItem();
		p.Print()

	} else {
		for i := 0; i < n; i++ {
			heapPermute(p, n-1)
			if n%2 == 1 { // if n is odd
				p.Swap(0, n-1)
			} else { // if n is even
				p.Swap(i, n-1)
			}
		}
	}
}

func AllPermsHeap(n int) {
	p := NewIntVector(n)
	p.Order()
	heapPermute(p, n)
}

func (perm IntVector) Rank(elem int) (rank int) {
	rank = -999
	smp := len(perm)
	for i := 0; i < smp; i++ {
		if perm[i] == elem { // find rank of sample i
			rank = i
			break
		}
	}
	if rank == -999 {
		panic("element not found")
	}
	return
}
