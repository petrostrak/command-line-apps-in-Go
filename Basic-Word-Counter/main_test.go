package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word 3word4\n")

	expected := 4
	response := count(b)

	if response != expected {
		t.Errorf("Expected %d; got %d instead.", expected, response)
	}
}
