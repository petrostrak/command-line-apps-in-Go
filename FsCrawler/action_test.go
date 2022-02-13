package main

import (
	"os"
	"testing"
)

func TestFilterOut(t *testing.T) {
	testCases := []struct {
		name     string
		file     string
		ext      string
		minSize  int64
		expected bool
	}{
		{"FilterNoExtention", "testdata/dir.log", "", 0, false},
		{"FilterExtentionMatch", "testdata/dir.log", ".log", 0, false},
		{"FilterExtentionNoMatch", "testdata/dir.log", ".sh", 0, true},
		{"FilterExtentionSizeMatch", "testdata/dir.log", ".log", 10, false},
		{"FilterExtentionSizeNoMatch", "testdata/dir.log", ".log", 20, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			// retrieve the file's attribute using os.Stat
			info, err := os.Stat(tc.file)
			if err != nil {
				t.Fatal(err)
			}

			f := filterOut(tc.file, tc.ext, tc.minSize, info)

			if f != tc.expected {
				t.Errorf("expected '%t', got '%t' instead\n", tc.expected, f)
			}
		})
	}
}
