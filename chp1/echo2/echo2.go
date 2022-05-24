package main

import (
	"fmt"
	"os"
)

func main() {
	var s string
	var sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}