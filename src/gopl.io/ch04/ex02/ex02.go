package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io"
	"os"
)

var shatype *int = flag.Int("t", 256, "sha256,sha384,sha512")

func main() {
	flag.Parse()
	fmt.Printf("t:%d\n", *shatype)
	if err := executeProc(flag.Args(), *shatype); err != nil {
		fmt.Fprintf(os.Stderr, "echo: %v\n", err)
		os.Exit(1)
	}
}

func executeProc(args []string, t int) error {
	input := bufio.NewScanner(os.Stdin)
	var str string
	for input.Scan() {
		str = input.Text()
	}
	//b := []byte("x")
	fmt.Printf("T:%d\n", t)
	if t == 256 {
		fmt.Printf("sha256\n")
		c1 := sha256.Sum256([]byte(str))
		fmt.Printf("%X\n%T\n", c1, c1)
	} else if t == 384 {
		fmt.Printf("sha384\n")
		s := sha512.New384()
		io.WriteString(s, str)
		fmt.Printf("%X\n%T\n", s.Sum(nil), s.Sum(nil))
	} else if t == 512 {
		fmt.Printf("sha512\n")
		c1 := sha512.Sum512([]byte(str))
		fmt.Printf("%X\n%T\n", c1, c1)
	} else {
		fmt.Printf("%d notfound,default sha256\n", t)
		c1 := sha256.Sum256([]byte(str))
		fmt.Printf("%X\n%T\n", c1, c1)
	}
	return nil
}

func countLine(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
