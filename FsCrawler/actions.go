package main

import "os"

func delFile(path string) error {
	return os.Remove(path)
}
