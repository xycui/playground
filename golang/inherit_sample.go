package main

import (
  "fmt";
)

type IBaseInterface interface{
  PrintData() error
  printData() error
}

type Impl1 struct{
  IBaseInterface
  data1 string
  Data string
}

func NewImpl1(data1 string, data string) *Impl1{
  return &Impl1{
    data1: data1,
    Data: data,
  }
}

func (imp *Impl1) PrintData() error{
  fmt.Println("Public:" + imp.Data)
  return nil
}

func (imp *Impl1) printData() error{
  fmt.Println("Private:"+ imp.data1)
  return nil
}

type Impl2 struct{
  Impl1
  newData string
}

func NewImpl2(data string, newData string) *Impl2{
  return &Impl2{
    Impl1: Impl1{
      data1: newData,
      Data: data,
    },
    newData: newData,
  }
}

//func (imp *Impl2)

func main(){
  imp := NewImpl1("test1", "test2")
  imp.PrintData()
  imp.printData()
  
  imp2 := NewImpl2("a", "b")
  imp2.PrintData()
  imp2.printData()
  fmt.Println(imp2.newData)
  fmt.Println(imp2.Data)
  data, ok := imp2.(Impl2)
  fmt.Println(ok)
}