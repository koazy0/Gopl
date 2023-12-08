// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count
// Modified by using signal
package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//create a new map
	counts := make(map[string]int)

	// Create a signal channel
	sigCh := make(chan os.Signal, 1)

	// Register the signal of Ctrl+C
	signal.Notify(sigCh, os.Interrupt, syscall.SIGINT)

	// Read the stdin
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
