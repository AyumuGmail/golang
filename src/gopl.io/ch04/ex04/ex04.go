package main

import (
	"fmt"
)

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	//v := [6]int{0, 1, 2, 3, 4, 5}
	b := rotate(a, 3)
	fmt.Println(b)
}

func rotate(s []int, d int) []int {
	s2 := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		if i+d < len(s) {
			s2[i+d] = s[i]
		} else {
			s2[i+d-len(s)] = s[i]
		}
	}
	return s2
}
