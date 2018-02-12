package comma

import (
	"fmt"
	"testing"
)

func TestFloat(t *testing.T) {
	var tests = []struct {
		arg      float64
		expected string
	}{
		{100, "100"},
		{1000, "1,000"},
		{1000000000, "1,000,000,000"},
		{-100, "-100"},
		{-1000, "-1,000"},
		{-2.1000, "-2.1"},
		{-1002.10000, "-1,002.1"},
		{-1002.10001, "-1,002.10001"},
	}
	for _, test := range tests {
		if commaFloat(test.arg) != test.expected {
			t.Errorf("expected:%s actual:%s", test.expected, commaFloat(test.arg))
		} else {
			fmt.Printf("%s\n", test.expected)
		}
	}
}
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
