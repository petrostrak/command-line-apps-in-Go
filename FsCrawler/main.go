package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

type config struct {
	ext  string // extention to filter out
	size int64  // minimum file size
	list bool   // list files
}

func main() {
	// parsing command line flags
	root := flag.String("root", ".", "Root directory to start")
	// action options
	list := flag.Bool("list", false, "List files only")
	// filter option
	ext := flag.String("ext", "", "File extention to filter out")
	size := flag.Int64("size", 0, "Minimum file size")
	flag.Parse()

	// create an instance of config struct
	c := config{
		*ext,
		*size,
		*list,
	}

	if err := run(*root, os.Stdout, c); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(root string, out io.Writer, cfg config) error {

}
