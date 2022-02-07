package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

// echo "My first command line tool with Go" | ./Basic-Word-Counter
func main() {

	// defining a boolean flag -l to count lines instead of words.
	lines := flag.Bool("l", false, "Count lines")

	// defining a boolean flag -b to count the total bytes read instead of words
	bytesRead := flag.Bool("b", false, "Count total bytes read")

	// parse the flags provided by the user.
	flag.Parse()

	// calling the count function to count the number of words (or lines)
	// received from the stdIn.
	fmt.Println(count(os.Stdin, *lines, *bytesRead))
}

func count(r io.Reader, countLines, countBytes bool) int {

	// a scanner is used to read text from a Reader.
	scanner := bufio.NewScanner(r)

	switch {
	case countLines:
		// if countLines flag is passed, we count lines
		scanner.Split(bufio.ScanLines)
	case countBytes:
		// if countBytes is passed, we count bytes
		scanner.Split(bufio.ScanBytes)
	default:
		// the default case is to count words
		scanner.Split(bufio.ScanWords)
	}

	// define the counter
	wc := 0

	// for every word scanned, increment the counter
	for scanner.Scan() {
		wc++
	}

	// return total
	return wc
}
