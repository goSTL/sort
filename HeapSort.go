package sort

import (
	"reflect"
)

func HeapSort(data interface{}, cmp func(i, j interface{}) bool) []interface{} {
	value := reflect.ValueOf(data)
	dataS := make([]interface{}, value.Len())
	for a := 0; a < value.Len(); a++ {
		dataS[a] = value.Index(a).Interface()
	}
	heapify(dataS, value.Len()/2-1, value.Len(), cmp)
	maxHeap(dataS, value.Len(), cmp)
	return dataS
}

func heapify(data []interface{}, now, last int, cmp func(i, j interface{}) bool) {
	if now >= last/2 || now < 0 {
		return
	}
	subHeapify(data, now, last, cmp)
	heapify(data, now-1, last, cmp)
	return
}

func subHeapify(data []interface{}, now, last int, cmp func(i, j interface{}) bool) {
	if (now*2+2) < last && !(cmp(data[now], data[now*2+1]) && cmp(data[now], data[now*2+2])) {
		var max int
		if cmp(data[now*2+1], data[now*2+2]) {
			max = now*2 + 1
		} else {
			max = now*2 + 2
		}
		data[now], data[max] = data[max], data[now]
		if max < (last / 2) {
			subHeapify(data, max, last, cmp)
		}
	} else if ((now*2 + 1) < last) && cmp(data[now*2+1], data[now]) {
		data[now], data[now*2+1] = data[now*2+1], data[now]
		if (now*2 + 1) < (last / 2) {
			subHeapify(data, now*2+1, last, cmp)
		}
	}
	return
}

func maxHeap(data []interface{}, len int, cmp func(i, j interface{}) bool) {
	if len <= 1 {
		return
	}
	data[0], data[len-1] = data[len-1], data[0]
	subHeapify(data, 0, len-1, cmp)
	maxHeap(data, len-1, cmp)
	return
}

// INT
func HeapSortInt(dataS []int) {
	heapifyInt(dataS, len(dataS)/2-1, len(dataS))
	maxHeapInt(dataS, len(dataS))
}

func heapifyInt(data []int, now, last int) {
	if now >= last/2 || now < 0 {
		return
	}
	subHeapifyInt(data, now, last)
	heapifyInt(data, now-1, last)
	return
}

func subHeapifyInt(data []int, now, last int) {
	if now*2+2 < last && !(data[now] >= data[now*2+1] && data[now] >= data[now*2+2]) {
		var max int
		if data[now*2+1] > data[now*2+2] {
			max = now*2 + 1
		} else {
			max = now*2 + 2
		}
		data[now], data[max] = data[max], data[now]
		if max < (last / 2) {
			subHeapifyInt(data, max, last)
		}
	} else if (now*2+1) < last && data[now] < data[now*2+1] {
		data[now], data[now*2+1] = data[now*2+1], data[now]
		if (now*2 + 1) < (last / 2) {
			subHeapifyInt(data, now*2+1, last)
		}
	}
	return
}

func maxHeapInt(data []int, len int) {
	if len <= 1 {
		return
	}
	data[0], data[len-1] = data[len-1], data[0]
	subHeapifyInt(data, 0, len-1)
	maxHeapInt(data, len-1)
	return
}

// INT32
func HeapSortInt32(dataS []int32) {
	heapifyInt32(dataS, len(dataS)/2-1, len(dataS))
	maxHeapInt32(dataS, len(dataS))
}

func heapifyInt32(data []int32, now, last int) {
	if now >= last/2 || now < 0 {
		return
	}
	subHeapifyInt32(data, now, last)
	heapifyInt32(data, now-1, last)
	return
}

func subHeapifyInt32(data []int32, now, last int) {
	if now*2+2 < last && !(data[now] >= data[now*2+1] && data[now] >= data[now*2+2]) {
		var max int
		if data[now*2+1] > data[now*2+2] {
			max = now*2 + 1
		} else {
			max = now*2 + 2
		}
		data[now], data[max] = data[max], data[now]
		if max < (last / 2) {
			subHeapifyInt32(data, max, last)
		}
	} else if (now*2+1) < last && data[now] < data[now*2+1] {
		data[now], data[now*2+1] = data[now*2+1], data[now]
		if (now*2 + 1) < (last / 2) {
			subHeapifyInt32(data, now*2+1, last)
		}
	}
	return
}

func maxHeapInt32(data []int32, len int) {
	if len <= 1 {
		return
	}
	data[0], data[len-1] = data[len-1], data[0]
	subHeapifyInt32(data, 0, len-1)
	maxHeapInt32(data, len-1)
	return
}

// INT64
func HeapSortInt64(dataS []int64) {
	heapifyInt64(dataS, len(dataS)/2-1, len(dataS))
	maxHeapInt64(dataS, len(dataS))
}

func heapifyInt64(data []int64, now, last int) {
	if now >= last/2 || now < 0 {
		return
	}
	subHeapifyInt64(data, now, last)
	heapifyInt64(data, now-1, last)
	return
}

func subHeapifyInt64(data []int64, now, last int) {
	if now*2+2 < last && !(data[now] >= data[now*2+1] && data[now] >= data[now*2+2]) {
		var max int
		if data[now*2+1] > data[now*2+2] {
			max = now*2 + 1
		} else {
			max = now*2 + 2
		}
		data[now], data[max] = data[max], data[now]
		if max < (last / 2) {
			subHeapifyInt64(data, max, last)
		}
	} else if (now*2+1) < last && data[now] < data[now*2+1] {
		data[now], data[now*2+1] = data[now*2+1], data[now]
		if (now*2 + 1) < (last / 2) {
			subHeapifyInt64(data, now*2+1, last)
		}
	}
	return
}

func maxHeapInt64(data []int64, len int) {
	if len <= 1 {
		return
	}
	data[0], data[len-1] = data[len-1], data[0]
	subHeapifyInt64(data, 0, len-1)
	maxHeapInt64(data, len-1)
	return
}

// FLOAT32
func HeapSortFloat32(dataS []float32) {
	heapifyFloat32(dataS, len(dataS)/2-1, len(dataS))
	maxHeapFloat32(dataS, len(dataS))
}

func heapifyFloat32(data []float32, now, last int) {
	if now >= last/2 || now < 0 {
		return
	}
	subHeapifyFloat32(data, now, last)
	heapifyFloat32(data, now-1, last)
	return
}

func subHeapifyFloat32(data []float32, now, last int) {
	if now*2+2 < last && !(data[now] >= data[now*2+1] && data[now] >= data[now*2+2]) {
		var max int
		if data[now*2+1] > data[now*2+2] {
			max = now*2 + 1
		} else {
			max = now*2 + 2
		}
		data[now], data[max] = data[max], data[now]
		if max < (last / 2) {
			subHeapifyFloat32(data, max, last)
		}
	} else if (now*2+1) < last && data[now] < data[now*2+1] {
		data[now], data[now*2+1] = data[now*2+1], data[now]
		if (now*2 + 1) < (last / 2) {
			subHeapifyFloat32(data, now*2+1, last)
		}
	}
	return
}

func maxHeapFloat32(data []float32, len int) {
	if len <= 1 {
		return
	}
	data[0], data[len-1] = data[len-1], data[0]
	subHeapifyFloat32(data, 0, len-1)
	maxHeapFloat32(data, len-1)
	return
}

// FLOAT64
func HeapSortFloat64(dataS []float64) {
	heapifyFloat64(dataS, len(dataS)/2-1, len(dataS))
	maxHeapFloat64(dataS, len(dataS))
}

func heapifyFloat64(data []float64, now, last int) {
	if now >= last/2 || now < 0 {
		return
	}
	subHeapifyFloat64(data, now, last)
	heapifyFloat64(data, now-1, last)
	return
}

func subHeapifyFloat64(data []float64, now, last int) {
	if now*2+2 < last && !(data[now] >= data[now*2+1] && data[now] >= data[now*2+2]) {
		var max int
		if data[now*2+1] > data[now*2+2] {
			max = now*2 + 1
		} else {
			max = now*2 + 2
		}
		data[now], data[max] = data[max], data[now]
		if max < (last / 2) {
			subHeapifyFloat64(data, max, last)
		}
	} else if (now*2+1) < last && data[now] < data[now*2+1] {
		data[now], data[now*2+1] = data[now*2+1], data[now]
		if (now*2 + 1) < (last / 2) {
			subHeapifyFloat64(data, now*2+1, last)
		}
	}
	return
}

func maxHeapFloat64(data []float64, len int) {
	if len <= 1 {
		return
	}
	data[0], data[len-1] = data[len-1], data[0]
	subHeapifyFloat64(data, 0, len-1)
	maxHeapFloat64(data, len-1)
	return
}
