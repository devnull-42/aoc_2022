package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var total int
	dirSizes := make([]int, 0)

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
			if dirSizes[len(dirSizes)-1] <= 100000 {
				total += dirSizes[len(dirSizes)-1]
			}
			dirSizes = dirSizes[:len(dirSizes)-1]
		}

	}

	for i := range dirSizes {
		if dirSizes[i] <= 100000 {
			total += dirSizes[i]
		}
	}

	fmt.Printf("day7 part1: %d\n", total)

}
