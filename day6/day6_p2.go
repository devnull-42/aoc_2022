package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	fileScanner.Scan()

	var result int

	for i := 0; i < len(fileScanner.Text())-14; i++ {
		if !signal(fileScanner.Text()[i : i+14]) {
			result = i + 14
			break
		}
	}

	fmt.Printf("day6 part1: %v\n", result)
}

func signal(packet string) bool {
	if len(packet) == 1 {
		return false
	}
	for i := 1; i < len(packet); i++ {
		if packet[0] == packet[i] {
			return true
		}
	}
	return signal(packet[1:])
}
