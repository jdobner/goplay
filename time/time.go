package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var week time.Duration

func main() {
	t := time.Now()
	fmt.Println("Time now:", t)
	fivemin, _ := time.ParseDuration("5m")
	fmt.Println("Time @5m:", t.Truncate(fivemin))
	reader := bufio.NewReader(os.Stdin)
	print("Change the clack and hit enter...\n")
	_, _ = reader.ReadString('\n')

	t2 := time.Now()
	fmt.Println(t2)
	d := t2.Sub(t)
	fmt.Println(d)

	f := func() {
		fmt.Println("fired")
    }

    twosec, _ := time.ParseDuration("2s")
    time.AfterFunc(twosec, f)
    fmt.Println("pausing for", twosec * 2)
    time.Sleep(twosec * 2)
    fmt.Println("done")

	// if 1 == 1 {
	// 	return
	// }

	fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())

	location, err := time.LoadLocation("america/new_york")
	if err != nil {
		panic(err)
	}
	fmt.Println(location)

	t = time.Now().UTC()
	fmt.Println(t)
	fmt.Println(time.Now())
	// calculating times:
	week = 60 * 60 * 24 * 7 * 1e9 // must be in nanosec
	week_from_now := t.Add(week)
	fmt.Println(week_from_now)
	// formatting times:
	fmt.Println("RFC822:", t.Format(time.RFC822))
	fmt.Println("ANSIC", t.Format(time.ANSIC))
	fmt.Println(t.Format("02 Jan 2006 15:04"))
	s := t.Format("2006 01 02")
	fmt.Println(t, "=>", s)
}