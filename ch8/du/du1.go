package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func walkDir(dir string, fileSize chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subDir := filepath.Join(dir, entry.Name())
			walkDir(subDir, fileSize)
		} else {
			fileSize <- entry.Size()
		}
	}

}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1:%v\n", err)
		return nil
	}
	return entries
}
