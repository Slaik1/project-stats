package main

import (
	"os"
)

func IsDir(path string) bool {
	info, err := os.Stat(path)

	if err != nil {
		return false
	}

	return info.IsDir()
}

func GetFileStat(path string) (os.FileInfo, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func GetDirFiles(path string) ([]os.DirEntry, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	return entries, nil
}

func GetFileLinesCount(path string) int {
	data, err := os.ReadFile(path)
	if err != nil {
		return 0
	}

	counter := 0
	for _, b := range data {
		if b == '\n' {
			counter++
		}
	}

	if len(data) > 0 && data[len(data)-1] != '\n' {
		counter++
	}

	return counter
}
