package main

import (
	"bufio"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
)

type Turn int

const (
	TurnLeft Turn = iota
	TurnRight
)

const (
	startSuffix = "A"
	endSuffix   = "Z"
)

type Path = []Turn

type Node struct {
	name string
	next map[Turn]*Node
}

type Nodes map[string]*Node

func main() {
	slog.Info("got answer", "answer", run(os.Stdin))
}

func run(input io.Reader) int {
	path, nodes := parseInput(input)

	answer := 1
	for name, node := range nodes {
		if !strings.HasSuffix(name, startSuffix) {
			continue
		}

		steps := simulateSteps(node, path)
		answer = lcm(answer, steps)
	}

	return answer
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	for b != 0 {
		tmp := a % b
		a = b
		b = tmp
	}
	return a
}

func simulateSteps(startNode *Node, path []Turn) int {
	steps := 0
	for !strings.HasSuffix(startNode.name, endSuffix) {
		for _, turn := range path {
			startNode = startNode.next[turn]
			steps++
		}
	}

	return steps
}

func parseInput(input io.Reader) (path Path, nodes Nodes) {
	scanner := bufio.NewScanner(input)

	scanner.Scan()
	pathStr := scanner.Text()
	path = parsePath(pathStr)
	nodes = parseNodes(scanner)

	return
}

func parseNodes(scanner *bufio.Scanner) (nodes Nodes) {
	nodes = make(Nodes)
	createNodeIfMissing := func(name string) {
		if _, ok := nodes[name]; !ok {
			nodes[name] = new(Node)
			nodes[name].name = name
			nodes[name].next = make(map[Turn]*Node)
		}
	}

	for scanner.Scan() {
		nodeStr := scanner.Text()
		if len(nodeStr) == 0 {
			continue
		}

		name, l, r := parseNode(nodeStr)
		createNodeIfMissing(name)
		createNodeIfMissing(l)
		createNodeIfMissing(r)

		nodes[name].next[TurnLeft] = nodes[l]
		nodes[name].next[TurnRight] = nodes[r]
	}

	return
}

func parseNode(s string) (name string, l string, r string) {
	parts := strings.Split(s, " = (")
	name = parts[0]

	parts[1], _ = strings.CutSuffix(parts[1], ")")
	parts = strings.Split(parts[1], ", ")
	l = parts[0]
	r = parts[1]

	return
}

func parsePath(s string) (path Path) {
	for _, r := range s {
		path = append(path, parseTurn(r))
	}

	return
}

func parseTurn(r rune) Turn {
	switch r {
	case 'L':
		return TurnLeft
	case 'R':
		return TurnRight
	default:
		panic(fmt.Sprintf("invalid turn: %c", r))
	}
}
