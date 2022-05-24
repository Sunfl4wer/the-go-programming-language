package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	filesWithDuplicates := make(map[string]map[string]bool)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, filesWithDuplicates, "")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, filesWithDuplicates, arg)
			f.Close()
		}
		for line, n := range counts {
			if n > 1 {
				files := filesWithDuplicates[line]
				var filenames string
				for filename,exists := range files {
					if (exists) {
						filenames += filename
						filenames += " "
					}
				}
				fmt.Printf("%d\t%s\t%s\n", n, line, filenames)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int, filesWithDuplicates map[string]map[string]bool, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		text := input.Text()
		counts[input.Text()]++
		addFilename(text, counts, filesWithDuplicates, filename)
	}
}

func addFilename(text string, counts map[string]int, filesWithDuplicates map[string]map[string]bool, filename string) {
	if (counts[text] > 1) {
		files := filesWithDuplicates[text]
		if (files == nil) {
			filesWithDuplicates[text] = make(map[string]bool)
			filesWithDuplicates[text][filename] = true
		}
		if (!filesWithDuplicates[text][filename]) {
			filesWithDuplicates[text][filename] = true
		}
	}
}