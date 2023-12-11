package main

import (
	"bufio"
	"io"
	"log/slog"
	"os"
	"strconv"
	"unicode"
)

func main() {
	slog.Info("finished", "answer", run(os.Stdin))
}

type (
	Row  = []rune
	Rows = []Row
)

var (
	dRow = []int{0, -1, -1, -1, 0, 1, 1, 1}
	dCol = []int{1, 1, 0, -1, -1, -1, 0, 1}
)

func run(input io.Reader) int {
	var rows Rows
	var vis [][]bool

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		row := ([]rune)(scanner.Text())
		rows = append(rows, row)
		vis = append(vis, make([]bool, len(row)))
	}

	rowsCount := len(rows)
	colsCount := len(rows[0])

	var isTouching func(int, int) bool
	isTouchingImpl := func(rowIdx, colIdx int) bool {
		if rowIdx < 0 || rowIdx >= rowsCount || colIdx < 0 || colIdx >= colsCount {
			return false
		}

		cell := rows[rowIdx][colIdx]
		if isSymbol(cell) {
			return true
		}
		if !unicode.IsDigit(cell) {
			return false
		}
		if vis[rowIdx][colIdx] {
			return false
		}

		vis[rowIdx][colIdx] = true
		for d := 0; d < 8; d++ {
			if isTouching(rowIdx+dRow[d], colIdx+dCol[d]) {
				return true
			}
		}

		return false
	}
	isTouching = isTouchingImpl

	isNumberStart := func(rowIdx, colIdx int) bool {
		row := rows[rowIdx]
		cell := row[colIdx]
		if !unicode.IsDigit(cell) {
			return false
		}

		if colIdx > 0 && unicode.IsDigit(row[colIdx-1]) {
			return false
		}

		return true
	}

	parseNumber := func(rowIdx, colIdx int) int {
		endIdx := colIdx
		for endIdx < colsCount && unicode.IsDigit(rows[rowIdx][endIdx]) {
			endIdx++
		}

		str := string(rows[rowIdx][colIdx:endIdx])
		slog.Info(str, "row", rowIdx, "col", colIdx, "end", endIdx)
		num, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			panic(err)
		}
		return int(num)
	}

	answer := 0
	for rowIdx, row := range rows {
		for colIdx := range row {
			if isNumberStart(rowIdx, colIdx) && isTouching(rowIdx, colIdx) {
				answer += parseNumber(rowIdx, colIdx)
			}
		}
	}

	return answer
}

func isSymbol(r rune) bool {
	return r != '.' && !unicode.IsDigit(r)
}
