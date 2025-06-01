package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var path = os.Args[1]

var linesCount = 0
var dirCount = 0
var filesCount = 0

func cycle(basepath string) {
	if isDir(basepath) {
		dirCount += 1

		files, err := getDirFiles(basepath)

		if err != nil {
			return
		}

		for _, strPath := range files {
			cycle(filepath.Join(basepath, strPath.Name()))
		}
		return
	}

	filesCount += 1
	linesCount += getFileLinesCount(basepath)
}

func main() {
	if path == "--help" {
		fmt.Println("run example: fileInfo  ./pathToFolder")
	}

	cycle(path)

	fmt.Println("-------------Total Amount-------------")
	fmt.Println("linesCount: ", linesCount)
	fmt.Println("dirCount: ", dirCount)
	fmt.Println("filesCount: ", filesCount)
}
