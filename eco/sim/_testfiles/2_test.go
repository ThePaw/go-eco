package sim

import (
	"fmt"
	. "go-eco.googlecode.com/hg/eco"
	"testing"
)

// Euclidean test against R:vegan
func TestEuclid(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("Euclidean test against R:vegan")
	data = GetData()
	out = Euclid_D(data)
	//known distances
	dist := [...]float64{0.00000000, 3.80120700, 3.33984600, 3.67449900, 4.27450000, 3.56335800,
		3.80120700, 0.00000000, 5.07576900, 3.79167000, 3.46983200, 4.70733000,
		3.33984600, 5.07576900, 0.00000000, 4.74485000, 3.81491100, 5.02369900,
		3.67449900, 3.79167000, 4.74485000, 0.00000000, 3.57503700, 4.42328400,
		4.27450000, 3.46983200, 3.81491100, 3.57503700, 0.00000000, 3.94087000,
		3.56335800, 4.70733000, 5.02369900, 4.42328400, 3.94087000, 0.00000000}
	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			if !check(out.Get(i, j), known.Get(i, j)) {
				t.Error()
			}

		}
	}
}

// Manhattan  test against R:vegan
func TestManhattan(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("Manhattan test against R:vegan")
	data = GetData()
	out = Manhattan_D(data)

	//known distances
	dist := [...]float64{0.000000, 8.924206, 7.606877, 10.411230, 9.159628, 9.836410,
		8.924206, 0.000000, 12.352389, 11.087330, 9.874214, 11.264396,
		7.606877, 12.352389, 0.000000, 14.220567, 10.124621, 14.126321,
		10.41123000, 11.087330, 14.220567, 0.000000, 7.966832, 11.625520,
		9.159628, 9.874214, 10.124621, 7.966832, 0.000000, 10.386242,
		9.836410, 11.264396, 14.126321, 11.625520, 10.386242, 0.000000}
	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			if !check(out.Get(i, j), known.Get(i, j)) {
				t.Error()
			}

		}
	}
}

// Canberra Scaled test against R:vegan
func TestCanberraSc(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("Canberra Scaled test against R:vegan")
	data = GetData()
	out = CanberraSc_D(data)

	//known distances
	dist := [...]float64{0.00000000, 0.07305403, 0.06705955, 0.09046813, 0.07580495, 0.08021418,
		0.07305403, 0.00000000, 0.10369549, 0.09178595, 0.07764712, 0.09047445,
		0.06705955, 0.10369549, 0.00000000, 0.12332237, 0.08709588, 0.11657391,
		0.09046813, 0.09178595, 0.12332237, 0.00000000, 0.06573803, 0.09637591,
		0.07580495, 0.07764712, 0.08709588, 0.06573803, 0.00000000, 0.08161942,
		0.08021418, 0.09047445, 0.11657391, 0.09637591, 0.08161942, 0.00000000}
	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			if !check(out.Get(i, j), known.Get(i, j)) {
				t.Error()
			}

		}
	}
}

/* FAILED !!
// Gower test against R:vegan
func TestGower(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	data = GetData()
	out = Gower_D(data)

//known distances
	dist := [...]float64{0.00000000,0.34582040,0.32576220,0.48013640,0.39378850,0.43220160,
0.34582040,0.00000000,0.50612320,0.51417230,0.44173250,0.47383420,
0.32576220,0.50612320,0.00000000,0.62548610,0.44607840,0.61632530,
0.48013640,0.51417230,0.62548610,0.00000000,0.31229460,0.47949750,
0.39378850,0.44173250,0.44607840,0.31229460,0.00000000,0.41614630,
0.43220160,0.47383420,0.61632530,0.47949750,0.41614630,0.00000000}
	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

// check
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			if !check(out.Get(i, j), known.Get(i, j)) {
				t.Error()
				fmt.Println(out.Get(i, j), known.Get(i, j))
			}

		}
	}
}
*/

// Raup-Crick test against R:vegan - big data
func TestRaupCrickBool2(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("Raup-Crick test against R:vegan, big data")
	data = GetBoolData()
	out = RaupCrickBool_D(data)

	//known distances
	dist := [...]float64{0.00000000, 0.83137125, 0.62972983, 0.91365497, 0.64752313, 0.42024290,
		0.83137125, 0.00000000, 0.94852481, 0.27939267, 0.50317018, 0.91956864,
		0.62972983, 0.94852481, 0.00000000, 0.66445339, 0.03635185, 0.34442637,
		0.91365497, 0.27939267, 0.66445339, 0.00000000, 0.94649405, 0.72578724,
		0.64752313, 0.50317018, 0.03635185, 0.94649405, 0.00000000, 0.78810701,
		0.42024290, 0.91956864, 0.34442637, 0.72578724, 0.78810701, 0.00000000}
	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			if !check(out.Get(i, j), known.Get(i, j)) {
				t.Error()
				fmt.Println(out.Get(i, j), known.Get(i, j))
			}
		}
	}
}

// Raup-Crick test against R:vegan, smaller data
func TestRaupCrickBool(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("Raup-Crick test against R:vegan, smaller data")
	data = GetBoolData2()
	out = RaupCrickBool_D(data)

	//known distances
	dist := [...]float64{0.00000000, 0.45238095, 0.66666667, 0.11904762, 1.00000000, 0.97619048,
		0.45238095, 0.00000000, 0.83333333, 0.07142857, 0.88095238, 0.97619048,
		0.66666667, 0.83333333, 0.00000000, 0.66666667, 0.66666667, 1.00000000,
		0.11904762, 0.07142857, 0.66666667, 0.00000000, 1.00000000, 0.97619048,
		1.00000000, 0.88095238, 0.66666667, 1.00000000, 0.00000000, 0.73809524,
		0.97619048, 0.97619048, 1.00000000, 0.97619048, 0.73809524, 0.00000000}

	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			if !check(out.Get(i, j), known.Get(i, j)) {
				t.Error()
				fmt.Println(out.Get(i, j), known.Get(i, j))
			}

		}
	}
}

// Mountford test against R:vegan
func TestMountford(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("Mountford test against R:vegan, big data")
	data = GetBoolData()
	out = Mountford_D(data)

	//known distances
	dist := [...]float64{0.0000000, 0.9748605, 0.9720504, 0.9766586, 0.9717896, 0.9684308,
		0.9748605, 0.0000000, 0.9801731, 0.9687023, 0.9719189, 0.9784538,
		0.9720504, 0.9801731, 0.0000000, 0.9745933, 0.9603015, 0.9695784,
		0.9766586, 0.9687023, 0.9745933, 0.0000000, 0.9795951, 0.9747592,
		0.9717896, 0.9719189, 0.9603015, 0.9795951, 0.0000000, 0.9754804,
		0.9684308, 0.9784538, 0.9695784, 0.9747592, 0.9754804, 0.0000000}
	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			if !check(out.Get(i, j), known.Get(i, j)) {
				t.Error()
				fmt.Println(out.Get(i, j), known.Get(i, j))
			}

		}
	}
}

// Mountford test against R:vegan, smaller data
func TestMountford2(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("Mountford test against R:vegan, small data")
	data = GetBoolData2()
	out = Mountford_D(data)

	//known distances
	dist := [...]float64{0.00000000, 0.53504158, 0.67824320, 0.01189135, 0.88138692, 0.85054888,
		0.53504158, 0.00000000, 0.83170094, 0.00000000, 0.79715738, 0.90738441,
		0.67824320, 0.83170094, 0.00000000, 0.67824320, 0.67824320, 1.00000000,
		0.01189135, 0.00000000, 0.67824320, 0.00000000, 0.88138692, 0.85054888,
		0.88138692, 0.79715738, 0.67824320, 0.88138692, 0.00000000, 0.69373111,
		0.85054888, 0.90738441, 1.00000000, 0.85054888, 0.69373111, 0.00000000}
	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			if !check(out.Get(i, j), known.Get(i, j)) {
				t.Error()
				fmt.Println(i, j, out.Get(i, j), known.Get(i, j))
			}
		}
	}
}

// Morisita test against R:vegan
func TestMorisita(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("Morisita test against R:vegan")
	data = GetCatData()
	out = Morisita_D(data)
	//known distances
	dist := [...]float64{0.0000000, 0.2630106, 0.2442520, 0.1880951, 0.2171524, 0.2220582,
		0.2630106, 0.0000000, 0.2948483, 0.2183969, 0.2657005, 0.2016453,
		0.2442520, 0.2948483, 0.0000000, 0.2815873, 0.2348226, 0.2100823,
		0.1880951, 0.2183969, 0.2815873, 0.0000000, 0.1898959, 0.2235514,
		0.2171524, 0.2657005, 0.2348226, 0.1898959, 0.0000000, 0.2046018,
		0.2220582, 0.2016453, 0.2100823, 0.2235514, 0.2046018, 0.0000000}

	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			if !check(out.Get(i, j), known.Get(i, j)) {
				t.Error()
				fmt.Println(i, j, out.Get(i, j), known.Get(i, j))
			}

		}
	}
}

// Horn-Morisita test against R:vegan
func TestHornMorisita(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("Horn-Morisita test against R:vegan")
	data = GetCatData()
	out = HornMorisita_D(data)
	//known distances
	dist := [...]float64{0.0000000, 0.2852076, 0.2676365, 0.2122724, 0.2410327, 0.2454808,
		0.2852076, 0.0000000, 0.3159566, 0.2408891, 0.2873345, 0.2248776,
		0.2676365, 0.3159566, 0.0000000, 0.3028561, 0.2580032, 0.2337172,
		0.2122724, 0.2408891, 0.3028561, 0.0000000, 0.2134713, 0.2458634,
		0.2410327, 0.2873345, 0.2580032, 0.2134713, 0.0000000, 0.2280157,
		0.2454808, 0.2248776, 0.2337172, 0.2458634, 0.2280157, 0.0000000}

	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			if !check(out.Get(i, j), known.Get(i, j)) {
				t.Error()
				fmt.Println(i, j, out.Get(i, j), known.Get(i, j))
			}

		}
	}
}

// Horn-Morisita test against R:vegan
func TestBrayCurtis(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("Bray-Curtis test against R:vegan")
	data = GetData()
	out = BrayCurtis_D(data)
	//known distances
	dist := [...]float64{0.00000000, 0.07300309, 0.06501591, 0.08912230, 0.07448023, 0.08149640,
		0.07300309, 0.00000000, 0.10144119, 0.09118747, 0.07729363, 0.08978032,
		0.06501591, 0.10144119, 0.00000000, 0.12222828, 0.08264632, 0.11750175,
		0.08912230, 0.09118747, 0.12222828, 0.00000000, 0.06512854, 0.09684579,
		0.07448023, 0.07729363, 0.08264632, 0.06512854, 0.00000000, 0.08229814,
		0.08149640, 0.08978032, 0.11750175, 0.09684579, 0.08229814, 0.00000000}

	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			if !check(out.Get(i, j), known.Get(i, j)) {
				t.Error()
				fmt.Println(i, j, out.Get(i, j), known.Get(i, j))
			}

		}
	}
}

// Chao test against R:vegan, smaller data
func TestChao(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("Chao test against R:vegan, small data")
	data = GetBoolData2()
	out = Chao_D(data)

	//known distances
	dist := [...]float64{0.0000000, 0.0000000, 0.4444444, 0.0000000, 0.5600000, 0.5074627,
		0.0000000, 0.0000000, 0.7608696, 0.0000000, 0.4460432, 0.8179669,
		0.4444444, 0.7608696, 0.0000000, 0.4444444, 0.4444444, 1.0000000,
		0.0000000, 0.0000000, 0.4444444, 0.0000000, 0.5600000, 0.5074627,
		0.5600000, 0.4460432, 0.4444444, 0.5600000, 0.0000000, 0.0000000,
		0.5074627, 0.8179669, 1.0000000, 0.5074627, 0.0000000, 0.0000000}

	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			if !check(out.Get(i, j), known.Get(i, j)) {
				t.Error()
				fmt.Println(i, j, out.Get(i, j), known.Get(i, j))
			}
		}
	}
}

/*
// Millar vs. Binomial
// now obsolete, Millar reimplemented as Binomial
func TestMillarBinomial(t *testing.T) {
	var (
		data, out1, out2 *Matrix
	)

	fmt.Println("test whether Binomial_D == Millar_D")
	data = GetBoolData2()
	out1 = Binomial_D(data)
	out2 = Millar_D(data)
	rows := data.R
// check
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			x:=out1.Get(i, j)
			y:=out2.Get(i, j)

			if !check(x, y) {
				t.Error()
				fmt.Println(i, j, x, y)
			}
		}
	}
}


// Jaccard test against R:vegan
func TestJaccard(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("Jaccard test against R:vegan, small data")
	data = GetBoolData2()
	out = JaccardBool_D(data)

	//known distances
	dist := [...]float64{0.0000000, 0.5714286, 0.7142857, 0.2857143, 0.8000000, 0.7777778,
		0.5714286, 0.0000000, 0.8333333, 0.3333333, 0.7500000, 0.8750000,
		0.7142857, 0.8333333, 0.0000000, 0.7142857, 0.7142857, 1.0000000,
		0.2857143, 0.3333333, 0.7142857, 0.0000000, 0.8000000, 0.7777778,
		0.8000000, 0.7500000, 0.7142857, 0.8000000, 0.0000000, 0.6250000,
		0.7777778, 0.8750000, 1.0000000, 0.7777778, 0.6250000, 0.0000000}

	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			x := out.Get(i, j)
			y := known.Get(i, j)

			if !check(x, y) {
				t.Error()
				fmt.Println(i, j, x, y)
			}
		}
	}
}
*/

// Růžička test against R:vegan
func TestRůžička(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("Růžička test against R:vegan, small data")
	data = GetData()
	out = Ruzicka_D(data)

	//known distances
	dist := [...]float64{0, 0.1360725, 0.1220938, 0.1636589, 0.1386349, 0.1507104,
		0.1360725, 0, 0.1841972, 0.1671344, 0.1434959, 0.1647677,
		0.1220938, 0.1841972, 0, 0.2178314, 0.1526746, 0.2102936,
		0.1636589, 0.1671344, 0.2178314, 0, 0.1222924, 0.1765896,
		0.1386349, 0.1434959, 0.1526746, 0.1222924, 0, 0.1520803,
		0.1507104, 0.1647677, 0.2102936, 0.1765896, 0.1520803, 0}

	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			x := out.Get(i, j)
			y := known.Get(i, j)

			if !check(x, y) {
				t.Error()
				fmt.Println(i, j, x, y)
			}
		}
	}
}

// Sørensen test against R:vegan
func TestSørensen(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("Sørensen test against R:vegan, small data")
	data = GetBoolData2()
	out = SorensenBool_D(data)

	//known distances
	dist := [...]float64{0, 0.4, 0.5555556, 0.1666667, 0.6666667, 0.6363636,
		0.4, 0, 0.7142857, 0.2, 0.6, 0.7777778,
		0.5555556, 0.7142857, 0, 0.5555556, 0.5555556, 1,
		0.1666667, 0.2, 0.5555556, 0, 0.6666667, 0.6363636,
		0.6666667, 0.6, 0.5555556, 0.6666667, 0, 0.4545455,
		0.6363636, 0.7777778, 1, 0.6363636, 0.4545455, 0}
	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			x := out.Get(i, j)
			y := known.Get(i, j)

			if !check(x, y) {
				t.Error()
				fmt.Println(i, j, x, y)
			}
		}
	}
}

// Arrhenius test against R:vegan
func TestArrhenius(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("Arrhenius test against R:vegan, big data")
	data = GetBoolData()
	out = ArrheniusBool_D(data)

	//known distances
	dist := [...]float64{0.0000000, 0.6033410, 0.5802860, 0.6214884, 0.5711567, 0.5388661,
		0.6033410, 0.0000000, 0.6780719, 0.5647846, 0.5899117, 0.6520767,
		0.5802860, 0.6780719, 0.0000000, 0.6199690, 0.5025003, 0.5700122,
		0.6214884, 0.5647846, 0.6199690, 0.0000000, 0.6668737, 0.6141088,
		0.5711567, 0.5899117, 0.5025003, 0.6668737, 0.0000000, 0.6185709,
		0.5388661, 0.6520767, 0.5700122, 0.6141088, 0.6185709, 0.0000000}

	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			x := out.Get(i, j)
			y := known.Get(i, j)

			if !check(x, y) {
				t.Error()
				fmt.Println(i, j, x, y)
			}
		}
	}
}

/*
// Dice test against Jaccard: similarity
func TestDice(t *testing.T) {
	fmt.Println("Dice test against Jaccard: dissimilarity")
	data := GetBoolData()
	out := DiceBool_S(data)

	rows := data.R

	// check
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			x := out.Get(i, j)
			jacc_mx := JaccardBool_S(data)
			jacc := jacc_mx.Get(i, j)  // similarity
			y := 2 * jacc / (1 + jacc) // similarity
			if !check(x, y) {
				t.Error()
				fmt.Println(i, j, x, y)
			}
		}
	}
}

// Dice test against Jaccard: dissimilarity
func TestDice(t *testing.T) {
	fmt.Println("Dice test against Jaccard: dissimilarity")
	data := GetBoolData()
	out := DiceBool_D(data)


	rows := data.R

	// check
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			x:=out.Get(i, j)
			jacc_mx := JaccardBool_D(data)
			jacc := 1-jacc_mx.Get(i, j)  // similarity
			y:= 1-(2*jacc / (1 + jacc))  // dissimilarity
			if !check(x, y) {
				t.Error()
				fmt.Println(i, j, x, y)
			}
		}
	}
}
*/

// Sorensen test against R:simba

func TestSorensenS(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("Soerensen sim test against R:simba")
	data = GetBoolData()
	out = SorensenBool_S(data)

	//known similarities
	dist := [...]float64{1,0.4807692,0.5048544,0.4615385,0.5142857,0.5471698,
0.4807692,1,0.4,0.5208333,0.4948454,0.4285714,
0.5048544,0.4,1,0.4631579,0.5833333,0.5154639,
0.4615385,0.5208333,0.4631579,1,0.4123711,0.4693878,
0.5142857,0.4948454,0.5833333,0.4123711,1,0.4646465,
0.5471698,0.4285714,0.5154639,0.4693878,0.4646465,1}

	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			x := out.Get(i, j)
			y := known.Get(i, j)

			if !check(x, y) {
				t.Error()
				fmt.Println(i, j, x, y)
			}
		}
	}
}

// Jaccard test against R:simba
func TestJaccardS(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("Jaccard sim test against R:simba")
	data = GetBoolData()
	out = JaccardBool_S(data)

	//known similarities
	dist := [...]float64{1,0.3164557,0.3376623,0.3,0.3461538,0.3766234,
0.3164557,1,0.25,0.3521127,0.3287671,0.2727273,
0.3376623,0.25,1,0.3013699,0.4117647,0.3472222,
0.3,0.3521127,0.3013699,1,0.2597403,0.3066667,
0.3461538,0.3287671,0.4117647,0.2597403,1,0.3026316,
0.3766234,0.2727273,0.3472222,0.3066667,0.3026316,1}

	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			x := out.Get(i, j)
			y := known.Get(i, j)

			if !check(x, y) {
				t.Error()
				fmt.Println(i, j, x, y)
			}
		}
	}
}

// Ochiai test against R:simba
func TestOchiaiS(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("Ochiai sim test against R:simba")
	data = GetBoolData()
	out = OchiaiBool_S(data)

	//known similarities
	dist := [...]float64{1,0.482198,0.5067928,0.4629100,0.5154324,0.5480485,
0.482198,1,0.4000222,0.5208333,0.4948717,0.4286607,
0.5067928,0.4000222,1,0.4631836,0.58346,0.5157106,
0.4629100,0.5208333,0.4631836,1,0.4123930,0.4694855,
0.5154324,0.4948717,0.58346,0.4123930,1,0.4646702,
0.5480485,0.4286607,0.5157106,0.4694855,0.4646702,1}

	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			x := out.Get(i, j)
			y := known.Get(i, j)

			if !check(x, y) {
				t.Error()
				fmt.Println(i, j, x, y)
			}
		}
	}
}

// MountfordBool test against R:simba
func TestMountfordBoolS(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("MountfordBool sim test against R:simba")
	data = GetBoolData()
	out = MountfordBool_S(data)

	//known similarities
	dist := [...]float64{1,0.01801153,0.02010828,0.01666667,0.02035432,0.02296120,
0.01801153,1,0.01403768,0.02264493,0.02020202,0.01531729,
0.02010828,0.01403768,1,0.01816680,0.02919708,0.02197802,
0.01666667,0.02264493,0.01816680,1,0.01447178,0.01806756,
0.02035432,0.02020202,0.02919708,0.01447178,1,0.01753717,
0.02296120,0.01531729,0.02197802,0.01806756,0.01753717,1}

	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			x := out.Get(i, j)
			y := known.Get(i, j)

			if !check(x, y) {
				t.Error()
				fmt.Println(i, j, x, y)
			}
		}
	}
}

// WhittakerBool test against R:simba
func TestWhittakerBoolD(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("WhittakerBool dis test against R:simba")
	data = GetBoolData()
	out = WhittakerBool_D(data)

	//known similarities
	dist := [...]float64{0,0.4807692,0.5048544,0.4615385,0.5142857,0.5471698,
0.4807692,0,0.4,0.5208333,0.4948454,0.4285714,
0.5048544,0.4,0,0.4631579,0.5833333,0.5154639,
0.4615385,0.5208333,0.4631579,0,0.4123711,0.4693878,
0.5142857,0.4948454,0.5833333,0.4123711,0,0.4646465,
0.5471698,0.4285714,0.5154639,0.4693878,0.4646465,0}
	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			x := out.Get(i, j)
			y := known.Get(i, j)

			if !check(x, y) {
				t.Error()
				fmt.Println(i, j, x, y)
			}
		}
	}
}

// Raup-Crick1 test against R:vegan, smaller data
func TestRaupCrickBool1(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("Raup-Crick1 test against R:vegan, smaller data")
	data = GetBoolData2()
	out = RaupCrickBool1_S(data)

	//known distances
	dist := [...]float64{0.00000000, 0.45238095, 0.66666667, 0.11904762, 1.00000000, 0.97619048,
		0.45238095, 0.00000000, 0.83333333, 0.07142857, 0.88095238, 0.97619048,
		0.66666667, 0.83333333, 0.00000000, 0.66666667, 0.66666667, 1.00000000,
		0.11904762, 0.07142857, 0.66666667, 0.00000000, 1.00000000, 0.97619048,
		1.00000000, 0.88095238, 0.66666667, 1.00000000, 0.00000000, 0.73809524,
		0.97619048, 0.97619048, 1.00000000, 0.97619048, 0.73809524, 0.00000000}

	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			if !check(out.Get(i, j), known.Get(i, j)) {
				t.Error()
				fmt.Println(1-out.Get(i, j), known.Get(i, j))
			}

		}
	}
}


