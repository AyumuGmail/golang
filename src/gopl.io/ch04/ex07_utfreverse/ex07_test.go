package main

import (
	"testing"
)

func TestCompspace(t *testing.T) {
	var tests = []struct {
		s        []byte
		expected []byte
	}{
		{[]byte("世界で　一番"), []byte("番一　で界世")},
		{[]byte("abc片岡カタオカ　　　"), []byte("　　　カオタカ岡片cba")},
	}

	for _, test := range tests {
		bs := reverse(test.s)
		for i := 0; i < len(bs); i++ {
			if bs[i] != test.expected[i] {
				t.Errorf("err actual:%s expected:%s\n", test.s, test.expected)
			}
		}
	}
}
