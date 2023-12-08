// Fetch prints the content found at a URL
// Testing go run fetch.go www.baidu.com >> baidu.html
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// Example
func main1() {

	// Loop to get URL
	for _, url := range os.Args[1:] {

		resp, err := http.Get(url)

		if err != nil {
			println("err:", err)
			continue
		}

		data, err1 := io.ReadAll(resp.Body)
		resp.Body.Close()

		if err1 != nil {
			println("err1:", err)
			continue
		}
		fmt.Printf("%s\n", data)
	}
}

// Exercise 1.7 ~ 1.9
func main() {

	// Loop to get URL
	for _, url := range os.Args[1:] {

		// Add header
		if !strings.HasPrefix(url, "h") && !strings.HasPrefix(url, "h") {
			url = "https://" + url
		}

		// Same effect
		//if url[0] != 'h' && url[0] != 'H' {
		//	url = "https://" + url
		//}

		resp, err := http.Get(url)

		defer resp.Body.Close()
		if err != nil {
			println("err:", err)
			continue
		}

		// Write to file
		//f, err1 := os.OpenFile("./Chapter1/target.html", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
		//defer f.Close()

		//if err1 != nil {
		//	println("err1-1:", err)
		//	continue
		//}
		//_, err1 = io.Copy(f, resp.Body)

		// Use io.Copy() to replace printf
		_, err1 := io.Copy(os.Stdout, resp.Body)

		if err1 != nil {
			println("err1:", err)
			continue
		}
		//fmt.Printf("\n\nRead %d bytes\n", n)
	}
}
