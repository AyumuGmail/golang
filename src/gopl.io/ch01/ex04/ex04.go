package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	if results, err := executeProc(flag.Args()); err != nil {
		fmt.Fprintf(os.Stderr, "echo: %v\n", err)
		os.Exit(1)
	} else {
		if len(results) > 0 {
			for _, fileName := range results {
				fmt.Println(fileName)
			}
		} else {
			fmt.Println("No duplicatedLine file")
		}
		os.Exit(0)
	}
}

func executeProc(args []string) (result []string, err error) {
	files := args[0:] //flagのargの場合は、本当の引数からっぽい
	var duplicatedLineFiles []string

	if len(files) == 0 {
		if isDuplicationLineFile(os.Stdin) {
			duplicatedLineFiles = append(duplicatedLineFiles, "STDIN")
		}
	} else {
		for _, arg := range files {
			//fmt.Printf("open:%s\n", arg)
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dub2: %v\n", err)
				continue
			}
			if isDuplicationLineFile(f) {
				duplicatedLineFiles = append(duplicatedLineFiles, arg)
				//fmt.Printf("deplicatedLineFile=%s\n", arg)
			}
			f.Close()
		}
	}
	result = duplicatedLineFiles
	return result, nil
}

func isDuplicationLineFile(f *os.File) bool {
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
