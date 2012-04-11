// The Hoover index of inequality
// Edgar Malone Hoover jr. (1936) The Measurement of Industrial Localization, Review of Economics and Statistics, 18, No. 162-171
// For the formula, a notation[1] is used, where the amount N of quantiles only appears as upper border of summations. Thus, inequities can be computed for quantiles with different widths A. For example, E_i could be the income in the quantile #i and A_i could be the amount (absolute or relative) of earners in the quantile #i. E_\text{total} then would be the sum of incomes of all N quantiles and A_\text{total} would be the sum of the income earners in all N quantiles.

package div

import (
	"math"
)

func Hoover_D(a, e []float64, n int) float64 {

	// A total
	a_tot := 0.0
	for i := 0; i < n; i++ {
		a_tot += a[i]
	}

	// E total
	e_tot := 0.0
	for i := 0; i < n; i++ {
		e_tot += e[i]
	}

	y := 0.0
	for i := 0; i < n; i++ {
		y += math.Abs(e[i]/e_tot - a[i]/a_tot)
	}
	return y / 2
}
