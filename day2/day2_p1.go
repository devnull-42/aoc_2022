package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	file, _ := os.Open("input.txt")
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var score int

	for fileScanner.Scan() {
		pair := strings.Split(fileScanner.Text(), " ")

		switch {
		case pair[0] == "A" && pair[1] == "X":
			score += 4
		case pair[0] == "B" && pair[1] == "Y":
			score += 5
		case pair[0] == "C" && pair[1] == "Z":
			score += 6
		case pair[0] == "A" && pair[1] == "Y":
			score += 8
		case pair[0] == "B" && pair[1] == "Z":
			score += 9
		case pair[0] == "C" && pair[1] == "X":
			score += 7
		case pair[0] == "A" && pair[1] == "Z":
			score += 3
		case pair[0] == "B" && pair[1] == "X":
			score += 1
		case pair[0] == "C" && pair[1] == "Y":
			score += 2
		}
	}

	fmt.Printf("day2 part1: %d\n", score)
}
