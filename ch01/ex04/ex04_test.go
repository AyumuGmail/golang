package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestExecuteProc(t *testing.T) {
	var tests = []struct {
		targetfilenames  []string
		resuletFilenames []string
	}{
		{[]string{"depfile", "depfile2", "no_depfile"},
			[]string{"depfile", "depfile2"}},
	}
	for _, test := range tests {
		results, _ := executeProc(test.targetfilenames)

		if len(results) != 2 {
			for _, result := range results {
				fmt.Printf("%s\n", result)
			}
			t.Errorf("Err expected=2files actual=[%d]files", len(results))
		}
		for i, resultfile := range results {
			if resultfile != test.targetfilenames[i] {
				t.Errorf("fileUnmatch %s,[%d],%s", resultfile, i, test.targetfilenames[i])
			}
		}
	}
}

func TestCountLine(t *testing.T) {
	var tests = []struct {
		lines     string
		wantCount int
	}{
		{"test\ntest", 1},
		{"test\ntest2", 2},
		{"test\ntest\ntest", 1},
	}
	{
		for _, test := range tests {
			//ファイルポインタを用いたテストは、ファイルがないとNGか。
			counts := make(map[string]int)
			tempfilname := "tempfile"
			ioutil.WriteFile(tempfilname, []byte(test.lines), os.ModePerm)
			f, err := os.Open(tempfilname)
			if err != nil {
				t.Error("system Error!")
			}
			countLine(f, counts)
			if len(counts) != test.wantCount {
				t.Errorf("countLine(%d) != %d %v \n", len(counts), test.wantCount, test.lines)
			}
			f.Close()
		}
	}
}
