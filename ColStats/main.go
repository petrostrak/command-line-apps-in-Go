package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
)

// cat << 'EOF' > testdata/example.csv
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
	// create the channel to receive results or errors of operations
	resCh := make(chan []float64)
	errCh := make(chan error)

	// Notice that you’re using an empty struct as the type for the doneCh channel.
	// This is a common pattern since this channel doesn’t need to send any data. It
	// only sends a signal indicating the processing is done. By using the empty
	// struct , the program doesn’t allocate any memory for this channel.
	doneCh := make(chan struct{})

	// This is the queue; you’ll add files to be processed to this channel, and
	// the worker goroutines will take them from this channel and process them.
	filesCh := make(chan string)

	wg := sync.WaitGroup{}

	// Loop through all files sending them through the channel
	// so each one will be processed when a worker is available
	go func() {
		defer close(filesCh)

		for _, fname := range filenames {
			filesCh <- fname
		}
	}()

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for fname := range filesCh {
				// open the file for reading
				f, err := os.Open(fname)
				if err != nil {
					errCh <- fmt.Errorf("cannot open file: %w", err)
					return
				}

				// parse the CSV into a slice of float64 numbers
				data, err := csvToFloat(f, column)
				if err != nil {
					errCh <- err
				}

				if err := f.Close(); err != nil {
					errCh <- err
				}

				resCh <- data

			}
		}()
	}

	go func() {
		wg.Wait()
		close(doneCh)
	}()

	for {
		select {
		case err := <-errCh:
			return err
		case data := <-resCh:
			consolidate = append(consolidate, data...)
		case <-doneCh:
			_, err := fmt.Fprintln(out, opFunc(consolidate))
			return err
		}
	}
}
