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
	if _, err := executeProc(flag.Args(), os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "echo: %v\n", err)
		os.Exit(1)
	}
}

func executeProc(args []string, writer io.Writer) ([]int, error) {
	var statusCodes []int
	for _, url := range args[0:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			return statusCodes, fmt.Errorf("fetch %v\n", err)
		}

		_, err = io.Copy(writer, resp.Body) //ここはイコール!!
		fmt.Printf("Fetch:%s  HTTP STATUS CODE:%d\n", url, resp.StatusCode)
		if err != nil {
			return statusCodes, fmt.Errorf("fetch:reading %s:%v\n", url, err)
		} else {
			statusCodes = append(statusCodes, resp.StatusCode)
		}
		resp.Body.Close()
	}
	return statusCodes, nil
}
