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

		for i := 0; i < amount; i++ {
			var val string
			val, crates[fromCrate] = pop(crates[fromCrate])
			crates[toCrate] = append(crates[toCrate], val)
		}
	}

	result := make([]string, 0)

	for _, stack := range crates {
		result = append(result, stack[len(stack)-1])
	}

	fmt.Printf("day5 part1: %v\n", strings.Join(result, ""))
}

func pop(strSlice []string) (string, []string) {
	val := (strSlice)[len(strSlice)-1]
	strSlice = (strSlice)[:len(strSlice)-1]
	return val, strSlice
}
