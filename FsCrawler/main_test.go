package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestRun(t *testing.T) {

	var (
		logBuffer bytes.Buffer
	)

	testCases := []struct {
		name     string
		root     string
		cfg      config
		expected string
	}{
		{"NoFilter", "testdata", config{"", 0, true, false, &logBuffer}, "testdata/dir.log\ntestdata/dir2/script.sh\n"},
		{"FilterExtensionMatch", "testdata", config{".log", 0, true, false, &logBuffer}, "testdata/dir.log\n"},
		{"FilterExtensionSizeMatch", "testdata", config{".log", 10, true, false, &logBuffer}, "testdata/dir.log\n"},
		{"FilterExtensionSizeNoMatch", "testdata", config{".log", 20, true, false, &logBuffer}, ""},
		{"FilterExtensionNoMatch", "testdata", config{".gz", 0, true, false, &logBuffer}, ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var buffer bytes.Buffer

			if err := run(tc.root, &buffer, tc.cfg); err != nil {
				t.Fatal(err)
			}

			res := buffer.String()

			if tc.expected != res {
				t.Errorf("expected %q, got %q instead\n", tc.expected, res)
			}
		})
	}
}

func createTempDir(t *testing.T, files map[string]int) (dirname string, cleanup func()) {
	t.Helper()

	tempDir, err := ioutil.TempDir("", "walktest")
	if err != nil {
		t.Fatal(err)
	}

	for k, n := range files {
		for j := 1; j <= n; j++ {
			fname := fmt.Sprintf("file%d%s", j, k)
			fpath := filepath.Join(tempDir, fname)
			if err := ioutil.WriteFile(fpath, []byte("dummy"), 0644); err != nil {
				t.Fatal(err)
			}
		}
	}

	return tempDir, func() { os.RemoveAll(tempDir) }
}

func TestRunDelExtension(t *testing.T) {

	var (
		buffer    bytes.Buffer
		logBuffer bytes.Buffer
	)

	testCases := []struct {
		name        string
		cfg         config
		extNoDelete string
		nDelete     int
		nNoDelete   int
		expected    string
	}{
		{"DeleteExtensionNoMatch", config{".log", 0, false, true, &logBuffer}, ".gz", 0, 10, ""},
		{"DeleteExtensionMatch", config{".log", 0, false, true, &logBuffer}, ".gz", 10, 0, ""},
		{"DeleteExtensionMixed", config{".log", 0, false, true, &logBuffer}, ".gz", 5, 5, ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			tempDir, cleanup := createTempDir(t, map[string]int{
				tc.cfg.ext:     tc.nDelete,
				tc.extNoDelete: tc.nNoDelete,
			})
			defer cleanup()

			if err := run(tempDir, &buffer, tc.cfg); err != nil {
				t.Fatal(err)
			}

			res := buffer.String()

			if tc.expected != res {
				t.Errorf("expected %q, got %q instead\n", tc.expected, res)
			}

			filesLeft, err := ioutil.ReadDir(tempDir)
			if err != nil {
				t.Error(err)
			}

			if len(filesLeft) != tc.nNoDelete {
				t.Errorf("expected %d files left, got %d instead\n", tc.nNoDelete, len(filesLeft))
			}

			expLogLines := tc.nDelete + 1
			lines := bytes.Split(logBuffer.Bytes(), []byte("\n"))
			if len(lines) != expLogLines {
				t.Errorf("Expected %d log lines, got %d instead\n",
					expLogLines, len(lines))
			}
		})
	}
}
