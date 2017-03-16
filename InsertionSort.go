package sort

import "reflect"

// This is a custom Merge Sort
// You MUST define the compare function by yourself
// compare function type: func(i,j interface{}) bool
//
// input: func([]interface{}, func(i,j interface{}) bool) []interface{}
//
// output slice and won't change the origin one
// data = InsertionSort(data, cmp), if you want to change origin slice
func InsertionSort(data interface{}, cmp func(i, j interface{}) bool) []interface{} {
	value := reflect.ValueOf(data)
	dataS := make([]interface{}, value.Len())
	len := value.Len()
	for a := 0; a < len; a++ {
		dataS[a] = value.Index(a).Interface()
	}
	for i := 1; i < len; i++ {
		tmp := dataS[i]
		j := i - 1
		for ; j >= 0 && cmp(dataS[j], tmp); j-- {
			dataS[j+1] = dataS[j]
		}
		dataS[j+1] = tmp
	}
	return dataS
}

// INT
// Insertion Sort in Int type, sorting from small to big
//
// Input slice func([]int)
//
// No output, just change origin slice
func InsertionSortInt(dataS []int) {
	len := len(dataS)
	for i := 1; i < len; i++ {
		tmp := dataS[i]
		j := i - 1
		for ; j >= 0 && dataS[j] > tmp; j-- {
			dataS[j+1] = dataS[j]
		}
		dataS[j+1] = tmp
	}
}

// INT32
// Insertion Sort in Int32 type, sorting from small to big
//
// Input slice func([]int32)
//
// No output, just change origin slice
func InsertionSortInt32(dataS []int32) {
	len := len(dataS)
	for i := 1; i < len; i++ {
		tmp := dataS[i]
		j := i - 1
		for ; j >= 0 && dataS[j] > tmp; j-- {
			dataS[j+1] = dataS[j]
		}
		dataS[j+1] = tmp
	}
}

// INT64
// Insertion Sort in Int64 type, sorting from small to big
//
// Input slice func([]int64)
//
// No output, just change origin slice
func InsertionSortInt64(dataS []int64) {
	len := len(dataS)
	for i := 1; i < len; i++ {
		tmp := dataS[i]
		j := i - 1
		for ; j >= 0 && dataS[j] > tmp; j-- {
			dataS[j+1] = dataS[j]
		}
		dataS[j+1] = tmp
	}
}

// FLOAT32
// Insertion Sort in Float32 type, sorting from small to big
//
// Input slice func([]float32)
//
// No output, just change origin slice
func InsertionSortFloat32(dataS []float32) {
	len := len(dataS)
	for i := 1; i < len; i++ {
		tmp := dataS[i]
		j := i - 1
		for ; j >= 0 && dataS[j] > tmp; j-- {
			dataS[j+1] = dataS[j]
		}
		dataS[j+1] = tmp
	}
}

// FLOAT64
// Insertion Sort in Float64 type, sorting from small to big
//
// Input slice func([]float64)
//
// No output, just change origin slice
func InsertionSortFloat64(dataS []float64) {
	len := len(dataS)
	for i := 1; i < len; i++ {
		tmp := dataS[i]
		j := i - 1
		for ; j >= 0 && dataS[j] > tmp; j-- {
			dataS[j+1] = dataS[j]
		}
		dataS[j+1] = tmp
	}
}
