package main

import "fmt"

type MyType int64

func main() {
	mt := MyType(5)
	fmt.Println(mt, "+ 2 =", mt.Add(2))
}

func (m MyType) Add(v int64) MyType {
	return MyType(int64(m) + v)
}
