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
	execute1(os.Stdin, counts)

	return nil
}

func execute1(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
