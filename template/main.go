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
	filenames := args[0:] //flagのargの場合は、本当の引数からっぽい
	var counts map[string]int
	err := execute1(filenames, counts)
	if err != nil {
		return fmt.Errorf("err")
	}
	return nil
}

func execute1(filenames []string, counts map[string]int) error {
	for _, filename := range filenames {
		f, _ := os.Open(filename) //err処理は省略
		input := bufio.NewScanner(f)
		for input.Scan() {
			counts[input.Text()]++
		}
	}
	return nil
}
