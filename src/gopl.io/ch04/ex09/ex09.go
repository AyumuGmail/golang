package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	flag.Parse()
	filename := flag.Args()[0]
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("file not open:%s", filename)
		os.Exit(1)
	}
	counts, err := executeProc(f, os.Stdout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "echo: %v\n", err)
		os.Exit(1)
	}
	for c, n := range counts {
		fmt.Printf("[%s]:%d\n", c, n)
	}
}
func executeProc(f *os.File, out io.Writer) (map[string]int, error) {
	counts := make(map[string]int)
	input := bufio.NewScanner(f)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		counts[input.Text()]++
	}
	return counts, nil
}
