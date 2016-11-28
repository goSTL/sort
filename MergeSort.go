package sort

import(
  "reflect"
)

//this is a custom Merge Sort. you can define the comparion regulations all by yourself
//
//you should write your own compare function with definition 
//" func(i,j interface{})bool "
//then MergeSort() will launch Merge Sort with the regulation you define
//(*you can use i.(YourInputType) to convert i to original type)
//
//the output will be a slice of interface{},you should use o.(YourInputType) to convert o to origibal type
func MergeSort(data interface{}, cmp func(i,j interface{})bool)[]interface{}{
  //data : interface{} to []interface
  value:=reflect.ValueOf(data)
  N:=make([]interface{},value.Len())
  for a:=0;a<value.Len();a++{
    N[a]=value.Index(a).Interface()
  }
  msHandle(N,cmp)
  return N
}

func msHandle(data []interface{}, cmp func(i,j interface{})bool){
  M:=len(data)/2
  if len(data)==1 {
    return
  }
  msHandle(data[:M],cmp)
  msHandle(data[M:],cmp)
  merge(data,M,cmp)
}

func merge(data []interface{}, M int, cmp func(i,j interface{})bool){
  tmp:=make([]interface{},len(data))
  length:=len(data)
  var a,b,p int = 0,0,0
  for a,b=0,M ; a<M && b<length ; {
    if cmp(data[a],data[b]) {
      tmp[p]=data[a]
      p++
      a++
    }else{
      tmp[p]=data[b]
      p++
      b++
    }
  }
  if a==M {
    for i:=b ; i<length ; i,p=i+1,p+1 {
      tmp[p]=data[i]
    }
  }else{
    for i:=a ; i<M ; i,p=i+1,p+1 {
      tmp[p]=data[i]
    }
  }
  for i:=0 ; i<length ; i++ {
    data[i]=tmp[i]
  }
}
