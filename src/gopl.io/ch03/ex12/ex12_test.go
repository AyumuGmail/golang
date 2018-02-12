package word

import (
	"fmt"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	var tests = []struct {
		s1   string
		s2   string
		want bool
	}{
		{"aba", "baa", true},
		{"ab", "abc", false},
		{"abc", "ab", false},
		{"abc", "ab", false},
		{"山本 海", "海山本", true},
	}
	var sucessCnt int
	for _, test := range tests {
		if IsAnagram(test.s1, test.s2) != test.want {
			t.Errorf("s1:%s s2:%s", test.s1, test.s2)
		} else {
			sucessCnt++
		}
	}
	fmt.Printf("SuccessCnt:%d\n", sucessCnt)
}
