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
	totalTime, recordDist := parseInput(input)

	waysToWin := 0
	for t := 0; t <= totalTime; t++ {
		dist := t * (totalTime - t)
		if dist > recordDist {
			waysToWin++
		}
	}

	return waysToWin
}

func parseInput(input io.Reader) (totalTime int, distance int) {
	scanner := bufio.NewScanner(input)

	scanner.Scan()
	line := scanner.Text()
	line, _ = strings.CutPrefix(line, "Time:")
	totalTime = parseNumber(line)

	scanner.Scan()
	line = scanner.Text()
	line, _ = strings.CutPrefix(line, "Distance:")
	distance = parseNumber(line)

	return
}

func parseNumber(s string) int {
	parts := strings.Split(s, " ")
	numStr := strings.Join(parts, "")
	res, _ := strconv.ParseInt(numStr, 10, 64)
	return int(res)
}
