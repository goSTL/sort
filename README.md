[![MIT Licence](https://badges.frapsoft.com/os/mit/mit.svg?v=103)](https://opensource.org/licenses/mit-license.php) 
[![GoDoc](https://godoc.org/github.com/go-Algorithm/sort?status.svg)](https://godoc.org/github.com/go-Algorithm/sort)
[![Build Status](https://travis-ci.org/go-Algorithm/sort.svg?branch=master)](https://travis-ci.org/go-Algorithm/sort)

Sort
====
sort whatever you want

#Index
-   [Merge Sort](https://github.com/goSTL/sort#merge-sort)
-   [Quick Sort](https://github.com/goSTL/sort#quick-sort)

#Merge Sort
**MergeSort()**

MergeSort() is a custom Merge Sort function. You can define the comparion regulations all by yourself.

> You should write your own compare function with definition below.
> `func(i,j interface{})bool`
> (you can use i.(YourInputType) to convert i to original type)
>  
>  The output will be a slice of interface{},you should use o.(YourInputType) to convert o to origibal type

here is a sample code:
```go
package main
    
import (
  "github.com/goSTL/sort"
  "fmt"
)

type student struct{
  id    int
  name  string
}

func cmp(i,j interface{})bool{
  ii:=i.(student)
  jj:=j.(student)
  return ii.id < jj.id
}

func main(){
  stu:=[]student{{5,"sam"},{3,"lily"},{7,"jacky"},{1,"willy"},{2,"steve"}}
  fmt.Println("original: \t",stu)

  out:=sort.MergeSort(stu,cmp)
  fmt.Println("[]interface{}: \t",out)

  stu2:=make([]student,len(out))
  for a:=0;a<len(out);a++{
    stu2[a]=out[a].(student)
  }
  fmt.Println("[]student: \t",stu2)
}
```    
output:

    original: 	 [{5 sam} {3 lily} {7 jacky} {1 willy} {2 steve}]
    []interface{}: 	 [{1 willy} {2 steve} {3 lily} {5 sam} {7 jacky}]
    []student: 	 [{1 willy} {2 steve} {3 lily} {5 sam} {7 jacky}]


#Quick Sort
**QuickSort()**

QuickSort() is a custom quick sort.

> You should write your own compare function with definition below.
> `func(i,j interface{})bool`
> (you can use i.(YourInputType) to convert i to original type)
>  
>  The output will be a slice of interface{},you should use o.(YourInputType) to convert o to origibal type

here is a sample code:
```go
package main
    
import (
  "github.com/goSTL/sort"
  "fmt"
)

type student struct{
  id    int
  name  string
}

func cmp(i,j interface{})bool{
  ii:=i.(student)
  jj:=j.(student)
  return ii.id < jj.id
}

func main(){
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
```    
output:

    original: 	 [{5 sam} {3 lily} {7 jacky} {1 willy} {2 steve}]
    []interface{}: 	 [{1 willy} {2 steve} {3 lily} {5 sam} {7 jacky}]
    []student: 	 [{1 willy} {2 steve} {3 lily} {5 sam} {7 jacky}]


