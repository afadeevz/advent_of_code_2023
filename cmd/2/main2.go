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
	if !isValidGame(parts[1]) {
		return 0
	}

	return parseGameID(parts[0])
}

func parseGameID(s string) int {
	slog.Info("got game id",
		"gameid", s,
	)
	parts := strings.Split(s, " ")
	gameID, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		panic(err)
	}

	return int(gameID)
}

func isValidGame(s string) bool {
	moves := strings.Split(s, "; ")
	for _, move := range moves {
		if !isValidMove(move) {
			return false
		}
	}
	return true
}

func isValidMove(s string) bool {
	cubesQtys := strings.Split(s, ", ")
	for _, cubesQty := range cubesQtys {
		if !isValidCubesQty(cubesQty) {
			return false
		}
	}

	return true
}

func isValidCubesQty(s string) bool {
	parts := strings.Split(s, " ")
	qty, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		panic(err)
	}

	color := parts[1]
	switch color {
	case "red":
		return qty <= 12
	case "green":
		return qty <= 13
	case "blue":
		return qty <= 14
	default:
		panic("invalid color: " + color)
	}
}
