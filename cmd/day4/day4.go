package main

import (
	"bufio"
	"io"
	"log/slog"
	"os"
	"strings"
)

type Void = struct{}
type Numbers = map[string]Void

func main() {
	answer := run(os.Stdin)
	slog.Info("got answer", "answer", answer)
}

func run(input io.Reader) (answer int) {
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		answer += processLine(scanner.Text())
	}

	return
}

func processLine(line string) int {
	postCard := strings.Split(line, ":")[1]
	parts := strings.Split(postCard, "|")
	winningNumbers := parseNumbers(parts[0])
	numbers := parseNumbers(parts[1])

	matches := 0

	for num := range numbers {
		if _, ok := winningNumbers[num]; !ok {
			continue
		}

		matches++
	}

	return calculatePoints(matches)
}

func parseNumbers(numbers string) (result Numbers) {
	parts := strings.Split(numbers, " ")
	result = make(Numbers, len(parts))

	for _, part := range parts {
		if len(part) == 0 {
			continue
		}

		result[part] = Void{}
	}

	return
}

func calculatePoints(matches int) int {
	if matches == 0 {
		return 0
	}

	return 1 << (matches - 1)
}
