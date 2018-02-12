package basename

import "testing"

func TestBasename1(t *testing.T) {
	var tests = []struct {
		arg      string
		expected string
	}{
		{"a", "a"},
		{"a.go", "a"},
		{"a/b/c.go", "c"},
		{"a/b.c.go", "b.c"},
	}
	for _, test := range tests {
		if basename1(test.arg) != test.expected {
			t.Errorf("expected:%s actual:%s", test.expected, basename1(test.arg))
		}
	}
}
