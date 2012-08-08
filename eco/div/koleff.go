// Beta diversity indices from Koleff, Gaston & Lennon (2003)
// Koleff, P., Gaston, K.J. and Lennon, J.J. (2003) Measuring beta diversity for presence-absence data. Journal of Animal Ecology 72, 367–382.

package div

import (
	. "code.google.com/p/go-eco/eco"
	"math"
)

// β1 = β2 = β8 = β9 = β20 diversity index from Koleff, Gaston & Lennon (2003)
func Koleff1Beta(data *Matrix) *Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = GetABCD(data, i, j)
			v := (b+c)/(2*a+b+c)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// β3 diversity index from Koleff, Gaston & Lennon (2003)
func Koleff3Beta(data *Matrix) *Matrix {
	var (
		b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			_, b, _, _ = GetABCD(data, i, j)
			v := (b+c)/2
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// β4 diversity index from Koleff, Gaston & Lennon (2003)
func Koleff4Beta(data *Matrix) *Matrix {
	var (
		b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			_, b, c, _ = GetABCD(data, i, j)
			v := b+c
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// β5 diversity index from Koleff, Gaston & Lennon (2003)
func Koleff5Beta(data *Matrix) *Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = GetABCD(data, i, j)
			v := 2*b*c/((a+b+c)*(a+b+c)-2*b*c)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// β6 diversity index from Koleff, Gaston & Lennon (2003)
func Koleff6Beta(data *Matrix) *Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = GetABCD(data, i, j)
			v := math.Log(2*a+b+c)-2*a*math.Log(2)/(2*a+b+c)-((a+b)*math.Log(a+b)+(a+c)*math.Log(a+c))/(2*a+b+c)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// β7 diversity index from Koleff, Gaston & Lennon (2003)
func Koleff7Beta(data *Matrix) *Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = GetABCD(data, i, j)
			v := math.Exp(math.Log(2*a+b+c)-2*a*math.Log(2)/(2*a+b+c)-((a+b)*math.Log(a+b)+(a+c)*math.Log(a+c))/(2*a+b+c))-1
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// β10 diversity index from Koleff, Gaston & Lennon (2003)
func Koleff10Beta(data *Matrix) *Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = GetABCD(data, i, j)
			v := a/(a+b+c)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// β11 diversity index from Koleff, Gaston & Lennon (2003)
func Koleff11Beta(data *Matrix) *Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = GetABCD(data, i, j)
			v := 2*a/(2*a+b+c)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// β12 diversity index from Koleff, Gaston & Lennon (2003)
func Koleff12Beta(data *Matrix) *Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = GetABCD(data, i, j)
			v := (2*a+b+c)*(b+c)/(a+b+c)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// β13 diversity index from Koleff, Gaston & Lennon (2003)
func Koleff13Beta(data *Matrix) *Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = GetABCD(data, i, j)
			v := math.Min(b,c)/(math.Max(b,c)+a)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// β14 diversity index from Koleff, Gaston & Lennon (2003)
func Koleff14Beta(data *Matrix) *Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = GetABCD(data, i, j)
			v := (a*c+a*b+2*b*c)/(2*(a+b)*(a+c))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// β15 = β16 diversity index from Koleff, Gaston & Lennon (2003)
func Koleff15Beta(data *Matrix) *Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = GetABCD(data, i, j)
			v := (b+c)/(a+b+c)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// β17 diversity index from Koleff, Gaston & Lennon (2003)
func Koleff17Beta(data *Matrix) *Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = GetABCD(data, i, j)
			v := math.Min(b,c)/(a+b+c)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// β18 diversity index from Koleff, Gaston & Lennon (2003)
func Koleff18Beta(data *Matrix) *Matrix {
	var (
		b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			_, b, c, _ = GetABCD(data, i, j)
			v := (b+c)/2
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// β19 diversity index from Koleff, Gaston & Lennon (2003)
func Koleff19Beta(data *Matrix) *Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = GetABCD(data, i, j)
			v := 2*(b*c+1)/((a+b+c)*(a+b+c)+(a+b+c))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// β21 diversity index from Koleff, Gaston & Lennon (2003)
func Koleff21Beta(data *Matrix) *Matrix {
	var (
		a, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, _, c, _ = GetABCD(data, i, j)
			v := a/(a+c)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// β22 diversity index from Koleff, Gaston & Lennon (2003)
func Koleff22Beta(data *Matrix) *Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = GetABCD(data, i, j)
			v := math.Min(b,c)/(math.Min(b,c)+a)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// β23 diversity index from Koleff, Gaston & Lennon (2003)
func Koleff23Beta(data *Matrix) *Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = GetABCD(data, i, j)
			v := 2*math.Abs(b-c)/(2*a+b+c)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// β24 diversity index from Koleff, Gaston & Lennon (2003)
func Koleff24Beta(data *Matrix) *Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = GetABCD(data, i, j)
			v := (math.Log(2)-math.Log(2*a+b+c)+math.Log(a+b+c))/math.Log(2)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}


