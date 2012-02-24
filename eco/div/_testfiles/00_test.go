package div

import (
	. "go-eco.googlecode.com/hg/eco"

)

func check(x, y float64) bool {
	const acc float64 = 1e-6 // accuracy
	var z float64
	if x/y > 1.00 {
		z = y / x
	} else {
		z = x / y
	}
	if 1-z > acc {
		return false
	}
	return true
}

// Get boolean data matrix
func GetBoolData() *Matrix {
	var (
		data *Matrix
	)
	rows := 6
	cols := 100
	arr := [...]float64{1, 0, 0, 1, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 0, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 1, 0, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 1, 1, 0, 1, 0, 1, 1,
		1, 1, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 1, 0, 1, 1, 0, 1, 0, 1, 0, 0, 1, 0, 0, 1, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 1, 1, 0, 0, 1, 0, 1, 0, 1, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 1, 1, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 0, 0, 1, 0, 1, 1, 0, 1, 0, 0, 0, 1, 1, 0, 0, 1, 0, 1, 1, 0,
		1, 0, 1, 0, 1, 1, 1, 0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 1, 0, 1, 1, 1, 0, 1, 0, 0, 1, 0, 0, 0, 1, 0, 1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 0, 1, 0, 1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 0, 0, 1, 0, 0, 0, 1, 1, 1, 1, 0, 0,
		0, 0, 0, 1, 0, 0, 1, 1, 1, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0, 1, 0, 1, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 1, 0, 0, 1, 0, 0, 1, 1, 0, 1, 0, 1, 1, 1, 0, 1, 0, 0, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 1, 1, 1,
		1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0, 1, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 1, 1, 0, 1, 0, 0, 1, 0, 1, 1, 0, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 0, 1, 1, 1, 0, 1, 1, 0, 1, 0, 0, 0, 1, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 1, 0, 0}

	data = NewMatrix(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			data.Set(i, j, arr[i*cols+j])
		}
	}
	return data
}

// Get smaller boolean data matrix
func GetBoolData2() *Matrix {
	var (
		data *Matrix
	)
	rows := 6
	cols := 10
	arr := [...]float64{0, 1, 0, 1, 1, 0, 1, 0, 1, 1,
		0, 1, 1, 0, 1, 0, 0, 0, 1, 0,
		0, 0, 1, 1, 0, 0, 1, 0, 0, 0,
		0, 1, 1, 0, 1, 0, 1, 0, 1, 1,
		1, 0, 1, 1, 1, 1, 0, 1, 0, 0,
		1, 0, 0, 0, 0, 1, 0, 1, 1, 1}

	data = NewMatrix(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			data.Set(i, j, arr[i*cols+j])
		}
	}
	return data
}

// Get float data matrix
func GetData() *Matrix {
	var (
		data *Matrix
	)
	rows := 6
	cols := 10
	arr := [...]float64{4.637511, 5.795001, 5.700484, 6.524882, 6.170708, 6.690082, 6.156994, 6.921186, 4.740336, 5.400573,
		7.535951, 5.804745, 6.697524, 5.297671, 5.77213, 7.187614, 7.470511, 6.559553, 5.870524, 5.31025,
		4.428564, 5.698517, 4.882601, 6.541425, 4.910434, 7.311253, 4.562559, 6.858137, 7.105823, 5.963177,
		6.614128, 4.282033, 6.146613, 5.819527, 6.797518, 5.657332, 5.61468, 5.180996, 5.374655, 6.594351,
		6.79283, 6.371214, 5.990534, 6.518546, 6.283301, 6.841622, 5.978732, 5.278547, 7.825815, 6.36177,
		4.813066, 6.990308, 6.809527, 7.83582, 6.256215, 4.981545, 7.230944, 5.322504, 5.981109, 5.738691}

	data = NewMatrix(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			data.Set(i, j, arr[i*cols+j])
		}
	}
	return data
}

// Get categorical data matrix
func GetCatData() *Matrix {
	var (
		data *Matrix
	)
	rows := 6
	cols := 100
	arr := [...]float64{44, 27, 12, 24, 30, 2, 12, 47, 21, 19, 11, 32, 19, 46, 38, 47, 10, 39, 15, 6, 21, 14, 48, 11, 25, 27, 17, 12, 28, 18, 17, 7, 26, 39, 13, 13, 43, 39, 35, 16, 7, 22, 24, 14, 28, 9, 37, 19, 20, 8, 8, 23, 10, 25, 43, 46, 29, 21, 40, 38, 6, 14, 14, 44, 18, 10, 37, 25, 33, 25, 33, 1, 39, 25, 38, 34, 22, 42, 23, 9, 12, 11, 10, 37, 23, 27, 50, 16, 33, 45, 3, 33, 33, 9, 12, 6, 48, 47, 17, 45,
		48, 2, 36, 18, 36, 45, 23, 44, 47, 15, 39, 34, 25, 10, 28, 2, 40, 20, 28, 21, 25, 3, 39, 8, 24, 7, 37, 2, 45, 18, 8, 49, 11, 29, 24, 36, 18, 29, 48, 44, 32, 16, 19, 34, 46, 37, 19, 8, 1, 30, 50, 15, 15, 34, 33, 6, 14, 4, 1, 4, 45, 49, 3, 15, 45, 18, 44, 10, 2, 5, 37, 14, 18, 41, 15, 15, 12, 47, 12, 37, 23, 43, 2, 2, 9, 16, 13, 48, 12, 45, 6, 27, 46, 7, 46, 30, 12, 17, 32, 22,
		7, 1, 41, 40, 28, 16, 28, 16, 35, 35, 32, 27, 7, 40, 12, 36, 2, 3, 38, 29, 41, 49, 46, 29, 4, 15, 24, 47, 30, 15, 7, 28, 30, 4, 47, 35, 15, 16, 2, 7, 9, 46, 23, 11, 15, 48, 38, 18, 4, 38, 6, 37, 4, 41, 16, 39, 20, 24, 35, 36, 20, 29, 8, 17, 48, 3, 7, 38, 24, 47, 39, 22, 22, 5, 2, 22, 13, 15, 47, 29, 2, 16, 19, 14, 39, 2, 27, 15, 2, 24, 11, 32, 0, 9, 46, 2, 37, 15, 10, 7,
		14, 36, 38, 37, 30, 30, 37, 49, 34, 26, 12, 46, 21, 10, 32, 50, 41, 32, 24, 39, 9, 5, 40, 16, 44, 31, 14, 7, 26, 7, 25, 12, 40, 45, 14, 14, 41, 17, 40, 43, 45, 28, 21, 25, 36, 6, 23, 4, 23, 20, 35, 2, 3, 29, 3, 38, 49, 17, 48, 1, 9, 13, 24, 17, 42, 15, 28, 4, 32, 50, 17, 49, 28, 22, 10, 27, 47, 10, 46, 14, 48, 13, 50, 1, 6, 17, 22, 41, 42, 42, 26, 5, 32, 6, 47, 17, 9, 50, 46, 27,
		5, 40, 7, 48, 14, 47, 18, 45, 21, 12, 40, 18, 35, 12, 30, 48, 11, 50, 11, 12, 18, 49, 18, 19, 27, 32, 36, 29, 27, 13, 19, 31, 46, 13, 11, 3, 28, 27, 17, 46, 20, 24, 6, 11, 40, 45, 5, 1, 48, 24, 14, 32, 47, 22, 2, 22, 45, 23, 24, 36, 39, 28, 34, 30, 16, 22, 5, 31, 39, 13, 39, 36, 32, 22, 24, 27, 43, 41, 42, 35, 5, 39, 50, 36, 12, 1, 28, 35, 13, 7, 44, 20, 26, 18, 39, 8, 10, 23, 9, 32,
		17, 19, 20, 30, 48, 35, 38, 23, 35, 29, 45, 42, 37, 24, 27, 8, 37, 18, 32, 7, 50, 49, 43, 9, 37, 12, 3, 37, 30, 48, 1, 46, 44, 22, 29, 45, 15, 31, 40, 33, 8, 38, 39, 6, 28, 40, 11, 3, 9, 19, 17, 9, 22, 42, 41, 45, 16, 40, 44, 2, 49, 36, 44, 4, 36, 34, 8, 2, 15, 3, 18, 4, 5, 10, 13, 34, 36, 33, 35, 12, 9, 9, 26, 33, 11, 24, 10, 42, 40, 30, 22, 37, 5, 20, 4, 22, 6, 47, 7, 35}

	data = NewMatrix(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			data.Set(i, j, arr[i*cols+j])
		}
	}
	return data
}

