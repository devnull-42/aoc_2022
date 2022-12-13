package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	fs := bufio.NewScanner(file)
	fs.Split(bufio.ScanLines)

	valid := 0

	packets := []any{}
	pairGroup := 1

	for fs.Scan() {
		var packet any
		if fs.Text() == "" {

			if compare(packets[0], packets[1]) <= 0 {
				valid += pairGroup
			}
			packets = []any{}
			pairGroup++
		} else {
			json.Unmarshal([]byte(fs.Text()), &packet)
			packets = append(packets, packet)
		}
	}

	fmt.Printf("day13 part1: %d\n", valid)
}

func compare(first, second any) int {
	var firstList, secondList []any
	firstIsFloat, secondIsFloat := false, false

	// check first and return a list if it is a single float
	switch first.(type) {
	case float64:
		firstList, firstIsFloat = []any{first}, true
	case []any, []float64:
		firstList = first.([]any)
	}

	// check second and return a list if it is a single float
	switch second.(type) {
	case float64:
		secondList, secondIsFloat = []any{second}, true
	case []any, []float64:
		secondList = second.([]any)
	}

	// if both the first and second are lists with a single
	// float return the second subtracted from the first
	if firstIsFloat && secondIsFloat {
		return int(firstList[0].(float64) - secondList[0].(float64))
	}

	// if both first and second were not floats then recursivly call
	// compare again in the elements until one list runs out
	for i := 0; i < len(firstList) && i < len(secondList); i++ {
		if c := compare(firstList[i], secondList[i]); c != 0 {
			return c
		}
	}

	// default to returning the length of the second list subtracted
	// from the first
	return len(firstList) - len(secondList)
}
