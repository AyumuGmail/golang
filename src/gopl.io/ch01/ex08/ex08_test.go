package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestExecuteProc(t *testing.T) {
	var tests = []struct {
		urls []string
		err  error
	}{
		{[]string{"http://gopl.io/"}, nil},
		{[]string{"http://bad.gopl.io/"}, errors.New("hoge")},
		{[]string{"test"}, errors.New("hoge")},
		{[]string{"gopl.io/"}, nil},
	}
	for _, test := range tests {
		err := executeProc(test.urls)
		if err != nil {
			if test.err == nil {
				t.Errorf("err %v\n", test.urls)
			} else {
				fmt.Printf("ErrMessage:%v", err)
			}
			//それ以外はSuccess
		} else {
			if test.err != nil {
				t.Errorf("err %v\n", test.urls)
			}
		}
	}
}
