package main

import (
	"testing"
)

func TestCompspace(t *testing.T) {
	var tests = []struct {
		s        []byte
		expected []byte
	}{
		{[]byte("世界で　一番"), []byte("世界で 一番")},
		{[]byte("カタオカ　　　アユム"), []byte("カタオカ   アユム")},
		{[]byte("カタオカ　　　"), []byte("カタオカ   ")},
	}

	for _, test := range tests {
		bs := compSpace(test.s)
		for i := 0; i < len(bs); i++ {
			if bs[i] != test.expected[i] {
				t.Errorf("err actual:%s expected:%s\n", test.s, test.expected)
			}
		}
	}
}
