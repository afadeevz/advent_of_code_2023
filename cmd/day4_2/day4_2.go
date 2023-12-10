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

func run(input io.Reader) int {
	scanner := bufio.NewScanner(input)

	results := make([]int, 0)

	for scanner.Scan() {
		results = append(results, processLine(scanner.Text()))
	}

	cardsCounts := make([]int, len(results))
	for i := range cardsCounts {
		cardsCounts[i] = 1
	}

	for i, result := range results {
		for j := i + 1; j <= i+result && j < len(results); j++ {
			cardsCounts[j] += cardsCounts[i]
		}
	}

	answer := 0
	for _, count := range cardsCounts {
		answer += count
	}

	return answer
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

	return matches
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
