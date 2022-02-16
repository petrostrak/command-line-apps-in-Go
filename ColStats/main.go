package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// verify and parse arguments
	op := flag.String("op", "sum", "Operation to be executed")
	column := flag.Int("col", 1, "CSV column on which to execute operation")
	flag.Parse()

	if err := run(flag.Args(), *op, *column, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// the run function accepts four input parameters
//
// filenames represent the file names to process
// op represents the operation to execute, such as sum or avg
// column represents the column on which to execute the operaiton
// out of type io.Writer interface to print out the results
func run(filenames []string, op string, column int, out io.Writer) error {

	// create a var of type statsFunc
	var opFunc statsFunc

	// validate the user-provided parameters
	if len(filenames) == 0 {
		return ErrNoFiles
	}

	// check the column parameter
	if column < 1 {
		return fmt.Errorf("%w: %d", ErrInvalidColumn, column)
	}

	// validate the operation and define the opFunc accordingly
	switch op {
	case "sum":
		opFunc = sum
	case "avg":
		opFunc = avg
	default:
		return fmt.Errorf("%w: %s", ErrInvalidOperation, op)
	}

	consolidate := make([]float64, 0)

	// loop through all files adding their data to consolidate
	for _, fname := range filenames {

		// open the file for reading
		f, err := os.Open(fname)
		if err != nil {
			return fmt.Errorf("cannot open file: %w", err)
		}

		// parse the CSV into a slice of float64 numbers
		data, err := csvToFloat(f, column)
		if err != nil {
			return err
		}

		if err := f.Close(); err != nil {
			return err
		}

		// append the data to consolidate
		consolidate = append(consolidate, data...)
	}

	_, err := fmt.Fprintln(out, opFunc(consolidate))

	return err
}
