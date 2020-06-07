package main

import (
	"strings"
)

func climStairs(n int) (nums int) {
	if n == 1 {
		nums = 1
		return
	}
	if n == 2 {
		nums = 2
		return
	}
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, b+a
	}
	return b
}

func simplifyPath(path string) string {
	pathSlice := strings.Split(path, "/")
	cleanPath := make([]string, 0, len(pathSlice))
	for _, dirName := range pathSlice {
		if len(cleanPath) == 0 {
			if dirName == ".." || dirName == "." {
				continue
			}
			cleanPath = append(cleanPath, dirName)
		} else {
			if dirName == ".." {
				cleanPath = cleanPath[:len(cleanPath)]
			}
			if dirName == "." {
				continue
			}
		}
	}
	return strings.Join(cleanPath, "/")
}
