// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package div

// Beta diversity indices from Koleff, Gaston & Lennon (2003). 
// Koleff, P., Gaston, K.J. and Lennon, J.J. (2003) Measuring beta diversity for presence-absence data. Journal of Animal Ecology 72, 367–382. 

import (
	"code.google.com/p/go-eco/eco/aux"
	"math"
)

// Koleff1Div returns vector of Koleff β1 diversities. 
// β1 = β2 = β8 = β9 = β20 diversity index from Koleff, Gaston & Lennon (2003). 
func Koleff1Div(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = aux.GetABCD(data, i, j)
			v := (b + c) / (2*a + b + c)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Koleff3Div returns vector of Koleff β3 diversities. 
// Koleff, Gaston & Lennon (2003). 
func Koleff3Div(data *aux.Matrix) *aux.Matrix {
	var (
		b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			_, b, _, _ = aux.GetABCD(data, i, j)
			v := (b + c) / 2
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Koleff4Div returns vector of Koleff β4 diversities. 
// Koleff, Gaston & Lennon (2003). 
func Koleff4Div(data *aux.Matrix) *aux.Matrix {
	var (
		b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			_, b, c, _ = aux.GetABCD(data, i, j)
			v := b + c
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Koleff5Div returns vector of Koleff β5 diversities. 
// Koleff, Gaston & Lennon (2003). 
func Koleff5Div(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = aux.GetABCD(data, i, j)
			v := 2 * b * c / ((a+b+c)*(a+b+c) - 2*b*c)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Koleff6Div returns vector of Koleff β6 diversities. 
// Koleff, Gaston & Lennon (2003). 
func Koleff6Div(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = aux.GetABCD(data, i, j)
			v := math.Log(2*a+b+c) - 2*a*math.Log(2)/(2*a+b+c) - ((a+b)*math.Log(a+b)+(a+c)*math.Log(a+c))/(2*a+b+c)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Koleff7Div returns vector of Koleff β7 diversities. 
// Koleff, Gaston & Lennon (2003). 
func Koleff7Div(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = aux.GetABCD(data, i, j)
			v := math.Exp(math.Log(2*a+b+c)-2*a*math.Log(2)/(2*a+b+c)-((a+b)*math.Log(a+b)+(a+c)*math.Log(a+c))/(2*a+b+c)) - 1
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Koleff10Div returns vector of Koleff β10 diversities. 
// Koleff, Gaston & Lennon (2003). 
func Koleff10Div(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = aux.GetABCD(data, i, j)
			v := a / (a + b + c)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Koleff11Div returns vector of Koleff β11 diversities. 
// Koleff, Gaston & Lennon (2003). 
func Koleff11Div(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = aux.GetABCD(data, i, j)
			v := 2 * a / (2*a + b + c)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Koleff12Div returns vector of Koleff β12 diversities. 
// Koleff, Gaston & Lennon (2003). 
func Koleff12Div(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = aux.GetABCD(data, i, j)
			v := (2*a + b + c) * (b + c) / (a + b + c)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Koleff13Div returns vector of Koleff β13 diversities. 
// Koleff, Gaston & Lennon (2003). 
func Koleff13Div(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = aux.GetABCD(data, i, j)
			v := math.Min(b, c) / (math.Max(b, c) + a)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Koleff14Div returns vector of Koleff β14 diversities. 
// Koleff, Gaston & Lennon (2003). 
func Koleff14Div(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = aux.GetABCD(data, i, j)
			v := (a*c + a*b + 2*b*c) / (2 * (a + b) * (a + c))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Koleff15Div returns vector of Koleff β15 diversities. 
// Koleff, Gaston & Lennon (2003). 
// β15 = β16
func Koleff15Div(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = aux.GetABCD(data, i, j)
			v := (b + c) / (a + b + c)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Koleff17Div returns vector of Koleff β17 diversities. 
// Koleff, Gaston & Lennon (2003). 
func Koleff17Div(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = aux.GetABCD(data, i, j)
			v := math.Min(b, c) / (a + b + c)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Koleff18Div returns vector of Koleff β18 diversities. 
// Koleff, Gaston & Lennon (2003). 
func Koleff18Div(data *aux.Matrix) *aux.Matrix {
	var (
		b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			_, b, c, _ = aux.GetABCD(data, i, j)
			v := (b + c) / 2
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Koleff19Div returns vector of Koleff β19 diversities. 
// Koleff, Gaston & Lennon (2003). 
func Koleff19Div(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = aux.GetABCD(data, i, j)
			v := 2 * (b*c + 1) / ((a+b+c)*(a+b+c) + (a + b + c))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Koleff21Div returns vector of Koleff β21 diversities. 
// Koleff, Gaston & Lennon (2003). 
func Koleff21Div(data *aux.Matrix) *aux.Matrix {
	var (
		a, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, _, c, _ = aux.GetABCD(data, i, j)
			v := a / (a + c)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Koleff22Div returns vector of Koleff β22 diversities. 
// Koleff, Gaston & Lennon (2003). 
func Koleff22Div(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = aux.GetABCD(data, i, j)
			v := math.Min(b, c) / (math.Min(b, c) + a)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Koleff23Div returns vector of Koleff β23 diversities. 
// Koleff, Gaston & Lennon (2003). 
func Koleff23Div(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = aux.GetABCD(data, i, j)
			v := 2 * math.Abs(b-c) / (2*a + b + c)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Koleff24Div returns vector of Koleff β24 diversities. 
// Koleff, Gaston & Lennon (2003). 
func Koleff24Div(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = aux.GetABCD(data, i, j)
			v := (math.Log(2) - math.Log(2*a+b+c) + math.Log(a+b+c)) / math.Log(2)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
