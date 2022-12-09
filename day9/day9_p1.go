package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type rope struct {
	head    []int
	tail    []int
	visited map[string]interface{}
}

func newRope() *rope {
	r := new(rope)
	r.head = []int{0, 0}
	r.tail = []int{0, 0}
	r.visited = make(map[string]interface{})

	return r
}

func (r *rope) moveHead(direction string) {
	switch direction {
	case "U":
		r.head[0]--
	case "D":
		r.head[0]++
	case "L":
		r.head[1]--
	case "R":
		r.head[1]++
	}
	return
}

func (r *rope) moveTail() {
	if r.tail[0] == r.head[0]-2 {
		if r.tail[1] == r.head[1] {
			r.tail[0]++
			return
		}
		if r.tail[1] == r.head[1]-1 {
			r.tail[0]++
			r.tail[1]++
			return
		}
		if r.tail[1] == r.head[1]+1 {
			r.tail[0]++
			r.tail[1]--
			return
		}
	}
	if r.tail[0] == r.head[0]+2 {
		if r.tail[1] == r.head[1] {
			r.tail[0]--
			return
		}
		if r.tail[1] == r.head[1]-1 {
			r.tail[0]--
			r.tail[1]++
			return
		}
		if r.tail[1] == r.head[1]+1 {
			r.tail[0]--
			r.tail[1]--
			return
		}
	}
	if r.tail[1] == r.head[1]-2 {
		if r.tail[0] == r.head[0] {
			r.tail[1]++
			return
		}
		if r.tail[0] == r.head[0]-1 {
			r.tail[1]++
			r.tail[0]++
			return
		}
		if r.tail[0] == r.head[0]+1 {
			r.tail[1]++
			r.tail[0]--
			return
		}
	}
	if r.tail[1] == r.head[1]+2 {
		if r.tail[0] == r.head[0] {
			r.tail[1]--
			return
		}
		if r.tail[0] == r.head[0]-1 {
			r.tail[1]--
			r.tail[0]++
			return
		}
		if r.tail[0] == r.head[0]+1 {
			r.tail[1]--
			r.tail[0]--
			return
		}
	}
}

func (r *rope) recordTailPosition() {
	pos := fmt.Sprintf("%d,%d", r.tail[0], r.tail[1])
	r.visited[pos] = struct{}{}
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	r := newRope()

	for fileScanner.Scan() {
		instruction := strings.Split(fileScanner.Text(), " ")
		direction := instruction[0]
		steps, _ := strconv.Atoi(instruction[1])
		// fmt.Printf("instruction: %v\n", instruction)

		for i := 0; i < steps; i++ {
			r.moveHead(direction)
			r.moveTail()
			r.recordTailPosition()
			// fmt.Printf("head: %v  tail: %v\n", r.head, r.tail)
		}
	}

	fmt.Printf("day 9 part1: %d\n", len(r.visited))
}
