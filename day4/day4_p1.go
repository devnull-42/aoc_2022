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

	var overlapCount int

	for fileScanner.Scan() {
		sections := strings.Split(fileScanner.Text(), ",")
		section1 := strings.Split(sections[0], "-")
		section2 := strings.Split(sections[1], "-")

		section1Start, _ := strconv.Atoi(section1[0])
		section1End, _ := strconv.Atoi(section1[1])
		section2Start, _ := strconv.Atoi(section2[0])
		section2End, _ := strconv.Atoi(section2[1])

		if (section1Start <= section2Start && section1End >= section2End) ||
			(section2Start <= section1Start && section2End >= section1End) {
			overlapCount++
		}
	}

	fmt.Printf("day 4 part1: %d\n", overlapCount)
}
