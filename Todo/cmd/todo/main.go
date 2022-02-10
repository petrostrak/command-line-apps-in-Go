package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	todo "github.com/petrostrak/command-line-apps-in-Go/Todo"
)

var (
	// hardcoding the filename
	todoFileName = ".todo.json"
)

// > ./todo -h
// Usage of ./todo:
//   -complete int
//     	Item to be completed
//   -list
//     	List all tasks
//   -task string
//     	Task to be included in the todo list
func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(),
			"%s tool. Developed for the Pragmatic Bookshelf\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Copyright 2022\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Usage information:")
		flag.PrintDefaults()
	}

	// parsing command line flags
	add := flag.Bool("add", false, "Add task to the todo list")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item to be completed")

	flag.Parse()

	// Check if the user defined the ENV VAR for a custom file name
	// export TODO_FILENAME=new-todo.json
	if os.Getenv("TODO_FILENAME") != "" {
		todoFileName = os.Getenv("TODO_FILENAME")
	}

	// define an item list
	l := &todo.List{}

	// use the Get method to read to do items from file
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// decide what to do based on the number of arguments provided
	switch {
	// for an extra argument, print the list
	case *list:
		// list current to do items
		fmt.Print(l)
	case *complete > 0:
		// complete the given item
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	case *add:
		// when any arguments (excluding flags) are provided, they will be
		// used as the new task
		t, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		task := todo.Item{t, false, time.Now(), time.Time{}}
		// add the task
		l.Add(task)

		// save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	default:
		// invalid flag provided
		fmt.Fprintln(os.Stderr, "invalid option")
		os.Exit(1)

	}

}

// getTask function decides where to get the description for a new
// task from: arguments or STDIN
func getTask(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	s := bufio.NewScanner(r)
	s.Scan()
	if err := s.Err(); err != nil {
		return "", err
	}

	if len(s.Text()) == 0 {
		return "", errors.New("task cannot be blank")
	}

	return s.Text(), nil
}
