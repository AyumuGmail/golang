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
		if comma(test.arg) != test.expected {
			t.Errorf("expected:%s actual:%s", test.expected, comma(test.arg))
		} else {
			fmt.Printf("%s\n", test.expected)
		}
	}
}
