package todo

import (
	"testing"
	"time"
)

// TestAdd tests the Add method of the List type
func TestAdd(t *testing.T) {
	l := List{}

	taskName := Item{
		Task:        "New Task",
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	l.Add(taskName)

	if l[0].Task != taskName.Task {
		t.Errorf("Expected %q, got %q instead", taskName.Task, l[0].Task)
	}
}

func TestComplete(t *testing.T) {
	l := List{}

	taskName := Item{
		Task:        "New Task",
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	l.Add(taskName)

	if l[0].Task != taskName.Task {
		t.Errorf("expected %q, got %q instead", taskName.Task, l[0].Task)
	}

	if l[0].Done {
		t.Errorf("New task should not be completed")
	}

	l.Complete(1)

	if !l[0].Done {
		t.Errorf("New task should be completed")
	}
}
