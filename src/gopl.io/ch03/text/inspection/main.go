package main

import "fmt"

func main() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2
	var z uint64 = 255
	var zz uint64 = 1000000000

	fmt.Printf("x\t%08b\t%d\t%o\t%x\n", x, x, x, x)
	fmt.Printf("x<<1\t%08b\t%d\t%o\t%x\n", x<<1, x<<1, x<<1, x<<1)
	fmt.Printf("x<<2\t%08b\t%d\t%o\t%x\n", x<<2, x<<2, x<<2, x<<2)
	fmt.Printf("x>>1\t%08b\t%d\t%o\t%x\n", x>>1, x>>1, x>>1, x>>1)
	fmt.Printf("y\t%08b\t%d\t%o\t%x\n", y, y, y, y)
	fmt.Printf("y<<1\t%08b\t%d\t%o\t%x\n", y<<1, y<<1, y<<1, y<<1)
	fmt.Printf("z\t%064b\t%d\t%o\t%x\n", z, z, z, z)
	fmt.Printf("zz\t%064b\t%d\t%o\t%x\n", zz, zz, zz, zz)

	var out string
	out = fmt.Sprintf("%b", z)
	//var num int
	var i uint = 0
	for ; i < 64; i++ {
		fmt.Print(zz >> i & 1)
	}
	fmt.Println()
	fmt.Println(out)

}
