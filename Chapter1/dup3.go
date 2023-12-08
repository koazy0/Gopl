// Only read from files
// Testing:dup3.exe dup2-1 dup2-2

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)

	// From index 1
	for _, file := range os.Args[1:] {

		// os.ReadFile() return the []byte
		data, err := os.ReadFile(file)

		if err != nil {
			//error
		}
		println(string(data))
		// Slice an array
		// Windows: \r\n
		// Unix: \n
		for _, line := range strings.Split(string(data), "\r\n") {
			counts[line]++
			//println(line)
		}

	}
	for line, n := range counts {

		if n > 1 {
			fmt.Printf("%s:%d\n", line, n)
		}
	}
}
