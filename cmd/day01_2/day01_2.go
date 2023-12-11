package main

import (
	"bufio"
	"io"
	"log/slog"
	"os"
	"unicode"
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
	digits := extractDigits(runes)

	first := digits[0]
	last := digits[len(digits)-1]

	slog.Info("parsed line",
		"line", line,
		"first", first,
		"last", last,
	)
	return 10*first + last
}

func extractDigits(runes []rune) (digits []int) {
	for i := range runes {
		if digit := extractDigit(runes[i:]); digit >= 0 {
			digits = append(digits, digit)
		}
	}

	return
}

func extractDigit(runes []rune) int {
	if unicode.IsDigit(runes[0]) {
		return parseDigit(runes[0])
	}

	mapping := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for word, digit := range mapping {
		if hasPrefix(runes, word) {
			return digit
		}
	}

	return -1
}

func hasPrefix(runes []rune, prefixStr string) bool {
	prefix := []rune(prefixStr)
	if len(runes) < len(prefixStr) {
		return false
	}

	for i := range prefix {
		if runes[i] != prefix[i] {
			return false
		}
	}

	return true
}

func parseDigit(b rune) int {
	return int(b - '0')
}
