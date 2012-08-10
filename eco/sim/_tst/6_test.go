// Tests of boolean sim/dis measures against R:simba
// Warning! Simba does not return proper diagonal for sim/dis, so that "for j := i+1; j < rows; j++" is used. Also, simba calls everything "sim", although some are actually "dis".

package sim

import (
	"code.google.com/p/go-eco/eco/aux"
	"fmt"
	"testing"
)

// SorensenBool test against R:simba
func TestSorensenBool(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("SorensenBool_S test against R:simba")
	data = GetBoolData()
	out = SorensenBool_S(data)
	//known values
	dist := [...]float64{0.0000000, 0.4807692, 0.5048544, 0.4615385, 0.5142857, 0.5471698,
		0.4807692, 0.0000000, 0.4000000, 0.5208333, 0.4948454, 0.4285714,
		0.5048544, 0.4000000, 0.0000000, 0.4631579, 0.5833333, 0.5154639,
		0.4615385, 0.5208333, 0.4631579, 0.0000000, 0.4123711, 0.4693878,
		0.5142857, 0.4948454, 0.5833333, 0.4123711, 0.0000000, 0.4646465,
		0.5471698, 0.4285714, 0.5154639, 0.4693878, 0.4646465, 0.0000000}

	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			if !check(out.Get(i, j), known.Get(i, j)) {
				t.Error()
			}

		}
	}
}

// JaccardBool test against R:simba
func TestJaccardBool(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("JaccardBool_S test against R:simba")
	data = GetBoolData()
	out = JaccardBool_S(data)
	//known values
	dist := [...]float64{0.0000000, 0.3164557, 0.3376623, 0.3000000, 0.3461538, 0.3766234,
		0.3164557, 0.0000000, 0.2500000, 0.3521127, 0.3287671, 0.2727273,
		0.3376623, 0.2500000, 0.0000000, 0.3013699, 0.4117647, 0.3472222,
		0.3000000, 0.3521127, 0.3013699, 0.0000000, 0.2597403, 0.3066667,
		0.3461538, 0.3287671, 0.4117647, 0.2597403, 0.0000000, 0.3026316,
		0.3766234, 0.2727273, 0.3472222, 0.3066667, 0.3026316, 0.0000000}

	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			if !check(out.Get(i, j), known.Get(i, j)) {
				t.Error()
			}

		}
	}
}

// OchiaiBool test against R:simba
func TestOchiaiBool(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("OchiaiBool_S test against R:simba")
	data = GetBoolData()
	out = OchiaiBool_S(data)
	//known values
	dist := [...]float64{0.0000000, 0.4821980, 0.5067928, 0.4629100, 0.5154324, 0.5480485,
		0.4821980, 0.0000000, 0.4000222, 0.5208333, 0.4948717, 0.4286607,
		0.5067928, 0.4000222, 0.0000000, 0.4631836, 0.5834600, 0.5157106,
		0.4629100, 0.5208333, 0.4631836, 0.0000000, 0.4123930, 0.4694855,
		0.5154324, 0.4948717, 0.5834600, 0.4123930, 0.0000000, 0.4646702,
		0.5480485, 0.4286607, 0.5157106, 0.4694855, 0.4646702, 0.0000000}

	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			if !check(out.Get(i, j), known.Get(i, j)) {
				t.Error()
			}

		}
	}
}

// MountfordBool test against R:simba
func TestMountfordBool(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("MountfordBool_S test against R:simba")
	data = GetBoolData()
	out = MountfordBool_S(data)
	//known values
	dist := [...]float64{0.00000000, 0.01801153, 0.02010828, 0.01666667, 0.02035432, 0.02296120,
		0.01801153, 0.00000000, 0.01403768, 0.02264493, 0.02020202, 0.01531729,
		0.02010828, 0.01403768, 0.00000000, 0.01816680, 0.02919708, 0.02197802,
		0.01666667, 0.02264493, 0.01816680, 0.00000000, 0.01447178, 0.01806756,
		0.02035432, 0.02020202, 0.02919708, 0.01447178, 0.00000000, 0.01753717,
		0.02296120, 0.01531729, 0.02197802, 0.01806756, 0.01753717, 0.00000000}

	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			if !check(out.Get(i, j), known.Get(i, j)) {
				t.Error()
			}

		}
	}
}

// WhittakerBool test against R:simba
func TestWhittakerBool(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("WhittakerBool_D test against R:simba")
	data = GetBoolData()
	out = WhittakerBool_D(data)
	//known values
	dist := [...]float64{0.0000000, 0.5192308, 0.4951456, 0.5384615, 0.4857143, 0.4528302,
		0.5192308, 0.0000000, 0.6000000, 0.4791667, 0.5051546, 0.5714286,
		0.4951456, 0.6000000, 0.0000000, 0.5368421, 0.4166667, 0.4845361,
		0.5384615, 0.4791667, 0.5368421, 0.0000000, 0.5876289, 0.5306122,
		0.4857143, 0.5051546, 0.4166667, 0.5876289, 0.0000000, 0.5353535,
		0.4528302, 0.5714286, 0.4845361, 0.5306122, 0.5353535, 0.0000000}

	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			if !check(out.Get(i, j), known.Get(i, j)) {
				t.Error()
			}

		}
	}
}

// LandeBool test against R:simba
func TestLandeBool(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("LandeBool_D test against R:simba")
	data = GetBoolData()
	out = LandeBool_D(data)
	//known values
	dist := [...]float64{0.0, 27.0, 25.5, 28.0, 25.5, 24.0,
		27.0, 0.0, 28.5, 23.0, 24.5, 28.0,
		25.5, 28.5, 0.0, 25.5, 20.0, 23.5,
		28.0, 23.0, 25.5, 0.0, 28.5, 26.0,
		25.5, 24.5, 20.0, 28.5, 0.0, 26.5,
		24.0, 28.0, 23.5, 26.0, 26.5, 0.0}

	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			if !check(out.Get(i, j), known.Get(i, j)) {
				t.Error()
				fmt.Println(out.Get(i, j), known.Get(i, j))
			}

		}
	}
}

// WilsonShmidaBool test against R:simba
func TestWilsonShmidaBool(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("WilsonShmidaBool_D test against R:simba")
	data = GetBoolData()
	out = WilsonShmidaBool_D(data)
	//known values
	dist := [...]float64{0.0000000, 0.5192308, 0.4951456, 0.5384615, 0.4857143, 0.4528302,
		0.5192308, 0.0000000, 0.6000000, 0.4791667, 0.5051546, 0.5714286,
		0.4951456, 0.6000000, 0.0000000, 0.5368421, 0.4166667, 0.4845361,
		0.5384615, 0.4791667, 0.5368421, 0.0000000, 0.5876289, 0.5306122,
		0.4857143, 0.5051546, 0.4166667, 0.5876289, 0.0000000, 0.5353535,
		0.4528302, 0.5714286, 0.4845361, 0.5306122, 0.5353535, 0.0000000}

	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			if !check(out.Get(i, j), known.Get(i, j)) {
				t.Error()
				fmt.Println(out.Get(i, j), known.Get(i, j))
			}

		}
	}
}

// CoCoGastonBool test against R:simba
func TestCoCoGastonBool(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("CoCoGastonBool_D test against R:simba")
	data = GetBoolData()
	out = CoCoGastonBool_D(data)
	//known values
	dist := [...]float64{0.0000000, 0.6835443, 0.6623377, 0.7000000, 0.6538462, 0.6233766,
		0.6835443, 0.0000000, 0.7500000, 0.6478873, 0.6712329, 0.7272727,
		0.6623377, 0.7500000, 0.0000000, 0.6986301, 0.5882353, 0.6527778,
		0.7000000, 0.6478873, 0.6986301, 0.0000000, 0.7402597, 0.6933333,
		0.6538462, 0.6712329, 0.5882353, 0.7402597, 0.0000000, 0.6973684,
		0.6233766, 0.7272727, 0.6527778, 0.6933333, 0.6973684, 0.0000000}

	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			if !check(out.Get(i, j), known.Get(i, j)) {
				t.Error()
				fmt.Println(out.Get(i, j), known.Get(i, j))
			}

		}
	}
}

// MagurranBool test against R:simba
func TestMagurranBool(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("MagurranBool_D test against R:simba")
	data = GetBoolData()
	out = MagurranBool_D(data)
	//known values
	dist := [...]float64{0.00000, 71.08861, 68.22078, 72.80000, 68.65385, 66.07792,
		71.08861, 0.00000, 71.25000, 62.19718, 65.10959, 71.27273,
		68.22078, 71.25000, 0.00000, 66.36986, 56.47059, 63.31944,
		72.80000, 62.19718, 66.36986, 0.00000, 71.80519, 67.94667,
		68.65385, 65.10959, 56.47059, 71.80519, 0.00000, 69.03947,
		66.07792, 71.27273, 63.31944, 67.94667, 69.03947, 0.00000}

	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			if !check(out.Get(i, j), known.Get(i, j)) {
				t.Error()
				fmt.Println(out.Get(i, j), known.Get(i, j))
			}

		}
	}
}

// HarrisonBool test against R:simba
func TestHarrisonBool(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("HarrisonBool_D test against R:simba")
	data = GetBoolData()
	out = HarrisonBool_D(data)
	//known values
	dist := [...]float64{0.0000000, 0.4107143, 0.3750000, 0.4285714, 0.3928571, 0.375,
		0.4107143, 0.0000000, 0.5833333, 0.4791667, 0.4897959, 0.540,
		0.3750000, 0.5833333, 0.0000000, 0.5208333, 0.3877551, 0.440,
		0.4285714, 0.4791667, 0.5208333, 0.0000000, 0.5714286, 0.500,
		0.3928571, 0.4897959, 0.3877551, 0.5714286, 0.0000000, 0.520,
		0.3750000, 0.5400000, 0.4400000, 0.5000000, 0.5200000, 0.000}

	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			if !check(out.Get(i, j), known.Get(i, j)) {
				t.Error()
				fmt.Println(out.Get(i, j), known.Get(i, j))
			}

		}
	}
}

// CodyBool test against R:simba
func TestCodyBool(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("CodyBool_D test against R:simba")
	data = GetBoolData()
	out = CodyBool_D(data)
	//known values
	dist := [...]float64{0.0000000, 0.5163690, 0.4912614, 0.5357143, 0.4834184, 0.4510714,
		0.5163690, 0.0000000, 0.5999557, 0.4791667, 0.5051020, 0.5712500,
		0.4912614, 0.5999557, 0.0000000, 0.5367908, 0.4164134, 0.4840426,
		0.5357143, 0.4791667, 0.5367908, 0.0000000, 0.5875850, 0.5304167,
		0.4834184, 0.5051020, 0.4164134, 0.5875850, 0.0000000, 0.5353061,
		0.4510714, 0.5712500, 0.4840426, 0.5304167, 0.5353061, 0.0000000}

	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			if !check(out.Get(i, j), known.Get(i, j)) {
				t.Error()
				fmt.Println(out.Get(i, j), known.Get(i, j))
			}

		}
	}
}

// Williams1Bool test against R:simba
func TestWilliams1Bool(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("Williams1Bool_D test against R:simba")
	data = GetBoolData()
	out = Williams1Bool_D(data)
	//known values
	dist := [...]float64{0.0000000, 0.2911392, 0.2727273, 0.3000000, 0.2820513, 0.2727273,
		0.2911392, 0.0000000, 0.3684211, 0.3239437, 0.3287671, 0.3506494,
		0.2727273, 0.3684211, 0.0000000, 0.3424658, 0.2794118, 0.3055556,
		0.3000000, 0.3239437, 0.3424658, 0.0000000, 0.3636364, 0.3333333,
		0.2820513, 0.3287671, 0.2794118, 0.3636364, 0.0000000, 0.3421053,
		0.2727273, 0.3506494, 0.3055556, 0.3333333, 0.3421053, 0.0000000}

	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			if !check(out.Get(i, j), known.Get(i, j)) {
				t.Error()
				fmt.Println(out.Get(i, j), known.Get(i, j))
			}

		}
	}
}

// Williams2Bool test against R:simba
func TestWilliams2Bool(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("Williams2Bool_D test against R:simba")
	data = GetBoolData()
	out = Williams2Bool_D(data)
	//known values
	dist := [...]float64{0.0000000, 0.2317429, 0.2156528, 0.2433544, 0.2127872, 0.1941217,
		0.2317429, 0.0000000, 0.2852632, 0.2132797, 0.2286910, 0.2679426,
		0.2156528, 0.2852632, 0.0000000, 0.2477169, 0.1755926, 0.2155712,
		0.2433544, 0.2132797, 0.2477169, 0.0000000, 0.2778537, 0.2436036,
		0.2127872, 0.2286910, 0.1755926, 0.2778537, 0.0000000, 0.2466667,
		0.1941217, 0.2679426, 0.2155712, 0.2436036, 0.2466667, 0.0000000}

	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			if !check(out.Get(i, j), known.Get(i, j)) {
				t.Error()
				fmt.Println(out.Get(i, j), known.Get(i, j))
			}

		}
	}
}

// HarteBool test against R:simba
func TestHarteBool(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("HarteBool_D test against R:simba")
	data = GetBoolData()
	out = HarteBool_D(data)
	//known values
	dist := [...]float64{0.0000000, 0.5192308, 0.4951456, 0.5384615, 0.4857143, 0.4528302,
		0.5192308, 0.0000000, 0.6000000, 0.4791667, 0.5051546, 0.5714286,
		0.4951456, 0.6000000, 0.0000000, 0.5368421, 0.4166667, 0.4845361,
		0.5384615, 0.4791667, 0.5368421, 0.0000000, 0.5876289, 0.5306122,
		0.4857143, 0.5051546, 0.4166667, 0.5876289, 0.0000000, 0.5353535,
		0.4528302, 0.5714286, 0.4845361, 0.5306122, 0.5353535, 0.0000000}

	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			if !check(out.Get(i, j), known.Get(i, j)) {
				t.Error()
				fmt.Println(out.Get(i, j), known.Get(i, j))
			}

		}
	}
}

// Simpson1Bool test against R:simba
func TestSimpson1Bool(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("Simpson1Bool_D test against R:simba")
	data = GetBoolData()
	out = Simpson1Bool_D(data)
	//known values
	dist := [...]float64{0.0000000, 0.4791667, 0.4468085, 0.5000000, 0.4489796, 0.4200000,
		0.4791667, 0.0000000, 0.5957447, 0.4791667, 0.5000000, 0.5625000,
		0.4468085, 0.5957447, 0.0000000, 0.5319149, 0.4042553, 0.4680851,
		0.5000000, 0.4791667, 0.5319149, 0.0000000, 0.5833333, 0.5208333,
		0.4489796, 0.5000000, 0.4042553, 0.5833333, 0.0000000, 0.5306122,
		0.4200000, 0.5625000, 0.4680851, 0.5208333, 0.5306122, 0.0000000}

	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			if !check(out.Get(i, j), known.Get(i, j)) {
				t.Error()
				fmt.Println(out.Get(i, j), known.Get(i, j))
			}

		}
	}
}

// Lennon1Bool test against R:simba
func TestLennon1Bool(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("Lennon1Bool_D test against R:simba")
	data = GetBoolData()
	out = Lennon1Bool_D(data)
	//known values
	dist := [...]float64{0.0000000, 0.15384615, 0.17475728, 0.15384615, 0.13333333, 0.11320755,
		0.1538462, 0.00000000, 0.02105263, 0.00000000, 0.02061856, 0.04081633,
		0.1747573, 0.02105263, 0.00000000, 0.02105263, 0.04166667, 0.06185567,
		0.1538462, 0.00000000, 0.02105263, 0.00000000, 0.02061856, 0.04081633,
		0.1333333, 0.02061856, 0.04166667, 0.02061856, 0.00000000, 0.02020202,
		0.1132075, 0.04081633, 0.06185567, 0.04081633, 0.02020202, 0.00000000}

	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			if !check(out.Get(i, j), known.Get(i, j)) {
				t.Error()
				fmt.Println(out.Get(i, j), known.Get(i, j))
			}

		}
	}
}

// WeiherBool test against R:simba
func TestWeiherBool(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("WeiherBool_D test against R:simba")
	data = GetBoolData()
	out = WeiherBool_D(data)
	//known values
	dist := [...]float64{0, 54, 51, 56, 51, 48,
		54, 0, 57, 46, 49, 56,
		51, 57, 0, 51, 40, 47,
		56, 46, 51, 0, 57, 52,
		51, 49, 40, 57, 0, 53,
		48, 56, 47, 52, 53, 0}

	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			if !check(out.Get(i, j), known.Get(i, j)) {
				t.Error()
				fmt.Println(out.Get(i, j), known.Get(i, j))
			}

		}
	}
}

/* this fails against R:simba
// RuggieroBool test against R:simba
func TestRuggieroBool(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("RuggieroBool_S test against R:simba")
	data = GetBoolData()
	out = RuggieroBool_S(data)
	//known values
	dist := [...]float64{0.0000000,0.4464286,0.4642857,0.4285714,0.4821429,0.5178571,
0.4464286,0.0000000,0.3958333,0.5208333,0.5000000,0.4375000,
0.4642857,0.3958333,0.0000000,0.4680851,0.5957447,0.5319149,
0.4285714,0.5208333,0.4680851,0.0000000,0.4166667,0.4791667,
0.4821429,0.5000000,0.5957447,0.4166667,0.0000000,0.4693878,
0.5178571,0.4375000,0.5319149,0.4791667,0.4693878,0.0000000}

	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i+1; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := i+1; j < rows; j++ {
			if !check(out.Get(i, j), known.Get(i, j)) {
				t.Error()
				fmt.Println(i, j, out.Get(i, j), known.Get(i, j))
			}

		}
	}
}
*/

// Lennon2Bool test against R:simba
func TestLennon2Bool(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("Lennon2Bool_D test against R:simba")
	data = GetBoolData()
	out = Lennon2Bool_D(data)
	//known values
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
		for j := i + 1; j < rows; j++ {
			if !check(out.Get(i, j), known.Get(i, j)) {
				t.Error()
				fmt.Println(out.Get(i, j), known.Get(i, j))
			}

		}
	}
}

// Routledge1Bool test against R:simba
func TestRoutledge1Bool(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("Routledge1Bool_D test against R:simba")
	data = GetBoolData()
	out = Routledge1Bool_D(data)
	//known values
	dist := [...]float64{0.0000000, 0.2961578, 0.2698651, 0.3157895, 0.2653910, 0.2364964,
		0.2961578, 0.0000000, 0.3911368, 0.2656289, 0.2906273, 0.3589273,
		0.2698651, 0.3911368, 0.0000000, 0.3226607, 0.2085729, 0.2693438,
		0.3157895, 0.2656289, 0.3226607, 0.0000000, 0.3772358, 0.3157895,
		0.2653910, 0.2906273, 0.2085729, 0.3772358, 0.0000000, 0.3211345,
		0.2364964, 0.3589273, 0.2693438, 0.3157895, 0.3211345, 0.0000000}

	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			if !check(out.Get(i, j), known.Get(i, j)) {
				t.Error()
				fmt.Println(out.Get(i, j), known.Get(i, j))
			}

		}
	}
}

// Routledge2Bool test against R:simba
func TestRoutledge2Bool(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("Routledge2Bool_D test against R:simba")
	data = GetBoolData()
	out = Routledge2Bool_D(data)
	//known values
	dist := [...]float64{0.0000000, 0.3569418, 0.3393864, 0.3702716, 0.3344476, 0.3122751,
		0.3569418, 0.0000000, 0.4158329, 0.3321330, 0.3500934, 0.3958758,
		0.3393864, 0.4158329, 0.0000000, 0.3720552, 0.2885943, 0.3353765,
		0.3702716, 0.3321330, 0.3720552, 0.0000000, 0.4072602, 0.3675841,
		0.3344476, 0.3500934, 0.2885943, 0.4072602, 0.0000000, 0.3710278,
		0.3122751, 0.3958758, 0.3353765, 0.3675841, 0.3710278, 0.0000000}

	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			if !check(out.Get(i, j), known.Get(i, j)) {
				t.Error()
				fmt.Println(out.Get(i, j), known.Get(i, j))
			}

		}
	}
}

// Routledge3Bool test against R:simba
func TestRoutledge3Bool(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("Routledge3Bool_D test against R:simba")
	data = GetBoolData()
	out = Routledge3Bool_D(data)
	//known values
	dist := [...]float64{0.0000000, 0.4289528, 0.4040858, 0.4481279, 0.3971684, 0.3665306,
		0.4289528, 0.0000000, 0.5156326, 0.3939383, 0.4192001, 0.4856848,
		0.4040858, 0.5156326, 0.0000000, 0.4507130, 0.3345502, 0.3984668,
		0.4481279, 0.3939383, 0.4507130, 0.0000000, 0.5026950, 0.4442413,
		0.3971684, 0.4192001, 0.3345502, 0.5026950, 0.0000000, 0.4492233,
		0.3665306, 0.4856848, 0.3984668, 0.4442413, 0.4492233, 0.0000000}

	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			if !check(out.Get(i, j), known.Get(i, j)) {
				t.Error()
				fmt.Println(out.Get(i, j), known.Get(i, j))
			}

		}
	}
}

// SokalSneath5Bool test against R:simba
func TestSokalSneath5Bool(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("SokalSneath5Bool_S test against R:simba")
	data = GetBoolData()
	out = SokalSneath5Bool_S(data)
	//known values
	dist := [...]float64{0.0000000, 0.1879699, 0.2031250, 0.1764706, 0.2093023, 0.2320000,
		0.1879699, 0.0000000, 0.1428571, 0.2136752, 0.1967213, 0.1578947,
		0.2031250, 0.1428571, 0.0000000, 0.1774194, 0.2592593, 0.2100840,
		0.1764706, 0.2136752, 0.1774194, 0.0000000, 0.1492537, 0.1811024,
		0.2093023, 0.1967213, 0.2592593, 0.1492537, 0.0000000, 0.1782946,
		0.2320000, 0.1578947, 0.2100840, 0.1811024, 0.1782946, 0.0000000}

	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			if !check(out.Get(i, j), known.Get(i, j)) {
				t.Error()
				fmt.Println(out.Get(i, j), known.Get(i, j))
			}

		}
	}
}

// DiceBool test against R:simba
func TestDiceBool(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("DiceBool_S test against R:simba")
	data = GetBoolData()
	out = DiceBool_S(data)
	//known values
	dist := [...]float64{0.0000000, 0.5208333, 0.5531915, 0.5000000, 0.5510204, 0.5800000,
		0.5208333, 0.0000000, 0.4042553, 0.5208333, 0.5000000, 0.4375000,
		0.5531915, 0.4042553, 0.0000000, 0.4680851, 0.5957447, 0.5319149,
		0.5000000, 0.5208333, 0.4680851, 0.0000000, 0.4166667, 0.4791667,
		0.5510204, 0.5000000, 0.5957447, 0.4166667, 0.0000000, 0.4693878,
		0.5800000, 0.4375000, 0.5319149, 0.4791667, 0.4693878, 0.0000000}

	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			if !check(out.Get(i, j), known.Get(i, j)) {
				t.Error()
				fmt.Println(out.Get(i, j), known.Get(i, j))
			}

		}
	}
}

// Kulczynski1Bool test against R:simba
func TestKulczynski11Bool(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("Kulczynski1Bool_S test against R:simba")
	data = GetBoolData()
	out = Kulczynski1Bool_S(data)
	//known values
	dist := [...]float64{0.0000000, 0.4629630, 0.5098039, 0.4285714, 0.5294118, 0.6041667,
		0.4629630, 0.0000000, 0.3333333, 0.5434783, 0.4897959, 0.3750000,
		0.5098039, 0.3333333, 0.0000000, 0.4313725, 0.7000000, 0.5319149,
		0.4285714, 0.5434783, 0.4313725, 0.0000000, 0.3508772, 0.4423077,
		0.5294118, 0.4897959, 0.7000000, 0.3508772, 0.0000000, 0.4339623,
		0.6041667, 0.3750000, 0.5319149, 0.4423077, 0.4339623, 0.0000000}

	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			if !check(out.Get(i, j), known.Get(i, j)) {
				t.Error()
				fmt.Println(out.Get(i, j), known.Get(i, j))
			}

		}
	}
}

// Kulczynski2Bool test against R:simba
func TestKulczynski21Bool(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("Kulczynski2Bool_S test against R:simba")
	data = GetBoolData()
	out = Kulczynski2Bool_S(data)
	//known values
	dist := [...]float64{0.0000000, 0.4836310, 0.5087386, 0.4642857, 0.5165816, 0.5489286,
		0.4836310, 0.0000000, 0.4000443, 0.5208333, 0.4948980, 0.4287500,
		0.5087386, 0.4000443, 0.0000000, 0.4632092, 0.5835866, 0.5159574,
		0.4642857, 0.5208333, 0.4632092, 0.0000000, 0.4124150, 0.4695833,
		0.5165816, 0.4948980, 0.5835866, 0.4124150, 0.0000000, 0.4646939,
		0.5489286, 0.4287500, 0.5159574, 0.4695833, 0.4646939, 0.0000000}

	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			if !check(out.Get(i, j), known.Get(i, j)) {
				t.Error()
				fmt.Println(out.Get(i, j), known.Get(i, j))
			}

		}
	}
}

// McConnaghBool test against R:simba
func TestMcConnagh1Bool(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("McConnaghBool_S test against R:simba")
	data = GetBoolData()
	out = McConnaghBool_S(data)
	//known values
	dist := [...]float64{0.00000000, -0.03273810, 0.01747720, -0.07142857, 0.03316327, 0.09785714,
		-0.03273810, 0.00000000, -0.19991135, 0.04166667, -0.01020408, -0.14250000,
		0.01747720, -0.19991135, 0.00000000, -0.07358156, 0.16717325, 0.03191489,
		-0.07142857, 0.04166667, -0.07358156, 0.00000000, -0.17517007, -0.06083333,
		0.03316327, -0.01020408, 0.16717325, -0.17517007, 0.00000000, -0.07061224,
		0.09785714, -0.14250000, 0.03191489, -0.06083333, -0.07061224, 0.00000000}

	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			if !check(out.Get(i, j), known.Get(i, j)) {
				t.Error()
				fmt.Println(out.Get(i, j), known.Get(i, j))
			}

		}
	}
}

// ManhattanBool test against R:simba
func TestManhattanBool(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("ManhattanBool_D test against R:simba")
	data = GetBoolData()
	out = ManhattanBool_D(data)
	//known values
	dist := [...]float64{0.00, 0.54, 0.51, 0.56, 0.51, 0.48,
		0.54, 0.00, 0.57, 0.46, 0.49, 0.56,
		0.51, 0.57, 0.00, 0.51, 0.40, 0.47,
		0.56, 0.46, 0.51, 0.00, 0.57, 0.52,
		0.51, 0.49, 0.40, 0.57, 0.00, 0.53,
		0.48, 0.56, 0.47, 0.52, 0.53, 0.00}

	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			if !check(out.Get(i, j), known.Get(i, j)) {
				t.Error()
				fmt.Println(out.Get(i, j), known.Get(i, j))
			}

		}
	}
}

// BraunBlanquetBool test against R:fossil
func TestBraunBlanquetBool(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("BraunBlanquetBool test against R:fossil")
	data = GetBoolData()
	out = BraunBlanquetBool_S(data)
	//known values
	dist := [...]float64{0.0000000, 0.4464286, 0.4642857, 0.4285714, 0.4821429, 0.5178571,
		0.4464286, 0.0000000, 0.3958333, 0.5208333, 0.4897959, 0.4200000,
		0.4642857, 0.3958333, 0.0000000, 0.4583333, 0.5714286, 0.5000000,
		0.4285714, 0.5208333, 0.4583333, 0.0000000, 0.4081633, 0.4600000,
		0.4821429, 0.4897959, 0.5714286, 0.4081633, 0.0000000, 0.4600000,
		0.5178571, 0.4200000, 0.5000000, 0.4600000, 0.4600000, 0.0000000}

	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			if !check(out.Get(i, j), known.Get(i, j)) {
				t.Error()
				fmt.Println(out.Get(i, j), known.Get(i, j))
			}

		}
	}
}

// BraunBlanquetBool test 2 against R:fossil
func TestBraunBlanquetBool2(t *testing.T) {
	var (
		data, out, known *Matrix
	)

	fmt.Println("BraunBlanquetBool test 2 against R:fossil")
	data = GetBoolData2()
	out = BraunBlanquetBool_S(data)
	//known values
	dist := [...]float64{0.0000000, 0.5000000, 0.3333333, 0.8333333, 0.3333333, 0.3333333,
		0.5000000, 0.0000000, 0.2500000, 0.6666667, 0.3333333, 0.2000000,
		0.3333333, 0.2500000, 0.0000000, 0.3333333, 0.3333333, 0.0000000,
		0.8333333, 0.6666667, 0.3333333, 0.0000000, 0.3333333, 0.3333333,
		0.3333333, 0.3333333, 0.3333333, 0.3333333, 0.0000000, 0.5000000,
		0.3333333, 0.2000000, 0.0000000, 0.3333333, 0.5000000, 0.0000000}

	rows := data.R
	known = NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			known.Set(i, j, dist[i*rows+j])
		}
	}

	// check
	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			if !check(out.Get(i, j), known.Get(i, j)) {
				t.Error()
				fmt.Println(i, j, out.Get(i, j), known.Get(i, j))
			}

		}
	}
}
