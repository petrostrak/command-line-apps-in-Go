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
