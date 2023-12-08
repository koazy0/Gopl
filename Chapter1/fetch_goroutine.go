package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	ch := make(chan string)

	start := time.Now()

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for _ = range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

}

func fetch(url string, ch chan string) bool {

	start := time.Now()
	if !strings.HasPrefix(url, "h") && !strings.HasPrefix(url, "h") {
		url = "https://" + url
	}
	resp, err := http.Get(url)
	defer resp.Body.Close()

	if err != nil {
		println("url:", url, " err:", err)
		return false
	}

	// Discard is a Writer on which all Write calls succeed
	// without doing anything.
	nbyte, _ := io.Copy(io.Discard, resp.Body)

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbyte, url)
	return true
}
