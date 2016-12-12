package sort

import (
	"reflect"
)

//Tis is a custom Merge Sort.
//You can define the comparion regulations all by yourself.
//
//You should write your own compare function with definition
//" func(i,j interface{})bool "
//then MergeSort() will launch Merge Sort with the regulation you define.
//(*you can use i.(YourInputType) to convert i to original type)
//
//The output will be a slice of interface{},you should use o.(YourInputType) to convert o to origibal type.
func MergeSort(data interface{}, cmp func(i, j interface{}) bool) []interface{} {
	//data : interface{} to []interface
	value := reflect.ValueOf(data)
	N := make([]interface{}, value.Len())
	for a := 0; a < value.Len(); a++ {
		N[a] = value.Index(a).Interface()
	}
	msHandle(N, cmp)
	return N
}

func msHandle(data []interface{}, cmp func(i, j interface{}) bool) {
	M := len(data) / 2
	if len(data) == 1 {
		return
	}
	msHandle(data[:M], cmp)
	msHandle(data[M:], cmp)
	merge(data, M, cmp)
}

func merge(data []interface{}, M int, cmp func(i, j interface{}) bool) {
	tmp := make([]interface{}, len(data))
	length := len(data)
	var a, b, p int = 0, 0, 0
	for a, b = 0, M; a < M && b < length; {
		if cmp(data[a], data[b]) {
			tmp[p] = data[a]
			p++
			a++
		} else {
			tmp[p] = data[b]
			p++
			b++
		}
	}
	if a == M {
		for i := b; i < length; i, p = i+1, p+1 {
			tmp[p] = data[i]
		}
	} else {
		for i := a; i < M; i, p = i+1, p+1 {
			tmp[p] = data[i]
		}
	}
	for i := 0; i < length; i++ {
		data[i] = tmp[i]
	}
}

//MergeSort for int, sorting from smallest to biggest.
func MergeSortInt(data []int) {
	M := len(data) / 2
	if len(data) == 1 {
		return
	}
	MergeSortInt(data[:M])
	MergeSortInt(data[M:])
	mergeInt(data, M)
}

func mergeInt(data []int, M int) {
	tmp := make([]int, len(data))
	length := len(data)
	var a, b, p int = 0, 0, 0
	for a, b = 0, M; a < M && b < length; {
		if data[a] < data[b] {
			tmp[p] = data[a]
			p++
			a++
		} else {
			tmp[p] = data[b]
			p++
			b++
		}
	}
	if a == M {
		for i := b; i < length; i, p = i+1, p+1 {
			tmp[p] = data[i]
		}
	} else {
		for i := a; i < M; i, p = i+1, p+1 {
			tmp[p] = data[i]
		}
	}
	for i := 0; i < length; i++ {
		data[i] = tmp[i]
	}
}

//MergeSort for int32, sorting from smallest to biggest.
func MergeSortInt32(data []int32) {
	M := len(data) / 2
	if len(data) == 1 {
		return
	}
	MergeSortInt32(data[:M])
	MergeSortInt32(data[M:])
	mergeInt32(data, M)
}

func mergeInt32(data []int32, M int) {
	tmp := make([]int32, len(data))
	length := len(data)
	var a, b, p int = 0, 0, 0
	for a, b = 0, M; a < M && b < length; {
		if data[a] < data[b] {
			tmp[p] = data[a]
			p++
			a++
		} else {
			tmp[p] = data[b]
			p++
			b++
		}
	}
	if a == M {
		for i := b; i < length; i, p = i+1, p+1 {
			tmp[p] = data[i]
		}
	} else {
		for i := a; i < M; i, p = i+1, p+1 {
			tmp[p] = data[i]
		}
	}
	for i := 0; i < length; i++ {
		data[i] = tmp[i]
	}
}

//MergeSort for int64, sorting from smallest to biggest.
func MergeSortInt64(data []int64) {
	M := len(data) / 2
	if len(data) == 1 {
		return
	}
	MergeSortInt64(data[:M])
	MergeSortInt64(data[M:])
	mergeInt64(data, M)
}

func mergeInt64(data []int64, M int) {
	tmp := make([]int64, len(data))
	length := len(data)
	var a, b, p int = 0, 0, 0
	for a, b = 0, M; a < M && b < length; {
		if data[a] < data[b] {
			tmp[p] = data[a]
			p++
			a++
		} else {
			tmp[p] = data[b]
			p++
			b++
		}
	}
	if a == M {
		for i := b; i < length; i, p = i+1, p+1 {
			tmp[p] = data[i]
		}
	} else {
		for i := a; i < M; i, p = i+1, p+1 {
			tmp[p] = data[i]
		}
	}
	for i := 0; i < length; i++ {
		data[i] = tmp[i]
	}
}

//MergeSort for float32, sorting from smallest to biggest.
func MergeSortFloat32(data []float32) {
	M := len(data) / 2
	if len(data) == 1 {
		return
	}
	MergeSortFloat32(data[:M])
	MergeSortFloat32(data[M:])
	mergeFloat32(data, M)
}

func mergeFloat32(data []float32, M int) {
	tmp := make([]float32, len(data))
	length := len(data)
	var a, b, p int = 0, 0, 0
	for a, b = 0, M; a < M && b < length; {
		if data[a] < data[b] {
			tmp[p] = data[a]
			p++
			a++
		} else {
			tmp[p] = data[b]
			p++
			b++
		}
	}
	if a == M {
		for i := b; i < length; i, p = i+1, p+1 {
			tmp[p] = data[i]
		}
	} else {
		for i := a; i < M; i, p = i+1, p+1 {
			tmp[p] = data[i]
		}
	}
	for i := 0; i < length; i++ {
		data[i] = tmp[i]
	}
}

//MergeSort for float64, sorting from smallest to biggest.
func MergeSortFloat64(data []float64) {
	M := len(data) / 2
	if len(data) == 1 {
		return
	}
	MergeSortFloat64(data[:M])
	MergeSortFloat64(data[M:])
	mergeFloat64(data, M)
}

func mergeFloat64(data []float64, M int) {
	tmp := make([]float64, len(data))
	length := len(data)
	var a, b, p int = 0, 0, 0
	for a, b = 0, M; a < M && b < length; {
		if data[a] < data[b] {
			tmp[p] = data[a]
			p++
			a++
		} else {
			tmp[p] = data[b]
			p++
			b++
		}
	}
	if a == M {
		for i := b; i < length; i, p = i+1, p+1 {
			tmp[p] = data[i]
		}
	} else {
		for i := a; i < M; i, p = i+1, p+1 {
			tmp[p] = data[i]
		}
	}
	for i := 0; i < length; i++ {
		data[i] = tmp[i]
	}
}
