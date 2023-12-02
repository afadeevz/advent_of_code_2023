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
	answer := run(os.Stdin)
	slog.Info("got answer",
		"answer", answer,
	)
}

func run(input io.Reader) (answer int) {
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		answer += parseLine(scanner.Text())
	}

	return
}

func parseLine(line string) int {
	slog.Info("got line",
		"line", line,
	)

	parts := strings.SplitN(line, ": ", 2)
	return getPower(parts[1])
}

func getPower(s string) int {
	power := 1
	for _, qty := range getMinQtys(s) {
		power *= qty
	}
	slog.Info("got power",
		"power", power,
	)
	return power
}

func getMinQtys(s string) map[string]int {
	result := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	moves := strings.Split(s, "; ")
	for _, move := range moves {
		for color, qty := range parseTotalQtys(move) {
			result[color] = max(result[color], qty)
		}
	}

	slog.Info("got minQtys",
		"result", result,
	)
	return result
}

func parseTotalQtys(s string) map[string]int {
	result := map[string]int{}

	cubesQtys := strings.Split(s, ", ")
	for _, cubesQty := range cubesQtys {
		color, qty := parseCubesQty(cubesQty)
		result[color] += qty
	}

	return result
}

func parseCubesQty(s string) (string, int) {
	parts := strings.Split(s, " ")
	qty, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		panic(err)
	}

	color := parts[1]
	return color, int(qty)
}
