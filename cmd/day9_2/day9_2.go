package main

import (
	"bufio"
	"io"
	"log/slog"
	"os"
	"slices"
	"strconv"
	"strings"
)

type History []int

func (h History) Extrapolate() int {
	l := len(h)
	if l == 0 {
		return 0
	}
	if l == 1 {
		return h[0]
	}

	allZeroes := true
	for _, value := range h {
		if value != 0 {
			allZeroes = false
			break
		}
	}
	if allZeroes {
		return 0
	}

	deltas := h.GetDeltas()
	return h[l-1] + deltas.Extrapolate()
}

func (h History) GetDeltas() History {
	l := len(h)
	deltas := make(History, 0, l-1)

	for i := 1; i < l; i++ {
		deltas = append(deltas, h[i]-h[i-1])
	}

	return deltas
}

func main() {
	slog.Info("got answer", "answer", run(os.Stdin))
}

func run(input io.Reader) int {
	histories := parseInput(input)

	answer := 0

	for _, history := range histories {
		answer += history.Extrapolate()
	}

	return answer
}

func parseInput(input io.Reader) (histories []History) {
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		historyStr := scanner.Text()
		history := parseHistory(historyStr)
		histories = append(histories, history)
	}

	return
}

func parseHistory(s string) (history History) {
	values := strings.Split(s, " ")
	for _, value := range values {
		tmp, _ := strconv.ParseInt(value, 10, 64)
		history = append(history, int(tmp))
	}

	slices.Reverse[History](history)
	return
}
