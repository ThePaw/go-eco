// For Beta-Binomial sampling model: Estimates the parameters a,b of beta distribution from expected proportion (pi), binomial denominator (m), and shape parameter (tau2). Solution (hopefully correct) of Exercise 4.17 of McCullagh & Nelder 1989, helped by Moore, Appl Stat 36, 8-14; 1987.
// thanks to Jari Oksanen, I think.
func betapara(pi, m, tau2 float64) (a, b float64) {
	t1 := tau2 * m
	t2 := t1 - m - tau2 + 1
	t3 := 1 / (1 + t1 - tau2)
	t4 := t2 * t3
	a = -t4 * pi
	b = t4 * (pi - 1)
	return
}


