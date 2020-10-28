package main

import (
	"fmt"
	"time"
)

func main() {
	if d, err := time.ParseDuration("1s"); err != nil {
		panic(err)
	} else {
		fmt.Println(d)
	}
	a := 1
	a, b := 2, 3
	println(a, b)

}

func doOrErr() (d time.Duration, err error) {
	d, err = time.ParseDuration("1s")
	return
}
