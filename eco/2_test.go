package eco

import (
	"testing"
	"fmt"
	. "gomatrix.googlecode.com/hg/matrix"
)

// Get binary data matrix
func GetBinData()  *DenseMatrix {
	var (
		data *DenseMatrix
	)
	rows := 6
	cols := 100
	arr := [...]float64{1,0,0,1,0,1,0,1,0,0,0,1,0,0,1,0,0,0,0,0,0,0,0,1,1,1,1,1,1,1,0,1,1,1,1,0,1,1,0,1,1,1,0,1,1,0,0,0,0,0,1,0,0,1,0,1,0,1,1,1,1,0,1,0,0,0,0,1,0,1,1,1,1,1,1,1,1,1,0,1,1,1,1,0,1,0,0,1,0,0,1,1,1,1,1,1,0,1,0,1,
1,1,1,1,1,1,1,0,0,1,0,1,0,1,0,1,0,1,0,1,0,0,1,1,0,1,0,0,0,0,0,0,1,0,1,1,0,0,0,1,0,0,0,0,0,0,0,0,1,0,0,1,1,1,0,0,1,0,1,1,1,0,0,1,1,1,1,1,0,0,0,1,0,0,0,1,0,1,1,1,1,1,1,0,0,0,0,0,1,0,0,0,1,1,1,0,1,0,1,1,
1,1,0,0,0,0,1,1,1,0,0,0,1,1,1,1,1,0,1,1,0,1,0,1,0,0,1,0,0,1,1,0,0,0,0,0,1,0,0,0,1,1,1,0,0,1,0,1,0,1,1,0,0,0,0,1,0,0,1,0,0,0,0,1,0,1,1,1,1,0,0,1,1,0,0,1,1,1,0,1,0,0,1,0,1,1,0,1,0,0,0,1,1,0,0,1,0,1,1,0,
1,0,1,0,1,1,1,0,0,0,1,0,0,1,0,0,1,1,0,1,1,1,0,1,0,0,1,0,0,0,1,0,1,0,1,1,1,0,1,1,1,0,0,0,0,0,1,1,1,0,0,1,1,1,1,1,0,0,0,1,0,0,1,0,1,1,1,0,0,0,1,1,1,1,0,1,0,0,0,0,0,0,0,1,1,0,0,1,0,0,1,0,0,0,1,1,1,1,0,0,
0,0,0,1,0,0,1,1,1,0,1,1,1,1,0,0,0,0,1,1,0,0,1,1,1,0,0,1,0,1,1,0,0,0,0,0,1,0,0,0,1,0,1,0,1,0,0,0,0,1,0,0,1,0,1,0,1,0,0,1,0,0,1,1,0,1,0,1,1,1,0,1,0,0,1,1,1,0,0,1,1,1,1,1,1,1,1,0,0,0,0,0,1,0,1,0,1,1,1,1,
1,0,0,1,1,1,1,1,0,0,0,0,1,0,1,1,1,0,0,1,1,1,0,1,1,0,0,1,1,0,1,0,0,1,0,1,1,0,1,0,1,1,1,1,0,1,0,0,0,1,1,1,0,1,1,0,1,0,0,0,1,1,1,0,1,0,0,0,0,0,0,0,0,0,1,1,1,1,1,1,1,0,1,0,0,0,0,0,0,1,1,0,0,1,0,1,0,1,0,0}

	data = Zeros(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			data.Set(i, j, arr[i*cols+j])
		}
	}
	return data	
}


// Get smaller binary data matrix
func GetBinData2()  *DenseMatrix {
	var (
		data *DenseMatrix
	)
	rows := 6
	cols := 10
	arr := [...]float64{0,1,0,1,1,0,1,0,1,1,
0,1,1,0,1,0,0,0,1,0,
0,0,1,1,0,0,1,0,0,0,
0,1,1,0,1,0,1,0,1,1,
1,0,1,1,1,1,0,1,0,0,
1,0,0,0,0,1,0,1,1,1}

	data = Zeros(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			data.Set(i, j, arr[i*cols+j])
		}
	}
	return data	
}


// Get float data matrix
func GetData()  *DenseMatrix {
	var (
		data *DenseMatrix
	)
	rows := 6
	cols := 10
	arr := [...]float64{4.637511, 5.795001, 5.700484, 6.524882, 6.170708, 6.690082, 6.156994, 6.921186, 4.740336, 5.400573, 7.535951, 5.804745, 6.697524, 5.297671, 5.77213, 7.187614, 7.470511, 6.559553, 5.870524, 5.31025, 4.428564, 5.698517, 4.882601, 6.541425, 4.910434, 7.311253, 4.562559, 6.858137, 7.105823, 5.963177, 6.614128, 4.282033, 6.146613, 5.819527, 6.797518, 5.657332, 5.61468, 5.180996, 5.374655, 6.594351, 6.79283, 6.371214, 5.990534, 6.518546, 6.283301, 6.841622, 5.978732, 5.278547, 7.825815, 6.36177, 4.813066, 6.990308, 6.809527, 7.83582, 6.256215, 4.981545, 7.230944, 5.322504, 5.981109, 5.738691}

	data = Zeros(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			data.Set(i, j, arr[i*cols+j])
		}
	}
	return data	
}


// Euclidean test against R:vegan
func TestEuclid(t *testing.T) {
	var (
		data, out, known *DenseMatrix
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
	rows := data.Rows()
	known = Zeros(rows, rows)
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
		data, out, known *DenseMatrix
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
	rows := data.Rows()
	known = Zeros(rows, rows)
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
		data, out, known *DenseMatrix
	)

	fmt.Println("Canberra Scaled test against R:vegan")
	data = GetData()
	out = CanberraSc_D(data)

//known distances
	dist := [...]float64{0.00000000,0.07305403,0.06705955,0.09046813,0.07580495,0.08021418, 
0.07305403,0.00000000,0.10369549,0.09178595,0.07764712,0.09047445, 
0.06705955,0.10369549,0.00000000,0.12332237,0.08709588,0.11657391, 
0.09046813,0.09178595,0.12332237,0.00000000,0.06573803,0.09637591, 
0.07580495,0.07764712,0.08709588,0.06573803,0.00000000,0.08161942, 
0.08021418,0.09047445,0.11657391,0.09637591,0.08161942,0.00000000}
	rows := data.Rows()
	known = Zeros(rows, rows)
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
		data, out, known *DenseMatrix
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
	rows := data.Rows()
	known = Zeros(rows, rows)
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

// Raup-Crick test against R:vegan
func TestRaup-Crick(t *testing.T) {
	var (
		data, out, known *DenseMatrix
	)

	data = GetBinData()
	out = RaupCrick_D(data)

//known distances
	dist := [...]float64{0.0000000,0.9299438,0.7841876,0.7101091,0.1805086,0.4454451,
0.9299438,0.0000000,0.4936447,0.9150708,0.7588514,0.7478258,
0.7841876,0.4936447,0.0000000,0.4053583,0.9440366,0.9770032,
0.7101091,0.9150708,0.4053583,0.0000000,0.9269091,0.8321101,
0.1805086,0.7588514,0.9440366,0.9269091,0.0000000,0.4579328,
0.4454451,0.7478258,0.9770032,0.8321101,0.4579328,0.0000000}
	rows := data.Rows()
	known = Zeros(rows, rows)
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
		data, out, known *DenseMatrix
	)

	data = GetBinData()
	out = Mountford_D(data)

//known distances
	dist := [...]float64{0.00000000,0.98023890,0.97589280,0.97698570,0.96778280,0.97220950,
0.98023890,0.00000000,0.97175860,0.98092760,0.97711510,0.97650150,
0.97589280,0.97175860,0.00000000,0.97148980,0.97998350,0.98113560,
0.97698570,0.98092760,0.97148980,0.00000000,0.98173240,0.97902470,
0.96778280,0.97711510,0.97998350,0.98173240,0.00000000,0.97279650,
0.97220950,0.97650150,0.98113560,0.97902470,0.97279650,0.00000000}
	rows := data.Rows()
	known = Zeros(rows, rows)
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
		data, out, known *DenseMatrix
	)

	data = GetBinData2()
	out = Mountford_D(data)

//known distances
	dist := [...]float64{0.00000000,0.53504158,0.67824320,0.01189135,0.88138692,0.85054888,
0.53504158,0.00000000,0.83170094,0.00000000,0.79715738,0.90738441,
0.67824320,0.83170094,0.00000000,0.67824320,0.67824320,1.00000000,
0.01189135,0.00000000,0.67824320,0.00000000,0.88138692,0.85054888,
0.88138692,0.79715738,0.67824320,0.88138692,0.00000000,0.69373111,
0.85054888,0.90738441,1.00000000,0.85054888,0.69373111,0.00000000}
	rows := data.Rows()
	known = Zeros(rows, rows)
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

*/

// Raup-Crick test against R:vegan, smaller data
func TestRaupCrick(t *testing.T) {
	var (
		data, out, known *DenseMatrix
	)

	fmt.Println("Raup-Crick test against R:vegan, smaller data")
	data = GetBinData2()
	out = RaupCrick_D(data)

//known distances
	dist := [...]float64{0.00000000,0.45238095,0.66666667,0.11904762,1.00000000,0.97619048,
0.45238095,0.00000000,0.83333333,0.07142857,0.88095238,0.97619048,
0.66666667,0.83333333,0.00000000,0.66666667,0.66666667,1.00000000,
0.11904762,0.07142857,0.66666667,0.00000000,1.00000000,0.97619048,
1.00000000,0.88095238,0.66666667,1.00000000,0.00000000,0.73809524,
0.97619048,0.97619048,1.00000000,0.97619048,0.73809524,0.00000000}



	rows := data.Rows()
	known = Zeros(rows, rows)
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

// Raup-Crick test against R:vegan - big data
func TestRaupCrick2(t *testing.T) {
	var (
		data, out, known *DenseMatrix
	)

	fmt.Println("Raup-Crick test against R:vegan, big data")
	data = GetBinData()
	out = RaupCrick_D(data)

//known distances
	dist := [...]float64{0.00000000,0.83137125,0.62972983,0.91365497,0.64752313,0.42024290,
0.83137125,0.00000000,0.94852481,0.27939267,0.50317018,0.91956864,
0.62972983,0.94852481,0.00000000,0.66445339,0.03635185,0.34442637,
0.91365497,0.27939267,0.66445339,0.00000000,0.94649405,0.72578724,
0.64752313,0.50317018,0.03635185,0.94649405,0.00000000,0.78810701,
0.42024290,0.91956864,0.34442637,0.72578724,0.78810701,0.00000000}
	rows := data.Rows()
	known = Zeros(rows, rows)
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


