package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

type config struct {
	ext  string    // extention to filter out
	size int64     // minimum file size
	list bool      // list files
	del  bool      // delete files
	wLog io.Writer // log destination writer
}

var (
	f   = os.Stdout
	err error
)

func main() {
	// parsing command line flags
	root := flag.String("root", ".", "Root directory to start")
	logFile := flag.String("log", "", "Log deletes to the file")
	// action options
	list := flag.Bool("list", false, "List files only")
	del := flag.Bool("del", false, "Delete files")
	// filter option
	ext := flag.String("ext", "", "File extention to filter out")
	size := flag.Int64("size", 0, "Minimum file size")
	flag.Parse()

	if *logFile != "" {
		f, err = os.OpenFile(*logFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer f.Close()
	}

	// create an instance of config struct
	c := config{
		*ext,
		*size,
		*list,
		*del,
		f,
	}

	if err := run(*root, os.Stdout, c); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(root string, out io.Writer, cfg config) error {
	delLogger := log.New(cfg.wLog, "DELETED FILE: ", log.LstdFlags)

	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if filterOut(path, cfg.ext, cfg.size, info) {
			return nil
		}

		// if list was explicitly set, don't do anything else
		if cfg.list {
			return listFile(path, out)
		}

		// delete files
		if cfg.del {
			return delFile(path, delLogger)
		}

		// list is the default option if nothing else was set
		return listFile(path, out)
	})
}

func filterOut(path, ext string, minSize int64, info os.FileInfo) bool {

	// if there is a directory and is within the given size limits
	if info.IsDir() || info.Size() < minSize {
		return true
	}

	// if given extention is not met
	if ext != "" && filepath.Ext(path) != ext {
		return true
	}

	return false
}

// prints out the path of the current file to the specified io.Writer
// returning any potential error from the operation.
func listFile(path string, out io.Writer) error {
	_, err := fmt.Fprintln(out, path)
	return err
}
