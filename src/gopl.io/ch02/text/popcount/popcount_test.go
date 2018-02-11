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
	}
}
