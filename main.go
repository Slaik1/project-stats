package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Pair struct {
	Key   string
	Value int
}

var path = os.Args[1]

var linesCount = 0
var dirCount = 0
var filesCount = 0
var counter = make(map[string]int)

func cycle(basepath string) {
	if IsDir(basepath) {
		dirCount += 1

		files, err := GetDirFiles(basepath)

		if err != nil {
			return
		}

		for _, strPath := range files {
			cycle(filepath.Join(basepath, strPath.Name()))
		}
		return
	}

	filesCount += 1
	linesCount += GetFileLinesCount(basepath)
	stat, err := GetFileStat(basepath)

	if err != nil {
		return
	}

	fileNameArr := strings.Split(stat.Name(), ".")

	var fileExtension = fileNameArr[len(fileNameArr)-1]
	counter[fileExtension]++
}

func main() {
	if path == "--help" {
		fmt.Println("run example: fileInfo  ./pathToFolder")
		return
	}

	cycle(path)

	fmt.Println("-------------Total Amount-------------")
	fmt.Println("linesCount: ", linesCount)
	fmt.Println("dirCount: ", dirCount)
	fmt.Println("filesCount: ", filesCount)
	fmt.Println()
	fmt.Println("-------------Files Statistic-------------")

	pairs := make([]Pair, 0, len(counter))
	for key, value := range counter {
		pairs = append(pairs, Pair{key, value})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value > pairs[j].Value
	})

	for _, pair := range pairs {
		fmt.Printf("%s: %d\n", pair.Key, pair.Value)
	}
}
