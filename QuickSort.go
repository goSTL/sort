package sort

import (
	"math/rand"
	"reflect"
	"time"
)

//QuickSort is a custom quick sort.
//You should build your own compare function to sort in your own regulations.
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

	rand.Seed(time.Now().Unix())
	randPivot := rand.Int()%(right-left) + left
	data[left], data[randPivot] = data[randPivot], data[left]

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

//QuickSortFolat32 is a QuickSort function for float32, sorting from smallest to biggest.
func QuickSortFloat32(data []float32) {
	qsHandleFloat32(data, 0, len(data)-1)
}

func qsHandleFloat32(data []float32, left, right int) {
	if left >= right {
		return
	}

	rand.Seed(time.Now().Unix())
	randPivot := rand.Int()%(right-left) + left
	data[left], data[randPivot] = data[randPivot], data[left]

	pivot := data[left]
	i, j := left+1, right

	for true {
		for i <= right {
			if data[i] > pivot {
				break
			}
			i++
		}

		for j > left {
			if data[j] < pivot {
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

	qsHandleFloat32(data, left, j-1)
	qsHandleFloat32(data, j+1, right)
}

//QuickSortFloat64 is a QuickSort function for float64, sorting from smallest to biggest.
func QuickSortFloat64(data []float64) {
	qsHandleFloat64(data, 0, len(data)-1)
}

func qsHandleFloat64(data []float64, left, right int) {
	if left >= right {
		return
	}

	rand.Seed(time.Now().Unix())
	randPivot := rand.Int()%(right-left) + left
	data[left], data[randPivot] = data[randPivot], data[left]

	pivot := data[left]
	i, j := left+1, right

	for true {
		for i <= right {
			if data[i] > pivot {
				break
			}
			i++
		}

		for j > left {
			if data[j] < pivot {
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

	qsHandleFloat64(data, left, j-1)
	qsHandleFloat64(data, j+1, right)
}

//QuickSortInt is a QuickSort function for int, sorting from smallest to biggest.
func QuickSortInt(data []int) {
	qsHandleInt(data, 0, len(data)-1)
}

func qsHandleInt(data []int, left, right int) {
	if left >= right {
		return
	}

	rand.Seed(time.Now().Unix())
	randPivot := rand.Int()%(right-left) + left
	data[left], data[randPivot] = data[randPivot], data[left]

	pivot := data[left]
	i, j := left+1, right

	for true {
		for i <= right {
			if data[i] > pivot {
				break
			}
			i++
		}

		for j > left {
			if data[j] < pivot {
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

	qsHandleInt(data, left, j-1)
	qsHandleInt(data, j+1, right)
}

//QuickSortInt32 is a QuickSort function for int32, sorting from smallest to biggest.
func QuickSortInt32(data []int32) {
	qsHandleInt32(data, 0, len(data)-1)
}

func qsHandleInt32(data []int32, left, right int) {
	if left >= right {
		return
	}

	rand.Seed(time.Now().Unix())
	randPivot := rand.Int()%(right-left) + left
	data[left], data[randPivot] = data[randPivot], data[left]

	pivot := data[left]
	i, j := left+1, right

	for true {
		for i <= right {
			if data[i] > pivot {
				break
			}
			i++
		}

		for j > left {
			if data[j] < pivot {
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

	qsHandleInt32(data, left, j-1)
	qsHandleInt32(data, j+1, right)
}

//QuickSortInt64 is a QuickSort function for int64, sorting from smallest to biggest.
func QuickSortInt64(data []int64) {
	qsHandleInt64(data, 0, len(data)-1)
}

func qsHandleInt64(data []int64, left, right int) {
	if left >= right {
		return
	}

	rand.Seed(time.Now().Unix())
	randPivot := rand.Int()%(right-left) + left
	data[left], data[randPivot] = data[randPivot], data[left]

	pivot := data[left]
	i, j := left+1, right

	for true {
		for i <= right {
			if data[i] > pivot {
				break
			}
			i++
		}

		for j > left {
			if data[j] < pivot {
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

	qsHandleInt64(data, left, j-1)
	qsHandleInt64(data, j+1, right)
}
