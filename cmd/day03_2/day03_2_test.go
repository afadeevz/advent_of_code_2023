package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay3(t *testing.T) {
	t.Parallel()

	input := bytes.NewBufferString(`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`)

	assert.Equal(t, 467835, run(input))
}
