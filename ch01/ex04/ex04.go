package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	if err := executeProc(flag.Args()); err != nil {
		fmt.Fprintf(os.Stderr, "echo: %v\n", err)
		os.Exit(1)
	}
}

func executeProc(args []string) error {
	files := args[0:] //flagのargの場合は、本当の引数からっぽい
	//var duplicatedLineFiles []string

	if len(files) == 0 {
		IsDuplicationLineFile(os.Stderr)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dub2: %v\n", err)
				continue
			}
			IsDuplicationLineFile(f)
			f.Close()
		}
	}

	return nil
}

func IsDuplicationLineFile(f *os.File) bool {
	counts := make(map[string]int)
	countLine(f, counts)
	for _, n := range counts {
		if n > 1 {
			return true
		}
	}
	return false
}

func countLine(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
