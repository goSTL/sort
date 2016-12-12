package sort

import(
  "reflect"
)

//This is a custom Merge Sort with parallel compution in 8 threads. 
//You can define the comparion regulations all by yourself.
//
//You should use runtime.GOMAXPROS(int) to set the maximum number of CPUs that can be executing simultaneously.
func MergeSortPara(data interface{}, cmp func(i,j interface{})bool)[]interface{}{
  //data : interface{} to []interface
  value:=reflect.ValueOf(data)
  N:=make([]interface{},value.Len())
  for a:=0;a<value.Len();a++{
    N[a]=value.Index(a).Interface()
  }
  ch:=make(chan bool)
  go msHandlePara8(N,cmp,ch,0)
  <-ch
  return N
}

func msHandlePara8(data []interface{}, cmp func(i,j interface{})bool, chparent chan bool,count int){
  M:=len(data)/2
  if len(data)==1 {
    chparent<-true
    return
  }
  if count<4 {
    chchild1:=make(chan bool)
    chchild2:=make(chan bool)
    go msHandlePara8(data[:M],cmp,chchild1,count+1)
    go msHandlePara8(data[M:],cmp,chchild2,count+1)
    <-chchild1
    <-chchild2
  }else{
    msHandle(data[:M],cmp)
    msHandle(data[M:],cmp)
  }
  merge(data,M,cmp)
  chparent<-true
}
