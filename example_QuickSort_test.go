package sort_test

import (
  "github.com/goSTL/sort"
  "fmt"
)

func ExampleQuickSort(){
  stu:=[]student{{5,"sam"},{3,"lily"},{7,"jacky"},{1,"willy"},{2,"steve"}}
  fmt.Println("original: \t",stu)

  out:=sort.QuickSort(stu,cmp)
  fmt.Println("[]interface{}: \t",out)

  stu2:=make([]student,len(out))
  for a:=0;a<len(out);a++{
    stu2[a]=out[a].(student)
  }
  fmt.Println("[]student: \t",stu2)
}
