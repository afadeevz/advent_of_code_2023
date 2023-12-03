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
	var numIDs [][]int

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		row := ([]rune)(scanner.Text())
		rows = append(rows, row)
		numIDs = append(numIDs, make([]int, len(row)))
	}

	rowsCount := len(rows)
	colsCount := len(rows[0])

	parseNumber := func(rowIdx, colIdx int) int {
		endIdx := colIdx
		for endIdx < colsCount && unicode.IsDigit(rows[rowIdx][endIdx]) {
			endIdx++
		}
		beginIdx := colIdx
		for beginIdx > 0 && unicode.IsDigit(rows[rowIdx][beginIdx-1]) {
			beginIdx--
		}

		str := string(rows[rowIdx][beginIdx:endIdx])
		slog.Info(str, "row", rowIdx, "begin", beginIdx, "end", endIdx)
		num, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			panic(err)
		}
		return int(num)
	}

	lastNumID := 0
	for rowIdx, row := range rows {
		for colIdx, cell := range row {
			if unicode.IsDigit(cell) {
				if colIdx == 0 || !unicode.IsDigit(row[colIdx-1]) {
					lastNumID++
				}
				numIDs[rowIdx][colIdx] = lastNumID
			}
		}
	}

	answer := 0
	for rowIdx, row := range rows {
		for colIdx, cell := range row {
			if cell == '*' {
				touchingIDs := make(map[int]struct{})
				product := 1
				for d := 0; d < 8; d++ {
					rowIdx2 := rowIdx + dRow[d]
					colIdx2 := colIdx + dCol[d]
					if rowIdx2 < 0 || rowIdx2 >= rowsCount || colIdx2 < 0 || colIdx2 >= colsCount {
						continue
					}

					numID := numIDs[rowIdx2][colIdx2]
					if numID == 0 {
						continue
					}

					if _, ok := touchingIDs[numID]; ok {
						continue
					}

					touchingIDs[numID] = struct{}{}
					product *= parseNumber(rowIdx2, colIdx2)
				}
				if len(touchingIDs) == 2 {
					answer += product
				}
			}
		}
	}

	return answer
}
