package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	data := "山本山　山本"
	b := []byte(data)
	reverse(b)
	fmt.Println(b)
}
func reverse(bs []byte) {
	//for i, j := 0, len(bs)-1; i < j; i, j = i+1, j-1 {
	for i := 0; i < len(bs); {
		r, size := utf8.DecodeRune(bs[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i = i + size
	}
}
