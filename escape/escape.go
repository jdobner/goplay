package main

import "fmt"

func main() {
	a := new(int)
	fmt.Println("value of a: ", *a)
	b := new(int)
	sum := *a + *b
	fmt.Println("sum = ", sum)
}
