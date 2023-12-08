package main

import (
	"fmt"
	"os"
)

func main() {
	for _, str := range os.Args[1:] {
		fmt.Printf("%s ", str)
	}
	fmt.Println()
}

func main01() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
