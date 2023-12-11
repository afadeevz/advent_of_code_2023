package main

import (
	"bufio"
	"io"
	"log/slog"
	"os"
)

const (
	part2 = true
)

var expansionRate int

func init() {
	if part2 {
		expansionRate = 1000000
	} else {
		expansionRate = 2
	}
}

type Void = struct{}

func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

type Pos struct {
	row, col int
}

func (p Pos) Dist(other Pos) int {
	return Abs(p.row-other.row) + Abs(p.col-other.col)
}

func main() {
	slog.Info("got answer", "answer", run(os.Stdin))
}

func run(input io.Reader) int {
	galaxies := parseInput(input)
	galaxiesCount := len(galaxies)

	rowDisplacements := getRowsDisplacements(galaxies)
	colDisplacements := getColsDisplacements(galaxies)

	// slog.Info("", "rowDisplacements", rowDisplacements)
	// slog.Info("", "colDisplacements", colDisplacements)

	// slog.Info("", "galaxies", galaxies)
	for i, g := range galaxies {
		galaxies[i].row += rowDisplacements[g.row]
		galaxies[i].col += colDisplacements[g.col]
	}
	// slog.Info("", "galaxies", galaxies)

	answer := 0
	for i := 0; i < galaxiesCount; i++ {
		for j := i + 1; j < galaxiesCount; j++ {
			dist := galaxies[i].Dist(galaxies[j])
			answer += dist
		}
	}

	return answer
}

func getRowsDisplacements(galaxies []Pos) []int {
	rows := make(map[int]Void)

	maxRow := 0
	for _, g := range galaxies {
		rows[g.row] = Void{}
		maxRow = max(maxRow, g.row)
	}

	res := make([]int, maxRow+1)

	displacement := 0
	for row := 0; row <= maxRow; row++ {
		if _, ok := rows[row]; !ok {
			displacement += expansionRate - 1
		}
		res[row] = displacement
	}

	return res
}

func getColsDisplacements(galaxies []Pos) []int {
	cols := make(map[int]Void)

	maxCol := 0
	for _, g := range galaxies {
		cols[g.col] = Void{}
		maxCol = max(maxCol, g.col)
	}

	res := make([]int, maxCol+1)

	displacement := 0
	for col := 0; col <= maxCol; col++ {
		if _, ok := cols[col]; !ok {
			displacement += expansionRate - 1
		}
		res[col] = displacement
	}

	return res
}

func parseInput(input io.Reader) (galaxies []Pos) {
	scanner := bufio.NewScanner(input)

	for row := 0; scanner.Scan(); row++ {
		line := scanner.Text()
		for col, r := range line {
			if r != '#' {
				continue
			}

			galaxies = append(galaxies, Pos{row, col})
		}
	}

	return
}
