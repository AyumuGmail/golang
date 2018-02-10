package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for i, url := range os.Args[1:] {
		filename := "tmp/" + strconv.Itoa(i) + ".out"
		go fetch(url, filename, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, filename string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	if _, err := os.Stat(filename); err == nil {
		//ファイルが存在する場合
		filename = filename + ".add"
	}
	dst, err := os.Create(filename)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	defer dst.Close()

	//nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	nbytes, err := io.Copy(dst, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
