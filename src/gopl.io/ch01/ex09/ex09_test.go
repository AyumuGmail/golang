package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestExecuteProc(t *testing.T) {
	var tests = []struct {
		urls     []string
		statuses []int
		err      error
	}{
		{[]string{"http://gopl.io/"}, []int{200}, nil},
		{[]string{"http://bad.gopl.io/"}, nil, errors.New("hoge")},
		{[]string{"test"}, nil, errors.New("hoge")},
		{[]string{"gopl.io/"}, []int{200}, nil},
		{[]string{"http://gopl.io/test/"}, []int{404}, nil},
	}
	for _, test := range tests {
		statuses, err := executeProc(test.urls, ioutil.Discard)
		if err != nil {
			if test.err == nil {
				t.Errorf("err %v \n", test.urls)
			} else {
				fmt.Printf("ErrMessage:%v", err)
			}
		} else {
			for i, _ := range statuses {
				if statuses[i] != test.statuses[i] {
					t.Errorf("err %v acutual:%d,expected:%d\n", test.urls, statuses[i], test.statuses[i])
				}
			}
		}
	}
}
