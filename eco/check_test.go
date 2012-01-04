package eco

func check(x, y float64) bool {
	const acc float64 = 1e-6	// accuracy
	var z float64
		if x/y > 1.00 {
			z = y/x
		} else {
			z = x/y
		}
		if 1-z > acc  {
			return false
		}
		return true
}


