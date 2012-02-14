// Chao distance
// Chao's index (Ecol. Lett. 8, 148-159; 2005) tries to take into
// account the number of unseen shared species using Chao's method for
// estimating the number of unseen species.
// Chao, A., Chazdon, R. L., Colwell, R. K. and Shen, T. (2005). A new statistical approach for assessing similarity of species composition with incidence and abundance data. Ecology Letters 8, 148–159. 

package sim

import (
	. "go-eco.googlecode.com/hg/eco"
	"math"
)

// Chao distance matrix
// Chao et al. (2005)
// Algorithm inspired by R:vegan
func Chao_D(data *Matrix) *Matrix {
	var v float64

	rows := data.R
	cols := data.C
	out := NewMatrix(rows, rows)
	// check whether data are integers; if not, truncate them
	WarnIfNotCounts(data)
	TruncData(data)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {

			itot := 0.0
			jtot := 0.0
			ionce := 0.0
			jonce := 0.0
			itwice := 0.0
			jtwice := 0.0
			ishare := 0.0
			jshare := 0.0
			ishar1 := 0.0
			jshar1 := 0.0

			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)

				itot += x
				jtot += y
				if x > 0 && y > 0 {
					ishare += x
					jshare += y
					if math.Abs(y-1) < 0.01 {
						ishar1 += x
						jonce += 1
					} else if math.Abs(y-2) < 0.01 {
						jtwice += 1
					}
					if math.Abs(x-1) < 0.01 {
						jshar1 += y
						ionce += 1
					} else if math.Abs(x-2) < 0.01 {
						itwice += 1
					}
				}

			}

			uu := ishare / itot
			if ishar1 > 0 {
				if jonce < 1 {
					jonce = 1 // Never true if got here?
				}
				if jtwice < 1 {
					jtwice = 1
				}
				uu += (jtot - 1) / jtot * jonce / jtwice / 2.0 * ishar1 / itot
			}
			if uu > 1 {
				uu = 1
			}
			vv := jshare / jtot
			if jshar1 > 0 {
				if ionce < 1 {
					ionce = 1 // Is this never true?
				}
				if itwice < 1 {
					itwice = 1
				}
				vv += (itot - 1) / itot * ionce / itwice / 2.0 * jshar1 / jtot
			}
			if vv > 1 {
				vv = 1
			}
			if uu <= 0.0 || vv <= 0.0 {
				v = 1.0
			} else {
				v = 1.0 - uu*vv/(uu+vv-uu*vv)
			}
			if v < 0.0 {
				v = 0.0
			}
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

/*
`chao.jaccard` <-
function(x,y)
{ n<-sum(x)
  m<-sum(y)
  if (length(x[y>0&x==2])==0) f2plus<-1
  else f2plus<-length(x[y>0&x==2])
  p1<-sum(x[y>0]/n)
  p2<-((m-1)/m)*(length(x[y>0&x==1])/(2*f2plus))
  p3<-sum(x[y==0]/n)
  u<-p1+p2*p3
  if (u>1) u<-1
  if (length(y[x>0&y==2])==0) fplus2<-1
  else fplus2<-length(y[x>0&y==2])
  q1<-sum(y[x>0]/m)
  q2<-((n-1)/n)*(length(y[x>0&y==1])/(2*fplus2))
  q3<-sum(y[x==0]/m)
  v<-q1+q2*q3
  if (v>1) v<-1
  if (u==0&v==0) c.j<-0
  else c.j<-(u*v)/(u+v-(u*v))
  return(c.j)
}
*/
// Chao - Jaccard similarity matrix
// Chao’s Jaccard shared species estimators for use with incomplete datasets
// Chao et al. (2005)
// Algorithm inspired by R:fossil
func ChaoJaccard_S(data *Matrix) *Matrix {
	rows := data.R
	cols := data.C
	out := NewMatrix(rows, rows)

	// check whether data are integers; if not, truncate them
	WarnIfNotCounts(data)
	TruncData(data)
	c := 0.0
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			sumX := 0.0
			sumY := 0.0
			f2plus := 0.0
			fplus2 := 0.0
			len1 := 0.0
			len2 := 0.0
			p1 := 0.0
			p2 := 0.0
			p3 := 0.0
			q1 := 0.0
			q2 := 0.0
			q3 := 0.0

			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				sumX += x
				sumY += y
				if x == 2 && y > 0 {
					f2plus++
				}
				if x == 1 && y > 0 {
					len1++
				}
				if y > 0 {
					p1 += x
				}
				if y == 0 {
					p3 += x
				}
				if y == 2 && x > 0 {
					fplus2++
				}
				if y == 1 && x > 0 {
					len2++
				}
				if x > 0 {
					q1 += y
				}
				if x == 0 {
					q3 += y
				}
			}
			if f2plus == 0 {
				f2plus = 1
			}
			if fplus2 == 0 {
				fplus2 = 1
			}
			p1 /= sumX
			p2 = ((sumY - 1) / sumY) * (len1 / (2 * f2plus))
			p3 /= sumX
			u := p1 + p2*p3
			if u > 1 {
				u = 1
			}
			q1 /= sumY
			q2 = ((sumX - 1) / sumX) * (len2 / (2 * fplus2))
			q3 /= sumY
			v := q1 + q2*q3
			if v > 1 {
				v = 1
			}
			if u == 0 && v == 0 {
				c = 0
			} else {
				c = (u * v) / (u + v - (u * v))
			}
			out.Set(i, j, c)
			out.Set(j, i, c)
		}
	}
	return out
}

// Chao - Sorensen similarity matrix
// Chao’s Sorensen shared species estimators for use with incomplete datasets
// Chao et al. (2005)
// Algorithm inspired by R:fossil
func ChaoSorensen_S(data *Matrix) *Matrix {
	rows := data.R
	cols := data.C
	out := NewMatrix(rows, rows)

	// check whether data are integers; if not, truncate them
	WarnIfNotCounts(data)
	TruncData(data)
	c := 0.0
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			sumX := 0.0
			sumY := 0.0
			f2plus := 0.0
			fplus2 := 0.0
			len1 := 0.0
			len2 := 0.0
			p1 := 0.0
			p2 := 0.0
			p3 := 0.0
			q1 := 0.0
			q2 := 0.0
			q3 := 0.0

			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				sumX += x
				sumY += y
				if x == 2 && y > 0 {
					f2plus++
				}
				if x == 1 && y > 0 {
					len1++
				}
				if y > 0 {
					p1 += x
				}
				if y == 0 {
					p3 += x
				}
				if y == 2 && x > 0 {
					fplus2++
				}
				if y == 1 && x > 0 {
					len2++
				}
				if x > 0 {
					q1 += y
				}
				if x == 0 {
					q3 += y
				}
			}
			if f2plus == 0 {
				f2plus = 1
			}
			if fplus2 == 0 {
				fplus2 = 1
			}
			p1 /= sumX
			p2 = ((sumY - 1) / sumY) * (len1 / (2 * f2plus))
			p3 /= sumX
			u := p1 + p2*p3
			if u > 1 {
				u = 1
			}
			q1 /= sumY
			q2 = ((sumX - 1) / sumX) * (len2 / (2 * fplus2))
			q3 /= sumY
			v := q1 + q2*q3
			if v > 1 {
				v = 1
			}
			if u == 0 && v == 0 {
				c = 0
			} else {
				c = (2*u*v)/(u+v)
			}
			out.Set(i, j, c)
			out.Set(j, i, c)
		}
	}
	return out
}
