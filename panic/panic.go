package main

import (
	"os"
	"path/filepath"
)

func main() {
	panic("a problem")

	path := filepath.Join(os.TempDir(), "files")
	_, err := os.Create(path)
	if err != nil {
		panic(err)
	}
}
