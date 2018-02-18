package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	b := []byte("x")
	var c1 [32]byte = sha256.Sum256([]byte("x"))
	var c2 [32]byte = sha256.Sum256([]byte("X"))
	fmt.Printf("d:%d b:%b\n", b, b)
	fmt.Printf("%X\n%X\n%t\n%T\n", c1, c2, c1 == c2, c1)
	fmt.Printf("%08b\n%08b\n", c1, c2)
	fmt.Printf("match count:%d\n", bitMatchCountSlow(c1, c2))
}

func bitMatchCountSlow(x [32]byte, y [32]byte) int {
	num := 0
	loop := 0
	var mask uint8 = 1

	for i := 0; i < 32; i++ {
		var j uint8
		fmt.Printf("x[i]:%X,y[i]:%X\n", x[i], y[i])
		for j = 0; j < 8; j++ {
			loop++
			if (x[i] & (mask << j)) == (y[i] & (mask << j)) {
				num++
			}
		}
	}
	fmt.Printf("loop cnter:%d\n", loop)
	return num
}
