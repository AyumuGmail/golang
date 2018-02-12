package comma

import (
	"fmt"
	"testing"
)

func TestComma(t *testing.T) {
	var tests = []struct {
		arg      string
		expected string
	}{
		{"100", "100"},
		{"1000", "1,000"},
		{"1000000000", "1,000,000,000"},
	}
	for _, test := range tests {
		if commaBuf(test.arg) != test.expected {
			t.Errorf("expected:%s actual:%s", test.expected, commaBuf(test.arg))
		} else {
			fmt.Printf("%s\n", test.expected)
		}
	}
}

var benchArgs = []string{"100", "1000", "100000", "1000000"}

func BenchmarkComma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, arg := range benchArgs {
			comma(arg)
		}
	}
}

func BenchmarkCommaBuf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, arg := range benchArgs {
			commaBuf(arg)
		}
	}
}
