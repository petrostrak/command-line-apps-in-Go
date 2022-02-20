package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

// ./CI -p testdata/tool
func main() {
	proj := flag.String("p", "", "Project directory")
	flag.Parse()

	if err := run(*proj, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(proj string, out io.Writer) error {

	// check if the project dir is provided
	if proj == "" {
		return fmt.Errorf("project dir is required: %w", ErrValidation)
	}

	pipeline := make([]step, 2)
	pipeline[0] = newStep(
		"go build",
		"go",
		"Go Build: SUCCESS",
		proj,
		[]string{"build", ".", "errors"},
	)

	pipeline[1] = newStep(
		"go test",
		"go",
		"Go TEST: SUCCESS",
		proj,
		[]string{"test", "-v"},
	)

	for _, s := range pipeline {
		msg, err := s.execute()
		if err != nil {
			return err
		}

		_, err = fmt.Fprintln(out, msg)
		if err != nil {
			return err
		}
	}

	return nil
}
