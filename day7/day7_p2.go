package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	dirSizes := make([]int, 0)
	sortedDirs := make([]int, 0)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if line == "$ ls" {
			continue
		}

		if strings.HasPrefix(line, "$ cd") && !strings.HasSuffix(line, "..") {
			dirSizes = append(dirSizes, 0)
		}

		if !strings.HasPrefix(line, "dir") && !strings.HasPrefix(line, "$ cd") {
			file := strings.Split(line, " ")
			size, _ := strconv.Atoi(file[0])
			for i := range dirSizes {
				dirSizes[i] += size
			}
		}

		if line == "$ cd .." {
			sortedDirs = append(sortedDirs, dirSizes[len(dirSizes)-1])
			dirSizes = dirSizes[:len(dirSizes)-1]
		}

	}

	for i := range dirSizes {
		sortedDirs = append(sortedDirs, dirSizes[i])
	}

	sort.Sort(sort.IntSlice(sortedDirs))
	const totalSpace = 70000000
	const neededSpace = 30000000

	goal := neededSpace - (totalSpace - sortedDirs[len(sortedDirs)-1])

	var dirSize int

	for i := 1; i < len(sortedDirs); i++ {
		if sortedDirs[i] >= goal {
			dirSize = sortedDirs[i]
			break
		}
	}

	fmt.Printf("day7 part1: %v\n", dirSize)

}
