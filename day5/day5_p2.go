package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	file, _ := os.Open("input.txt")
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	crates := make([][]string, 9)

	for fileScanner.Text() != " 1   2   3   4   5   6   7   8   9 " {
		for i, r := range fileScanner.Text() {
			if r != ' ' && r != '[' && r != ']' {
				crates[i/4] = append([]string{string(r)}, crates[i/4]...)
			}
		}
		fileScanner.Scan()
	}

	// read black line
	fileScanner.Scan()

	instructionRE := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

	for fileScanner.Scan() {
		matches := instructionRE.FindStringSubmatch(fileScanner.Text())

		amount, _ := strconv.Atoi(matches[1])
		fromCrate, _ := strconv.Atoi(matches[2])
		toCrate, _ := strconv.Atoi(matches[3])

		fromCrate--
		toCrate--

		movedCrates := crates[fromCrate][len(crates[fromCrate])-amount:]
		crates[toCrate] = append(crates[toCrate], movedCrates...)
		crates[fromCrate] = crates[fromCrate][:len(crates[fromCrate])-amount]
	}

	result := make([]string, 0)

	for _, stack := range crates {
		result = append(result, stack[len(stack)-1])
	}

	fmt.Printf("day5 part2: %v\n", strings.Join(result, ""))
}
