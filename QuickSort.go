package sort

import (
	"reflect"
)

//QuickSort() is a custom quick sort.
//Sou should build your own compare function to sort in your own regulations.
func QuickSort(data interface{}, cmp func(i, j interface{}) bool) []interface{} {
	value := reflect.ValueOf(data)
	dataS := make([]interface{}, value.Len())
	for a := 0; a < value.Len(); a++ {
		dataS[a] = value.Index(a).Interface()
	}
	qsHandle(dataS, 0, len(dataS)-1, cmp)
	return dataS
}

func qsHandle(data []interface{}, left, right int, cmp func(i, j interface{}) bool) {
	if left >= right {
		return
	}

	pivot := data[left]
	i, j := left+1, right

	for true {
		for i <= right {
			if cmp(pivot, data[i]) {
				break
			}
			i++
		}

		for j > left {
			if cmp(data[j], pivot) {
				break
			}
			j = j - 1
		}

		if i > j {
			break
		} else {
			data[i], data[j] = data[j], data[i]
		}
	}

	data[left], data[j] = data[j], data[left]

	qsHandle(data, left, j-1, cmp)
	qsHandle(data, j+1, right, cmp)
}
