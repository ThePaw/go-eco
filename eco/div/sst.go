// Sen-Shorrocks-Thon index of poverty

package div

import (
	"code.google.com/p/go-eco/eco/aux"
	"sort"
)

// Sen-Shorrocks-Thon index of poverty
// A poverty index proposed by Shorrocks (1995) based on the pioneering work of Sen (1976). 
// It has also received the name of modified Sen index in Shorrocks (1995) and Sen (1997). 
// As noted by Zheng (1997), this index is identical to the limit of Thon's modified Sen index (Thon, 1979 and 1983).
// Zheng, B. (1997). Aggregate Poverty Measures. Journal of Economic Surveys, 11, 123–162.
/*
Sen, A.K. (1976) "Poverty: An ordinal Approach to Measurement," Econometrica, 44, 219-231
Sen, A.K. (1997) "On Economic Inequality," Expanded edition (Oxford: Clarendon Press)
Shorrocks, A. F. (1997) "Revisiting the Sen poverty index," Econometrica, 63, 1225,1230
Thon, D. (1979) "On measuring poverty," Review of Income and Wealth, 25, 429-440
Thon, D. (1983) "A poverty measure," The Indian Economic Journal, 30, 55-70
Xu, K. (1998) "The statistical inference for the Sen-Shorrocks-Thon index of poverty intensity," Journal of Income Distribution, 8, 143-152
Xu, K. and L. Osberg (2001) "How to decompose the Sen-Shorrocks-Thon poverty index? A practitioner's Guide," Journal of Income Distribution
*/

func SST_D(data *aux.Matrix, k float64) *Vector {
	rows := data.R
	cols := data.C
	out := NewVector(rows)

	for i := 0; i < rows; i++ {
		// unload data row to slice
		arr := make([]float64, cols)
		arr = data.A[i*cols : i*cols+cols]
		sort.Float64s(arr)

		n := 0.0
		q := 0.0
		v := 0.0

		for j := 0; j < cols; j++ {
			x := arr[j]
			if x > 0.0 {
				n++
				if x < k {
					q++
					//					(2 * n - 2 * 1:q + 1) * (k - x2)/k
					v += (2*(n-float64(j)) + 1) * (k - x) / k
				}
			}
		}
		if q > 0 {
			v /= n * n
		}
		out.Set(i, v)
	}
	return out
}

/*
SST <- function(x, k)
{
  x2 <- sort(x[x < k])
  n <- length(x)
  q <- length(x2)
  if(q < 1) 0 else {
    sum((2 * n - 2 * 1:q + 1) * (k - x2)/k)/n^2
  }
}
*/
