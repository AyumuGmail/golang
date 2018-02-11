package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	flag.Parse()
	if err := executeProc(flag.Args()); err != nil {
		fmt.Fprintf(os.Stderr, "echo: %v\n", err)
		os.Exit(1)
	}
}

func executeProc(args []string) error {
	for _, url := range args[0:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			return fmt.Errorf("fetch %v\n", err)
		}

		_, err = io.Copy(os.Stdout, resp.Body) //ここはイコール!!
		resp.Body.Close()
		if err != nil {
			return fmt.Errorf("fetch:reading %s:%v\n", url, err)
		}
	}
	return nil
}
