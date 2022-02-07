package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// echo "My first command line tool with Go" | ./Basic-Word-Counter
func main() {
	fmt.Println(count(os.Stdin))
}

func count(r io.Reader) int {

	// a scanner is used to read text from a Reader.
	scanner := bufio.NewScanner(r)

	// define the scanner split type to words (default is split by lines)
	scanner.Split(bufio.ScanWords)

	// define the counter
	wc := 0

	// for every word scanned, increment the counter
	for scanner.Scan() {
		wc++
	}

	// return total
	return wc
}
