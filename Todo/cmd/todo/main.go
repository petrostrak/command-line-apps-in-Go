package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	todo "github.com/petrostrak/command-line-apps-in-Go/Todo"
)

const (
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
	task := flag.String("task", "", "Task to be included in the todo list")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item to be completed")

	flag.Parse()

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

	case *task != "":
		task := todo.Item{*task, false, time.Now(), time.Time{}}
		// add the task
		l.Add(task)

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
