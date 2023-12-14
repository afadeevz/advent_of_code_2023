package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"os"
)

type Direction int

const (
	Right Direction = iota
	Up
	Left
	Down
)

func (d Direction) dRow() int {
	return map[Direction]int{
		Right: 0,
		Up:    -1,
		Left:  0,
		Down:  1,
	}[d]
}

func (d Direction) dCol() int {
	return map[Direction]int{
		Right: 1,
		Up:    0,
		Left:  -1,
		Down:  0,
	}[d]
}

type Position struct {
	row, col int
}

func (p Position) toDir(d Direction) Position {
	return Position{
		row: p.row + d.dRow(),
		col: p.col + d.dCol(),
	}
}

type Cell = rune

const (
	Empty Cell = '.'
	Cube  Cell = '#'
	Round Cell = 'O'
)

type Row = []Cell
type Field []Row

func (f Field) rowsCount() int {
	return len(f)
}

func (f Field) colsCount() int {
	return len(f[0])
}

func (f Field) contains(pos Position) bool {
	return 0 <= pos.row && pos.row < f.rowsCount() && 0 <= pos.col && pos.col < f.colsCount()
}

func (f Field) at(pos Position) Cell {
	return f[pos.row][pos.col]
}

func (f Field) set(pos Position, cell Cell) {
	f[pos.row][pos.col] = cell
}

func (f Field) tiltAt(pos Position, dir Direction) {
	// if !f.contains(pos) {
	// 	return
	// }

	if f.at(pos) != Round {
		return
	}

	nextPos := pos.toDir(dir)
	if !f.contains(nextPos) {
		return
	}

	f.tiltAt(nextPos, dir)

	if f.at(nextPos) != Empty {
		return
	}

	f.set(pos, Empty)
	f.set(nextPos, Round)
	f.tiltAt(nextPos, dir)
}

func (f Field) tilt(dir Direction) {
	for i := range f {
		for j := range f[i] {
			f.tiltAt(Position{i, j}, dir)
		}
	}
}

func (f Field) CalcLoad() int {
	load := 0

	for i, row := range f {
		for _, cell := range row {
			if cell != Round {
				continue
			}

			load += f.colsCount() - i
		}
	}

	return load
}

func main() {
	inputData, _ := io.ReadAll(os.Stdin)

	input1 := bytes.NewBuffer(inputData)
	slog.Info("got answer part1", "answer", part1(input1))

	input2 := bytes.NewBuffer(inputData)
	slog.Info("got answer part2", "answer", part2(input2))
}

func part1(input io.Reader) int {
	field := parseInput(input)
	field.tilt(Up)
	return field.CalcLoad()
}

func part2(input io.Reader) int {
	cycle := []Direction{Up, Left, Down, Right}
	field := parseInput(input)

	mapping := make(map[string]int)

	lim := int(1e9)
	for i := 0; i < lim; i++ {
		str := fmt.Sprintf("%+v", field)
		if j, ok := mapping[str]; ok {
			loopLen := i - j
			skip := (lim - i) / loopLen * loopLen
			slog.Info("found loop", "i", i, "j", j, "loopLen", loopLen, "skip", skip)
			i += skip
		} else {
			mapping[str] = i
		}

		for _, dir := range cycle {
			field.tilt(dir)
		}

	}

	return field.CalcLoad()
}

func parseInput(input io.Reader) (field Field) {
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		rowStr := scanner.Text()
		field = append(field, parseRow(rowStr))
	}

	return
}

func parseRow(rowStr string) (row Row) {
	row = make(Row, 0, len(rowStr))
	for _, r := range rowStr {
		row = append(row, Cell(r))
	}

	return
}
