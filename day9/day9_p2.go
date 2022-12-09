package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type rope struct {
	knot    [][]int
	visited map[string]interface{}
}

func newRope(knots int) *rope {
	r := new(rope)
	r.knot = make([][]int, knots)
	for i := 0; i < len(r.knot); i++ {
		r.knot[i] = []int{0, 0}
	}
	r.visited = make(map[string]interface{})

	return r
}

func (r *rope) moveHead(direction string) {
	switch direction {
	case "U":
		r.knot[0][0]--
	case "D":
		r.knot[0][0]++
	case "L":
		r.knot[0][1]--
	case "R":
		r.knot[0][1]++
	}
	return
}

func (r *rope) moveTail(i int) {
	switch {
	case r.knot[i][0] == r.knot[i-1][0]-2:
		switch {
		case r.knot[i][1] == r.knot[i-1][1]:
			r.knot[i][0]++
		case r.knot[i][1] < r.knot[i-1][1]:
			r.knot[i][0]++
			r.knot[i][1]++
		case r.knot[i][1] > r.knot[i-1][1]:
			r.knot[i][0]++
			r.knot[i][1]--
		}
	case r.knot[i][0] == r.knot[i-1][0]+2:
		switch {
		case r.knot[i][1] == r.knot[i-1][1]:
			r.knot[i][0]--
		case r.knot[i][1] < r.knot[i-1][1]:
			r.knot[i][0]--
			r.knot[i][1]++
		case r.knot[i][1] > r.knot[i-1][1]:
			r.knot[i][0]--
			r.knot[i][1]--
		}
	case r.knot[i][1] == r.knot[i-1][1]-2:
		switch {
		case r.knot[i][0] == r.knot[i-1][0]:
			r.knot[i][1]++
		case r.knot[i][0] < r.knot[i-1][0]:
			r.knot[i][1]++
			r.knot[i][0]++
		case r.knot[i][0] > r.knot[i-1][0]:
			r.knot[i][1]++
			r.knot[i][0]--
		}
	case r.knot[i][1] == r.knot[i-1][1]+2:
		switch {
		case r.knot[i][0] == r.knot[i-1][0]:
			r.knot[i][1]--
		case r.knot[i][0] < r.knot[i-1][0]:
			r.knot[i][1]--
			r.knot[i][0]++
		case r.knot[i][0] > r.knot[i-1][0]:
			r.knot[i][1]--
			r.knot[i][0]--
		}
	}
}

func (r *rope) moveTails() {
	for i := 1; i < len(r.knot); i++ {
		r.moveTail(i)
	}
}

func (r *rope) recordTailPosition() {
	pos := fmt.Sprintf("%d,%d", r.knot[len(r.knot)-1][0], r.knot[len(r.knot)-1][1])
	r.visited[pos] = struct{}{}
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	// model rope
	r := newRope(10)

	for fileScanner.Scan() {
		instruction := strings.Split(fileScanner.Text(), " ")
		direction := instruction[0]
		steps, _ := strconv.Atoi(instruction[1])

		for i := 0; i < steps; i++ {
			r.moveHead(direction)
			r.moveTails()
			r.recordTailPosition()
		}
	}

	fmt.Printf("day 9 part2: %d\n", len(r.visited))
}
