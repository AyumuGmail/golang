package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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
		resp, err := http.Get(url)
		if err != nil {
			return fmt.Errorf("fetch %v\n", err)
		}

		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return fmt.Errorf("fetch:reading %s:%v\n", url, err)
		}
		fmt.Printf("%s", b)
	}
	return nil
}
