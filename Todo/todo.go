package todo

import "time"

// item struct represents a ToDo item
type Item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

// list represents a list of DoTo items
type List []Item

// Add creates a new todo item and appends it to the list
func (l *List) Add(task Item) {
	t := Item{
		Task:        task.Task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*l = append(*l, t)
}
