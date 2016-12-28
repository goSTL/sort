package sort_test

import (
  "github.com/goSTL/sort"
  "fmt"
)

func ExampleQuickSortInt(){
  sample:=[]int{5,7,4,10,99,2,72,37}
  fmt.Println(sample)

  sort.QuickSortInt(sample)
  fmt.Println(sample)
  //Output:
  //[5 7 4 10 99 2 72 37]
  //[2 4 5 7 10 37 72 99]
}
