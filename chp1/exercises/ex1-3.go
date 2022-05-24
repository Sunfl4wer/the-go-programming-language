package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var noIteration = 40000

	fmt.Println("Process A started")
	var start = time.Now()
	var sep = ""
	var s = ""
	for i := 0; i < noIteration; i++ {
		s += sep + "mew"
		sep = " "
	}
	fmt.Println(s)
	var secs = time.Since(start).Seconds()
	fmt.Println("Process A stopped")
	fmt.Printf("Execution time: %.2fs\n", secs)

	var mews = make([]string, noIteration)
	for i := 0; i < noIteration; i++ {
		mews[i] = "mew"
	}

	fmt.Println("Process B started")
	start = time.Now()
	fmt.Println(strings.Join(mews, " "))
	secs = time.Since(start).Seconds()
	fmt.Println("Process B stopped")
	fmt.Printf("Execution time: %.2fs\n", secs)
}