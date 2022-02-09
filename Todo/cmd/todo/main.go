package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	todo "github.com/petrostrak/command-line-apps-in-Go/Todo"
)

const (
	// hardcoding the filename
	todoFileName = ".todo.json"
)

func main() {

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
	case len(os.Args) == 1:
		// list current to do items
		for _, item := range *l {
			fmt.Println(item.Task)
		}

		// concatenate all provided arguments with a space and
		// add to the list as an item
	default:
		// concatenate all arguments with a space
		task := strings.Join(os.Args[1:], " ")

		// add the task
		item := todo.Item{task, false, time.Now(), time.Time{}}
		l.Add(item)

		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

}
