package main

import (
	"bufio"
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

var Directions = []Direction{Right, Up, Left, Down}

func (d Direction) Opposize() Direction {
	return (d + 2) % 4
}

func (d Direction) DRow() int {
	return map[Direction]int{
		Right: 0,
		Up:    -1,
		Left:  0,
		Down:  1,
	}[d]
}

func (d Direction) DCol() int {
	return map[Direction]int{
		Right: 1,
		Up:    0,
		Left:  -1,
		Down:  0,
	}[d]
}

type Cell int

const (
	PipeVertical = iota
	PipeHorizontal
	PipeUpRight
	PipeDownRight
	PipeUpLeft
	PipeDownLeft
	None
	Start
)

func (c Cell) Connections() []Direction {
	return map[Cell][]Direction{
		PipeVertical:   {Up, Down},
		PipeHorizontal: {Left, Right},
		PipeUpRight:    {Up, Right},
		PipeDownRight:  {Down, Right},
		PipeUpLeft:     {Up, Left},
		PipeDownLeft:   {Down, Left},
		None:           {},
		Start:          {},
	}[c]
}

type Row []Cell
type Field []Row

type Pos struct {
	row, col int
}

func (p Pos) ToDir(d Direction) Pos {
	return Pos{
		row: p.row + d.DRow(),
		col: p.col + d.DCol(),
	}
}

func (f Field) FindStartPos() Pos {
	for rowIdx, row := range f {
		for colIdx, cell := range row {
			if cell != Start {
				continue
			}

			return Pos{
				row: rowIdx,
				col: colIdx,
			}
		}
	}

	panic("start pos was not found")
}

func (f Field) RowsCount() int {
	return len(f)
}

func (f Field) ColsCount() int {
	return len(f[0])
}

func (f Field) Contains(pos Pos) bool {
	return 0 <= pos.row && pos.row < f.RowsCount() && 0 <= pos.col && pos.col < f.ColsCount()
}

func (f Field) At(pos Pos) Cell {
	return f[pos.row][pos.col]
}

func main() {
	slog.Info("got answer", "answer", run(os.Stdin))
}

func traverse(f Field, pos Pos, dirFrom Direction) int {
	if !f.Contains(pos) {
		return 0
	}
	if f.At(pos) == Start {
		return 1
	}

	for _, dir := range f.At(pos).Connections() {
		if dir == dirFrom.Opposize() {
			continue
		}

		dist := traverse(f, pos.ToDir(dir), dir)
		if dist > 0 {
			return dist + 1
		}
	}

	return 0
}

func run(input io.Reader) int {
	field := parseInput(input)
	startPos := field.FindStartPos()

	for _, dir := range Directions {
		dist := traverse(field, startPos.ToDir(dir), dir)
		if dist == 0 {
			continue
		}

		return dist / 2
	}

	return 0
}

func parseInput(input io.Reader) (field Field) {
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		rowStr := scanner.Text()
		row := parseRow(rowStr)
		field = append(field, row)
	}

	return
}

func parseRow(s string) (row Row) {
	for _, r := range s {
		row = append(row, parseCell(r))
	}

	return
}

func parseCell(r rune) Cell {
	mapping := map[rune]Cell{
		'|': PipeVertical,
		'-': PipeHorizontal,
		'L': PipeUpRight,
		'F': PipeDownRight,
		'J': PipeUpLeft,
		'7': PipeDownLeft,
		'.': None,
		'S': Start,
	}

	return mapping[r]
}
