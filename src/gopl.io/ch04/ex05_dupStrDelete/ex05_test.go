package main

import (
	"testing"
)

func TestDupDelete(t *testing.T) {
	var tests = []struct {
		s1       []string
		expected []string
	}{
		{[]string{"aba", "baa"}, []string{"aba", "baa"}},
		{[]string{"", "", "kataoka", "kataoka", "test", "kataoka"}, []string{"", "kataoka", "test", "kataoka"}},
	}

	for _, test := range tests {
		actualStrs := dupDelete(test.s1)
		for i, s := range actualStrs {
			if s != test.expected[i] {
				t.Errorf("err:actual %s,expected %s %q\n", s, test.expected[i], s)
			}
		}
	}
}
