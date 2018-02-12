package main

import (
	"fmt"
)

const (
	_ = 1 << (10 * iota)
	KB
	MB
	GB
	TB
	PB
)

const (
	KiB = 1000
	MiB = KiB * KiB
	GiB = MiB * KiB
	TiB = GiB * KiB
	PiB = TiB * KiB
)

func main() {
	fmt.Printf("KB:%d\n", KB)
	fmt.Printf("MB:%d\n", MB)
	fmt.Printf("GB:%d\n", GB)
	fmt.Printf("TB:%d\n", TB)
	fmt.Printf("PB:%d\n", PB)
	fmt.Printf("KB:%d\n", KiB)
	fmt.Printf("MB:%d\n", MiB)
	fmt.Printf("GB:%d\n", GiB)
	fmt.Printf("TB:%d\n", TiB)
	fmt.Printf("PB:%d\n", PiB)
}
