// Copyright 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package ser

import (
	"math/rand"
)

func proposePerm1(p IntVector) {
	rows := p.Len()
	a := rand.Intn(rows)
	b := rand.Intn(rows)
	for b == a {
		b = rand.Intn(rows)
	}
	c := rand.Intn(rows)
	for c == a || c == b {
		c = rand.Intn(rows)
	}

	x := rand.Float64()
	switch {
	case x < 0.5:
		//    swap
		p.Swap(a, b)
	default:
		// invert
		p.InvertFromTo(a, b)
	}
}

func proposePerm2(p IntVector) {
	rows := p.Len()
	c := rand.Intn(rows)
	b := rand.Intn(c)
	a := rand.Intn(b)

	x := rand.Float64()
	switch {
	case x < 0.091:
		//    swap
		p.Swap(a, b)
	case x < 0.182:
		// invert
		p.InvertFromTo(a, b)
	case x < 0.273:
		p.InvertTo(a)
	case x < 0.364:
		p.InsideOut(a)
	case x < 0.455:
		p.InsertAt(a, b)
	case x < 0.546:
		p.InvertHeadAndTail(a, b)
	case x < 0.637:
		p.Displace(a, b, c)
	case x < 0.728:
		p.DisplaceInv(a, b, c)
	case x < 0.818:
		p.Scramble3(a)
	case x < 0.908:
		p.Scramble4(a)
	default:
		p.TwoPointSwapInv(a, b)
	}
}

func proposePerm(p IntVector) {
	var a, b, c, d int
	rows := p.Len()

	n := 14
	x := rand.Intn(n)

	switch {
	case x < 5:
		a = rand.Intn(rows)
	case x < 10:
		b = rand.Intn(rows-1)+1 // avoid zero
		a = rand.Intn(b)
	case x < 13:
		c = rand.Intn(rows-2)+2
		b = rand.Intn(c-1)+1
		a = rand.Intn(b)
	default:
		d = rand.Intn(rows-3)+3
		c = rand.Intn(d-2)+2
		b = rand.Intn(c-1)+1
		a = rand.Intn(b)
	}

	switch {
	case x == 0:
		p.InvertFrom(a)
	case x == 1:
		p.InvertTo(a)
	case x == 2:
		p.InsideOut(a)
	case x == 3:
		p.Scramble3(a)
	case x == 4:
		p.Scramble4(a)
	case x == 5:
		p.Swap(a, b)
	case x == 6:
		p.InvertFromTo(a, b)
	case x == 7:
		p.InsertAt(a, b)
	case x == 8:
		p.InvertHeadAndTail(a, b)
	case x == 9:
		p.TwoPointSwapInv(a, b)
	case x == 10:
		p.Displace(a, b, c)
	case x == 11:
		p.DisplaceInv(a, b, c)
	case x == 12:
		p.ThreePointExch(a, b, c)
	case x == 13:
		p.FourPointExch(a, b, c, d)
	default:
	}
}
