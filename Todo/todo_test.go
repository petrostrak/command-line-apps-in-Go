package todo

import (
	"os"
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

// TestComplete tests the Complete method of the List type
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

// TestDelete tests the Delete method of the List type
func TestDelete(t *testing.T) {
	l := List{}

	tasks := []Item{
		{"New task 1", true, time.Now(), time.Now()},
		{"New task 2", false, time.Now(), time.Time{}},
		{"New task 3", false, time.Now(), time.Time{}},
	}

	for _, i := range tasks {
		l.Add(i)
	}

	if l[0].Task != tasks[0].Task {
		t.Errorf("expected %q, got %q", tasks[0].Task, l[0].Task)
	}

	l.Delete(2)

	if len(l) != 2 {
		t.Errorf("expected list length %d, got %d instead", 2, len(l))
	}

	if l[1].Task != tasks[2].Task {
		t.Errorf("expected %q, got %q instead", tasks[2].Task, l[1].Task)
	}
}

// TestSaveGet tests the Save and Get methods of the List type
func TestSaveGet(t *testing.T) {
	l1 := List{}
	l2 := List{}

	task := Item{"New task 1", true, time.Now(), time.Now()}
	l1.Add(task)

	if l1[0].Task != task.Task {
		t.Errorf("expected %q, got %q instead", task.Task, l1[0].Task)
	}

	tf, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatalf("error creating temp file: %s", err)
	}
	defer os.Remove(tf.Name())

	if err := l1.Save(tf.Name()); err != nil {
		t.Fatalf("error saving list to file: %s", err)
	}

	if err := l2.Get(tf.Name()); err != nil {
		t.Fatalf("error getting list from file: %s", err)
	}

	if l1[0].Task != l2[0].Task {
		t.Errorf("task %q should match %q task", l1[0].Task, l2[0].Task)
	}
}
