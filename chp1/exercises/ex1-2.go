package main

import (
	"fmt"
	"os"
)

func main() {
	var s string
	for idx, arg := range os.Args[:] {
		fmt.Printf("%d %s\n", idx, arg)
	}
}