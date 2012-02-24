package div

import (
	"fmt"
	"testing"
)


// Shannon test against R:vegan
func TestShannon(t *testing.T) {
	fmt.Println("Shannon test against R:vegan, float data")
	data := GetData()
	out := Shannon(data, 'e', false)

	//known diversities
	known := [...]float64{2.294555, 2.294458, 2.286693, 2.294462, 2.297777, 2.290620}

	rows := data.R

	// check
	for i := 0; i < rows; i++ {
		x := out.Get(i)
		y := known[i]

		if !check(x, y) {
			t.Error()
			fmt.Println(i, x, y)
		}
	}

	fmt.Println("Shannon test against R:vegan, int data, ln")
	data = GetCatData()
	out = Shannon(data, 'e', false)

	//known diversities
	known = [...]float64{4.448998, 4.381241, 4.365418, 4.423159, 4.444995, 4.410991}

	rows = data.R

	// check
	for i := 0; i < rows; i++ {
		x := out.Get(i)
		y := known[i]

		if !check(x, y) {
			t.Error()
			fmt.Println(i, x, y)
		}
	}

	fmt.Println("Shannon test against R:vegan, int data, log2")
	data = GetCatData()
	out = Shannon(data, '2', false)

	//known diversities
	known = [...]float64{6.418548, 6.320795, 6.297966, 6.381270, 6.412773, 6.363715}

	rows = data.R

	// check
	for i := 0; i < rows; i++ {
		x := out.Get(i)
		y := known[i]

		if !check(x, y) {
			t.Error()
			fmt.Println(i, x, y)
		}
	}

	fmt.Println("Shannon test against R:vegan, int data, log10")
	data = GetCatData()
	out = Shannon(data, 'd', false)

	//known diversities
	known = [...]float64{1.932175, 1.902749, 1.895877, 1.920954, 1.930437, 1.915669}

	rows = data.R

	// check
	for i := 0; i < rows; i++ {
		x := out.Get(i)
		y := known[i]

		if !check(x, y) {
			t.Error()
			fmt.Println(i, x, y)
		}
	}
}

// Simpson test against R:vegan
func TestSimpson(t *testing.T) {
	fmt.Println("Simpson test against R:vegan, complement type")
	data := GetCatData()
	out := Simpson(data, 'c', false)

	//known diversities
	known := [...]float64{0.9870751, 0.9860367, 0.9857834, 0.9868089, 0.9871495, 0.9866519}

	rows := data.R

	// check
	for i := 0; i < rows; i++ {
		x := out.Get(i)
		y := known[i]

		if !check(x, y) {
			t.Error()
			fmt.Println(i, x, y)
		}
	}

	fmt.Println("Simpson test against R:vegan, inverse type")
	data = GetCatData()
	out = Simpson(data, 'i', false)

	//known diversities
	known = [...]float64{77.36975, 71.61625, 70.34054, 75.80850, 77.81822, 74.91707}

	rows = data.R

	// check
	for i := 0; i < rows; i++ {
		x := out.Get(i)
		y := known[i]

		if !check(x, y) {
			t.Error()
			fmt.Println(i, x, y)
		}
	}
}

// Atkinson test against R:ineq
func TestAtkinson(t *testing.T) {
	fmt.Println("Atkinson test against R:ineq, epsilon =1.0")
	data := GetData()
	out := Atkinson(data, 1.0)

	//known inequalities
	known := [...]float64{0.00820619, 0.008100257, 0.01597159, 0.008376986, 0.004774476, 0.01199175}

	rows := data.R

	// check
	for i := 0; i < rows; i++ {
		x := out.Get(i)
		y := known[i]

		if !check(x, y) {
			t.Error()
			fmt.Println(i, x, y)
		}
	}

	fmt.Println("Atkinson test against R:ineq, epsilon =0.5")
	out = Atkinson(data, 0.5)

	//known inequalities
	known = [...]float64{0.004060306, 0.004058885, 0.00797419, 0.004124863, 0.002395535, 0.005992798}

	// check
	for i := 0; i < rows; i++ {
		x := out.Get(i)
		y := known[i]

		if !check(x, y) {
			t.Error()
			fmt.Println(i, x, y)
		}
	}
}

// RS test against R:ineq
func TestRS(t *testing.T) {
	fmt.Println("Ricci-Schutz test against R:ineq")
	data := GetData()
	out := RicciSchutz_D(data)

	//known inequalities
	known := [...]float64{0.05269138,0.05822897,0.07978667,0.05046708,0.03551596,0.06686520}

	// check
	rows := data.R
	for i := 0; i < rows; i++ {
		x := out.Get(i)
		y := known[i]

		if !check(x, y) {
			t.Error()
			fmt.Println(i, x, y)
		}
	}
}

// Gini test against R:ineq
func TestGini(t *testing.T) {
	fmt.Println("Gini test against R:ineq")
	data := GetData()
	out := Gini_D(data)

	//known inequalities
	known := [...]float64{0.07081562,0.07197870,0.10089229,0.06977700,0.05244517,0.08840777}

	// check
	rows := data.R
	for i := 0; i < rows; i++ {
		x := out.Get(i)
		y := known[i]

		if !check(x, y) {
			t.Error()
			fmt.Println(i, x, y)
		}
	}
}

// Kolm test against R:ineq
func TestKolm(t *testing.T) {
	fmt.Println("Kolm test against R:ineq, parameter =1.0")
	data := GetData()
	out := Kolm_D(data, 1.0)

	//known inequalities
	known := [...]float64{0.2803687,0.2930953,0.4706786,0.2902386,0.1845231,0.4090321}

	rows := data.R

	// check
	for i := 0; i < rows; i++ {
		x := out.Get(i)
		y := known[i]

		if !check(x, y) {
			t.Error()
			fmt.Println(i, x, y)
		}
	}

	fmt.Println("Kolm test against R:ineq, parameter =0.5")
	out = Kolm_D(data, 0.5)

	//known inequalities
	known = [...]float64{0.13996084,0.15756255,0.25776774,0.14034360,0.09573755,0.22034813}

	// check
	for i := 0; i < rows; i++ {
		x := out.Get(i)
		y := known[i]

		if !check(x, y) {
			t.Error()
			fmt.Println(i, x, y)
		}
	}
}

