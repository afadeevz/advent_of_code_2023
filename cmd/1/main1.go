package main

import (
	"bufio"
	"io"
	"log/slog"
	"os"
	"slices"
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
	slog.Debug("got line", "line", line)
	runes := []rune(line)

	first := firstDigit(runes)

	slices.Reverse(runes)
	last := firstDigit(runes)

	slog.Info("parsed line",
		"line", line,
		"first", first,
		"last", last,
	)
	return 10*first + last
}

func firstDigit(runes []rune) int {
	for _, b := range runes {
		if isDigit(b) {
			return parseDigit(b)
		}
	}

	panic("not found")
}

func isDigit(b rune) bool {
	return '0' < b && b <= '9'
}

func parseDigit(b rune) int {
	return int(b - '0')
}
