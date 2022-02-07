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

	// defineing a boolean flag -l to count lines instead of words.
	lines := flag.Bool("l", false, "Count lines")

	// parse the flags provided by the user.
	flag.Parse()

	// calling the count function to count the number of words (or lines)
	// received from the stdIn.
	fmt.Println(count(os.Stdin, *lines))
}

func count(r io.Reader, countLines bool) int {

	// a scanner is used to read text from a Reader.
	scanner := bufio.NewScanner(r)

	// if the count lines flag is not set, we want to count words
	if !countLines {
		// define the scanner split type to words (default is split by lines)
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
