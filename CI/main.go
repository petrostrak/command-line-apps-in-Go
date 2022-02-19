package main

import (
	"fmt"
	"io"
	"os/exec"
)

func run(proj string, out io.Writer) error {

	// check if the project dir is provided
	if proj == "" {
		return fmt.Errorf("project dir is required")
	}

	args := []string{"build", ".", "errors"}

	cmd := exec.Command("go", args...)
	cmd.Dir = proj

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("'go build' failed: %s", err)
	}

	_, err := fmt.Fprintln(out, "Go Build: SUCCESS")

	return err
}
