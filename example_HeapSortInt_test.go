package sort_test

import (
	"fmt"

	"github.com/goSTL/sort"
)

func main() {
	sample := []int{3, 138, 1, 674, 213, 23, 5, 2}
	fmt.Println(sample)

	sort.HeapSortInt(sample)
	fmt.Println(sample)
	//Output:
	//[3 138 1 674 213 23 5 2]
	//[1 2 3 5 23 138 213 674]
}
