package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word 3word4\n")

	expected := 4
	response := count(b, false)

	if response != expected {
		t.Errorf("Expected %d; got %d instead.", expected, response)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\nline2\nline3 word1")

	expected := 3
	response := count(b, true)

	if response != expected {
		t.Errorf("Expected %d; got %d instead.", expected, response)
	}
}
