package popcount

import (
	"encoding/binary"
	"fmt"
	"testing"
)

func TestExec(t *testing.T) {
	var tests = []struct {
		x uint64
	}{

		{0}, {1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}, {10}, {22}, {23},
		{255}, {256}, {1023}, {1024}, {1025}, {256 * 256}, {256*256 - 1}, {256*256 + 1},
		{256*256*256 - 1}, {256 * 256 * 256}, {256*256*256 + 1}, {256*256*256 + 256},
	}
	for _, test := range tests {
		b := make([]byte, 8)
		binary.LittleEndian.PutUint64(b, test.x)
		fmt.Printf("%v %d %d  %T:%v\n", test.x, PopCount(test.x), PopCountLoop(test.x), b, b)

		want := PopCount(test.x)

		if want != PopCountLoop(test.x) {
			t.Errorf("PopCountLoop:expected:%v atcual:%v\n", want, PopCountLoop(test.x))
		}
		if want != PopCountLoopArg(test.x) {
			t.Errorf("PopCountLoop:expected:%v atcual:%v\n", want, PopCountLoopArg(test.x))
		}
		if want != PopCountXminusOneLoop(test.x) {
			t.Errorf("PopCountLoop:expected:%v atcual:%v\n", want, PopCountXminusOneLoop(test.x))
		}

	}
}

var benchArgs []uint64 = []uint64{
	0, 1, 2, 3, 4, 5, 6, 7, 8, 255, 256, 256*256 - 1, 256 * 256, 256*256 + 1, 9, 10, 11, 12, 13,
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, arg := range benchArgs {
			PopCount(arg)
		}
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, arg := range benchArgs {
			PopCountLoop(arg)
		}
	}
}

func BenchmarkPopCountXminusOneLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, arg := range benchArgs {
			PopCountXminusOneLoop(arg)
		}
	}
}

func BenchmarkPopCountLoopArg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, arg := range benchArgs {
			PopCountLoopArg(arg)
		}
	}
}
