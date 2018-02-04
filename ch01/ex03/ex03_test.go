package main

import (
	"fmt"
	"testing"
)

func BenchmarkConcatArg(b *testing.B) {
	var args []string = []string{"aaa", "bbb", "ccc"}
	fmt.Printf("[exec nomal]\n")
	for i := 0; i < b.N; i++ {
		concatArg(args)
	}
}

func BenchmarkConcatStringJoin(b *testing.B) {
	var args []string = []string{"aaa", "bbb", "ccc"}
	fmt.Printf("[exec strings.Join]\n")
	for i := 0; i < b.N; i++ {
		concatArgJoin(args)
	}
}
