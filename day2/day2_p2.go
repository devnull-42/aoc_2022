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
			score += 3
		case pair[0] == "B" && pair[1] == "X":
			score += 1
		case pair[0] == "C" && pair[1] == "X":
			score += 2
		case pair[0] == "A" && pair[1] == "Y":
			score += 4
		case pair[0] == "B" && pair[1] == "Y":
			score += 5
		case pair[0] == "C" && pair[1] == "Y":
			score += 6
		case pair[0] == "A" && pair[1] == "Z":
			score += 8
		case pair[0] == "B" && pair[1] == "Z":
			score += 9
		case pair[0] == "C" && pair[1] == "Z":
			score += 7
		}
	}

	fmt.Printf("day2 part2: %d\n", score)
}
