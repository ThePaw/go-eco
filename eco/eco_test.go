package eco

import (
	"testing"
//	"fmt"
	. "gomatrix.googlecode.com/hg/matrix"
)

// test against known values
func TestEuclid(t *testing.T) {
	var (
		data, out *DenseMatrix
		d float64 = 1.7320508075688771
		s float64 = 0.36602540378443865
	)

	data = Zeros(2, 3) 
	data.Set(0, 0, 0)
	data.Set(0, 1, 0)
	data.Set(0, 2, 0)
	data.Set(1, 0, 1)
	data.Set(1, 1, 1)
	data.Set(1, 2, 1)

	out = Euclid_D(data)

	if !check(out.Get(0, 0), out.Get(1, 1)){
		t.Error()
	}
	if !check(out.Get(0, 0), 0.0){
		t.Error()
	}
	if !check(out.Get(0, 1), out.Get(1, 0)){
		t.Error()
	}
	if !check(out.Get(0, 1), d){
		t.Error()
	}

	out = Euclid_S(data)

	if !check(out.Get(0, 0), out.Get(1, 1)){
		t.Error()
	}
	if !check(out.Get(0, 0), 1.0){
		t.Error()
	}
	if !check(out.Get(0, 1), out.Get(1, 0)){
		t.Error()
	}
	if !check(out.Get(0, 1), s){
		t.Error()
	}

}


