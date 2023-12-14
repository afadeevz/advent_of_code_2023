package main

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

const inputData = `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`

//go:embed input14.txt
var inputFileData string

func TestDay14Part1(t *testing.T) {
	t.Parallel()
	input := bytes.NewBufferString(inputData)
	assert.Equal(t, 136, part1(input))
}

func TestDay14InputFile(t *testing.T) {
	t.Parallel()
	input := bytes.NewBufferString(inputFileData)
	assert.Equal(t, 112773, part1(input))
}

func TestDay14Part2(t *testing.T) {
	t.Parallel()
	input := bytes.NewBufferString(inputData)
	assert.Equal(t, 64, part2(input))
}
