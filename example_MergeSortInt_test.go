package sort_test

import (
	"fmt"
	"github.com/goSTL/sort"
)

func ExampleMergeSortInt() {
	sample := []int{5, 7, 4, 10, 99, 2, 72, 37}
	fmt.Println(sample)

	sort.MergeSortInt(sample)
	fmt.Println(sample)
	//Output:
	//[5 7 4 10 99 2 72 37]
	//[2 4 5 7 10 37 72 99]
}
