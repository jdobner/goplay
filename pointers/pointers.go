package main

import (
	"fmt"
	"unsafe"
)

func main() {
	i := 99
	fmt.Println("original val", i)
	incr(&i)
	fmt.Println("new val", i)
	i2, _ := newInt()
	fmt.Println(i2, *i2)

}

func incr(i *int) {
	*i = *i + 1
	fmt.Println(i, *i)
	//	&i = &i + 1;
}

func newInt() (*int, *int) {
	i1 := 11
	i2 := 22
	i3 := 33
	i4 := 0
	i5 := 11
	i4 = i3 + i5

	fmt.Println(&i1, i1)
	fmt.Println(&i2, i2)
	fmt.Println(&i3, i3)
	fmt.Println(&i4, i4)
	fmt.Print(&i5, i5, "\n\n")

	formatPointer(&i1, &i1)
	formatPointer(&i1, &i2)
	formatPointer(&i2, &i3)
	formatPointer(&i3, &i4)
	//formatPointer(&i4, &i5)
	return &i1, &i2

}

func formatPointer(old *int, newp *int) {
	diff := diff(old, newp)
	fmt.Println(diff, newp, *newp)

}

func diff(old *int, newp *int) uintptr {
	p1 := unsafe.Pointer(old)
	i1 := uintptr(p1)
	p2 := unsafe.Pointer(newp)
	i2 := uintptr(p2)

	return i2 - i1

}
