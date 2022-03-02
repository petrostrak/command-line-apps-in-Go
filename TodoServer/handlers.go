package main

import (
	"errors"
	"net/http"
	"sync"

	todo "github.com/petrostrak/command-line-apps-in-Go/Todo"
)

var (
	ErrNotFound    = errors.New("not found")
	ErrInvalidData = errors.New("invalid data")
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	content := "There's an API here"
	replyTextContent(w, r, http.StatusOK, content)
}

func replyTextContent(w http.ResponseWriter, r *http.Request, status int, content string) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(status)
	w.Write([]byte(content))
}

func todoRouter(todoFile string, l sync.Locker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		list := &todo.List{}

		// we use the l.Lock() from the sync.Locker interface to lock
		// the entire request handling. This prevents concurrent access
		// to the file represented by the variable todoFile which could
		// lead to data loss.
		l.Lock()
		defer l.Unlock()
		if err := list.Get(todoFile); err != nil {
			replyError(w, r, http.StatusInternalServerError, err.Error())
			return
		}
	}
}
