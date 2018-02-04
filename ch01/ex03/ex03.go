package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args
	s := concatArg(args)
	fmt.Println(s)
}

func concatArg(args []string) string {
	var s, sep string
	for i := 0; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	return s
}

func concatArgJoin(args []string) string {
	return strings.Join(args, " ")
}
