package main

import (
	"bufio"
	"io"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

func main() {
	slog.Info("got answer", "answer", run(os.Stdin))
}

func run(input io.Reader) (answer int) {
	times, distances := parseInput(input)

	answer = 1
	for i, totalTime := range times {
		recordDist := distances[i]

		waysToWin := 0
		for t := 0; t <= totalTime; t++ {
			dist := t * (totalTime - t)
			if dist > recordDist {
				waysToWin++
			}
		}

		answer *= waysToWin
	}

	return
}

func parseInput(input io.Reader) (times []int, distances []int) {
	scanner := bufio.NewScanner(input)

	scanner.Scan()
	line := scanner.Text()
	line, _ = strings.CutPrefix(line, "Time:")
	times = parseNumbers(line)

	scanner.Scan()
	line = scanner.Text()
	line, _ = strings.CutPrefix(line, "Distance:")
	distances = parseNumbers(line)

	return
}

func parseNumbers(s string) (result []int) {
	parts := strings.Split(s, " ")

	for _, part := range parts {
		if len(part) == 0 {
			continue
		}

		num, _ := strconv.ParseInt(part, 10, 64)
		result = append(result, int(num))
	}

	return
}
