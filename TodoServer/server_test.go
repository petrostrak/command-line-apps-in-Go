package main

import (
	"net/http/httptest"
	"testing"
)

func setupAPI(t *testing.T) (string, func()) {
	t.Helper()

	ts := httptest.NewServer(newMux(""))

	return ts.URL, func() {
		ts.Close()
	}
}
