package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

const inputData = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

func TestDay11(t *testing.T) {
	t.Parallel()

	input := bytes.NewBufferString(inputData)

	expansionRate = 2
	assert.Equal(t, 374, run(input))
}

func TestDay11_2(t *testing.T) {
	t.Parallel()

	input := bytes.NewBufferString(inputData)

	expansionRate = 100
	assert.Equal(t, 8410, run(input))
}
