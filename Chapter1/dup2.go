// Dup2 prints the count and text of lines that appears more than once
// in the input. It reads from stdin or from a list of named files
// Modified by using signal
// Testing:dup2.exe dup2-1 dup2-2

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// Init some variables

	counts := make(map[string]int)

	// Register the signal of Ctrl+C

	if len(os.Args) > 1 {
		//Read from files
		ReadFromFiles(os.Args[1:], counts)

	} else {
		//Read from stdin
		ReadFromStdin(os.Stdin, counts)
	}

}

func ReadFromStdin(file *os.File, counts map[string]int) {

	// Register the signal
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGINT)

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}

	go func() {
		// Block the goroutine
		<-sigCh

		// Output
		for line, n := range counts {
			if n > 1 {
				fmt.Println(line, ":", n)
			}
		}

		//Exit the program
		os.Exit(0)
	}()

	// Block the main routine
	// Wait for print
	select {}
}

func ReadFromFiles(list []string, counts map[string]int) {
	for _, fileStr := range list {
		file, err := os.Open(fileStr)

		if err != nil {
			fmt.Println("Open ", fileStr, "err:", err)
			continue
		}
		// Read the lines of files
		Countlines(file, counts)

		// Close the files
		file.Close()
	}

	// Summary
	for line, n := range counts {
		if n > 1 {
			fmt.Println(line, ":", n)
		}
	}
}

func Countlines(file *os.File, counts map[string]int) {
	input := bufio.NewScanner(file)

	for input.Scan() {
		counts[input.Text()]++
	}

	// Ignore the possible error
}
