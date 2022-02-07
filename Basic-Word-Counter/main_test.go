package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word 3word4\n")

	expected := 4
	response := count(b, false, false)

	if response != expected {
		t.Errorf("Expected %d; got %d instead.", expected, response)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\nline2\nline3 word1")

	expected := 3
	response := count(b, true, false)

	if response != expected {
		t.Errorf("Expected %d; got %d instead.", expected, response)
	}
}

func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("My first command line tool with Go\n")

	expected := 35
	response := count(b, false, true)

	if response != expected {
		t.Errorf("Expected %d; got %d instead.", expected, response)
	}
}
