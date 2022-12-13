package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	graph, startPoints, endNode := getGraph()

	shortestPath := 1000

	for _, start := range startPoints {
		maxY := len(graph)
		maxX := len(graph[0])
		mpq := new(minPQ)
		visitedNodes := make(map[string]*node)

		values := getMapValues(maxY, maxX, graph)

		mpq.nodes = append(mpq.nodes, newNode(point{y: start[0], x: start[1]}, 0, nil))

		for len(mpq.nodes) > 0 {
			da(mpq, visitedNodes, maxY, maxX, values)
		}
		if visitedNodes[endNode].totalCost < shortestPath {
			shortestPath = visitedNodes[endNode].totalCost
		}
	}
	// fmt.Println(len(visitedNodes))
	// printPath(endNode, visitedNodes, values)
	fmt.Printf("endNode: %s\n", endNode)
	fmt.Printf("day12 part2: %d\n", shortestPath)
}

func getGraph() ([][]rune, [][]int, string) {
	file, _ := os.Open("input.txt")
	defer file.Close()

	fs := bufio.NewScanner(file)
	fs.Split(bufio.ScanLines)

	graph := make([][]rune, 0)
	startPoints := make([][]int, 0)
	var endNode string

	i := 0
	for fs.Scan() {
		line := strings.Split(fs.Text(), "")
		graphRow := make([]rune, 0)

		for j, v := range line {
			switch v {
			case "S":
				startPoints = append(startPoints, []int{i, j})
				graphRow = append(graphRow, 'a')
			case "E":
				endNode = fmt.Sprintf("%d,%d", i, j)
				graphRow = append(graphRow, 'z')
			default:
				if v == "a" {
					startPoints = append(startPoints, []int{i, j})
				}
				graphRow = append(graphRow, rune(v[0]))
			}
		}
		graph = append(graph, graphRow)
		i++
	}
	return graph, startPoints, endNode
}

type node struct {
	point
	totalCost    int
	previousNode *node
}

func newNode(p point, tc int, previousNode *node) *node {
	n := new(node)
	n.point = p
	n.totalCost = tc
	n.previousNode = previousNode
	return n
}

func getNeighbors(n *node, maxY, maxX int, nodeValues map[point]rune, vistedNodes map[string]*node) []*node {
	pSlice := make([]*node, 0)
	// get up
	if n.point.y > 0 {
		p := point{y: n.point.y - 1, x: n.point.x}
		if _, exists := vistedNodes[p.ToString()]; !exists {
			if nodeValues[p]-nodeValues[n.point] < 2 {
				pSlice = append(pSlice, newNode(p, n.totalCost+1, n))
			} else {
				pSlice = append(pSlice, newNode(p, 1000, n))
			}
		}
	}
	// get down
	if n.point.y < maxY-1 {
		p := point{y: n.point.y + 1, x: n.point.x}
		if _, exists := vistedNodes[p.ToString()]; !exists {
			if nodeValues[p]-nodeValues[n.point] < 2 {
				pSlice = append(pSlice, newNode(p, n.totalCost+1, n))
			} else {
				pSlice = append(pSlice, newNode(p, 1000, n))
			}
		}
	}
	// get left
	if n.point.x > 0 {
		p := point{y: n.point.y, x: n.point.x - 1}
		if _, exists := vistedNodes[p.ToString()]; !exists {
			if nodeValues[p]-nodeValues[n.point] < 2 {
				pSlice = append(pSlice, newNode(p, n.totalCost+1, n))
			} else {
				pSlice = append(pSlice, newNode(p, 1000, n))
			}
		}
	}
	// get right
	if n.point.x < maxX-1 {
		p := point{y: n.point.y, x: n.point.x + 1}
		if _, exists := vistedNodes[p.ToString()]; !exists {
			if nodeValues[p]-nodeValues[n.point] < 2 {
				pSlice = append(pSlice, newNode(p, n.totalCost+1, n))
			} else {
				pSlice = append(pSlice, newNode(p, 1000, n))
			}
		}
	}
	return pSlice
}

type minPQ struct {
	nodes []*node
}

func (m minPQ) Print() {
	fmt.Println(len(m.nodes))
	for _, n := range m.nodes {
		fmt.Printf("%+v\n", *n)
	}
	fmt.Println()
}

func (m minPQ) NodeExists(n *node) (int, bool) {
	for i, node := range m.nodes {
		if node.point.x == n.point.x && node.point.y == n.point.y {
			return i, true
		}
	}
	return 0, false
}

func (m *minPQ) Push(n *node) {
	i, exists := m.NodeExists(n)
	if !exists {
		m.nodes = append(m.nodes, n)
	} else {
		if n.totalCost < m.nodes[i].totalCost {
			m.nodes[i].totalCost = n.totalCost
			m.nodes[i].previousNode = n.previousNode
		}
	}
}

func (m *minPQ) Pop() *node {
	n := m.nodes[0]
	m.nodes = m.nodes[1:]
	return n
}

func (m *minPQ) Sort() {
	sort.Slice(m.nodes, func(i, j int) bool {
		return m.nodes[i].totalCost < m.nodes[j].totalCost
	})
}

type point struct {
	y, x int
}

func (p point) ToString() string {
	return fmt.Sprintf("%d,%d", p.y, p.x)
}

func getMapValues(maxY, maxX int, graph [][]rune) map[point]rune {
	nodes := make(map[point]rune)
	for i := 0; i < maxY; i++ {
		for j := 0; j < maxX; j++ {
			nodes[point{y: i, x: j}] = graph[i][j]
		}
	}
	return nodes
}

func da(mpq *minPQ, visitedNodes map[string]*node, maxY, maxX int, nodeValues map[point]rune) {
	// pop from the mpq
	n := mpq.Pop()

	// add to visited nodes
	visitedNodes[n.point.ToString()] = n

	// push neighbors into the mpq
	neighbors := getNeighbors(n, maxY, maxX, nodeValues, visitedNodes)
	for _, neighbor := range neighbors {
		mpq.Push(neighbor)
	}

	// sort mpq
	mpq.Sort()
}

func printPath(end string, visitedNodes map[string]*node, values map[point]rune) {
	node := visitedNodes[end]

	for node.point.ToString() != "0,0" {
		fmt.Printf("%s -> %d -> %d -> %d -> %s\n", node.point.ToString(), node.totalCost, values[node.point], node.previousNode.totalCost, node.previousNode.point.ToString())
		node = visitedNodes[node.previousNode.point.ToString()]
	}
}
