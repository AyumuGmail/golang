package printints

import (
	"fmt"
	"testing"
)

func TestPrintints(t *testing.T) {
	var tests = []struct {
		arg      []int
		expected string
	}{
		{[]int{1, 2, 3}, "[1,2,3]"},
	}
	for _, test := range tests {
		if intsToString(test.arg) != test.expected {
			t.Errorf("expected:%s actual:%s", test.expected, intsToString(test.arg))
		} else {
			fmt.Printf("%s\n", test.expected)
		}
	}
}
