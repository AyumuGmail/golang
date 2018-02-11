package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	n = flag.Bool("n", false, "omit trailing newline")
	s = flag.String("s", " ", "separator")
)

var out io.Writer = os.Stdout

func main() {
	flag.Parse()
	input := bufio.NewScanner(os.Stdin)
	if err := executeProc(input); err != nil {
		fmt.Fprintf(os.Stderr, "echo: %v\n", err)
		os.Exit(1)
	}
}

func executeProc(input *bufio.Scanner) error {
	counts := make(map[string]int)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
	return nil
}
